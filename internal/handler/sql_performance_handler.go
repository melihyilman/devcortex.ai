package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"reflect"

	"devcortex.ai/internal/view"
	"github.com/xwb1989/sqlparser"
)

// Recommendation holds structured advice about a SQL query.
type Recommendation struct {
	Severity string // "Critical", "Warning", "Info"
	Title    string
	Detail   string
}

func SQLPerformanceTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "SQL Performance Analyzer",
	}

	if r.Method == http.MethodPost {
		sqlQuery := r.FormValue("sql_query")
		recommendations, ast := analyzeSQL(sqlQuery)

		var astGraph string
		if ast != nil {
			astGraph = generateASTGraph(ast)
		}

		data.ToolSpecificData = map[string]interface{}{
			"SQLQuery":                   sqlQuery,
			"PerformanceRecommendations": recommendations,
			"ASTGraph":                   astGraph,
		}
	}

	view.Render(w, r, "sql-performance.html", data)
}

func analyzeSQL(query string) ([]Recommendation, sqlparser.SQLNode) {
	var recommendations []Recommendation
	stmt, err := sqlparser.Parse(query)
	if err != nil {
		return []Recommendation{{
			Severity: "Critical",
			Title:    "SQL Parsing Error",
			Detail:   fmt.Sprintf("Could not parse the SQL query: %v", err),
		}}, nil
	}

	// Use a map to avoid duplicate recommendations
	recMap := make(map[string]Recommendation)

	sqlparser.Walk(func(node sqlparser.SQLNode) (kontinue bool, err error) {
		switch n := node.(type) {
		case *sqlparser.Select:
			// Rule: Avoid SELECT *
			for _, sel := range n.SelectExprs {
				if _, ok := sel.(*sqlparser.StarExpr); ok {
					recMap["select-star"] = Recommendation{
						Severity: "Warning",
						Title:    "Avoid using 'SELECT *'",
						Detail:   "Specify the exact columns you need. This reduces data transfer, processing load, and makes code more maintainable.",
					}
					break
				}
			}
			// Rule: ORDER BY without LIMIT
			if n.OrderBy != nil && n.Limit == nil {
				recMap["orderby-no-limit"] = Recommendation{
					Severity: "Info",
					Title:    "ORDER BY without LIMIT",
					Detail:   "Sorting a large result set without a LIMIT can be resource-intensive. Ensure you are not sorting more data than necessary.",
				}
			}
		case *sqlparser.Update:
			if n.Where == nil {
				recMap["update-no-where"] = Recommendation{
					Severity: "Critical",
					Title:    "UPDATE without WHERE clause",
					Detail:   "This query will update all rows in the table. This is rarely intended. Double-check your logic.",
				}
			}
		case *sqlparser.Delete:
			if n.Where == nil {
				recMap["delete-no-where"] = Recommendation{
					Severity: "Critical",
					Title:    "DELETE without WHERE clause",
					Detail:   "This query will delete all rows from the table. This is rarely intended. Double-check your logic.",
				}
			}
		case *sqlparser.ComparisonExpr:
			// Rule: LIKE with leading wildcard
			if n.Operator == sqlparser.LikeStr {
				if val, ok := n.Right.(*sqlparser.SQLVal); ok {
					if len(val.Val) > 0 && val.Val[0] == '%' {
						recMap["leading-wildcard"] = Recommendation{
							Severity: "Warning",
							Title:    "LIKE with leading wildcard",
							Detail:   "Using a wildcard '%' at the beginning of a LIKE pattern (e.g., LIKE '%value') prevents the database from using an index on that column, forcing a full table scan.",
						}
					}
				}
			}
			// Rule: Function on a column in WHERE clause
			if _, ok := n.Left.(*sqlparser.FuncExpr); ok {
				recMap["function-on-column"] = Recommendation{
					Severity: "Warning",
					Title:    "Function on indexed column",
					Detail:   "Applying a function (e.g., LOWER(), DATE()) to a column in the WHERE clause can prevent the database from using an index on that column. Try to apply the function to the value instead.",
				}
			}
			// Rule: IN clause with a subquery
			if n.Operator == sqlparser.InStr {
				if _, ok := n.Right.(*sqlparser.Subquery); ok {
					recMap["in-subquery"] = Recommendation{
						Severity: "Info",
						Title:    "IN clause with a subquery",
						Detail:   "For large subqueries, using 'IN' can be inefficient. Consider rewriting the query using a JOIN or an EXISTS clause, which can often be optimized better by the database.",
					}
				}
			}
		}
		return true, nil
	}, stmt)

	if len(recMap) == 0 {
		recommendations = append(recommendations, Recommendation{
			Severity: "Success",
			Title:    "No obvious issues found",
			Detail:   "Based on a static analysis of common anti-patterns, this query looks good. For a complete picture, always analyze the query's execution plan against your actual database.",
		})
	} else {
		for _, r := range recMap {
			recommendations = append(recommendations, r)
		}
	}

	return recommendations, stmt
}

func generateASTGraph(root sqlparser.SQLNode) string {
	var buf bytes.Buffer
	buf.WriteString("digraph AST {\n")
	buf.WriteString("  rankdir=TB;\n")
	buf.WriteString("  bgcolor=\"transparent\";\n")
	buf.WriteString("  node [shape=box, style=rounded, fontname=\"Helvetica\", fontcolor=\"white\", color=\"#cccccc\"];\n")
	buf.WriteString("  edge [fontname=\"Helvetica\", color=\"#cccccc\"];\n")

	var counter int
	var walk func(node sqlparser.SQLNode, parentID int)

	walk = func(node sqlparser.SQLNode, parentID int) {
		if node == nil || (reflect.ValueOf(node).Kind() == reflect.Ptr && reflect.ValueOf(node).IsNil()) {
			return
		}

		nodeID := counter
		counter++

		nodeType := reflect.TypeOf(node).Elem().Name()
		var nodeLabel string
		switch n := node.(type) {
		case *sqlparser.SQLVal:
			nodeLabel = fmt.Sprintf("%s\\n(%s)", nodeType, n.Val)
		case *sqlparser.ColName:
			nodeLabel = fmt.Sprintf("%s\\n(%s)", nodeType, n.Name.String())
		default:
			nodeLabel = nodeType
		}
		buf.WriteString(fmt.Sprintf("  node%d [label=\"%s\"];\n", nodeID, nodeLabel))

		if parentID != -1 {
			buf.WriteString(fmt.Sprintf("  node%d -> node%d;\n", parentID, nodeID))
		}

		// Recursively walk children for common node types
		switch n := node.(type) {
		case *sqlparser.Select:
			for _, expr := range n.SelectExprs {
				walk(expr, nodeID)
			}
			walk(n.From, nodeID)
			walk(n.Where, nodeID)
		case *sqlparser.AliasedTableExpr:
			walk(n.Expr, nodeID)
		case *sqlparser.JoinTableExpr:
			walk(n.LeftExpr, nodeID)
			walk(n.RightExpr, nodeID)
			// Check if the On condition is not empty
			if n.Condition.On != nil {
				walk(n.Condition.On, nodeID)
			}
		case *sqlparser.Where:
			walk(n.Expr, nodeID)
		case *sqlparser.ComparisonExpr:
			walk(n.Left, nodeID)
			walk(n.Right, nodeID)
		case *sqlparser.AndExpr:
			walk(n.Left, nodeID)
			walk(n.Right, nodeID)
		case *sqlparser.OrExpr:
			walk(n.Left, nodeID)
			walk(n.Right, nodeID)
		case *sqlparser.AliasedExpr:
			walk(n.Expr, nodeID)
		}
	}

	walk(root, -1)
	buf.WriteString("}\n")
	return buf.String()
}
