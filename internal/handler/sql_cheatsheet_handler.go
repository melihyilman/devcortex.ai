package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func SQLCheatsheetTool(w http.ResponseWriter, r *http.Request) {
	cheatsheetData := map[string][]Command{
		"Data Query Language (DQL)": {
			{Command: "SELECT [column1], [column2]", Description: "Selects specific columns from a table."},
			{Command: "SELECT * FROM [table_name]", Description: "Selects all columns from a table."},
			{Command: "WHERE [condition]", Description: "Filters records based on a condition."},
			{Command: "ORDER BY [column] ASC|DESC", Description: "Sorts the result set."},
			{Command: "GROUP BY [column]", Description: "Groups rows that have the same values into summary rows."},
			{Command: "HAVING [condition]", Description: "Filters groups based on a condition, used with GROUP BY."},
		},
		"Data Manipulation Language (DML)": {
			{Command: "INSERT INTO [table] (col1, col2) VALUES (val1, val2)", Description: "Adds a new row to a table."},
			{Command: "UPDATE [table] SET [col1]=[val1] WHERE [condition]", Description: "Modifies existing records in a table."},
			{Command: "DELETE FROM [table] WHERE [condition]", Description: "Deletes existing records from a table."},
		},
		"Data Definition Language (DDL)": {
			{Command: "CREATE TABLE [table_name] (...)", Description: "Creates a new table in the database."},
			{Command: "ALTER TABLE [table_name] ADD [column] [datatype]", Description: "Adds a column to a table."},
			{Command: "DROP TABLE [table_name]", Description: "Deletes a table."},
			{Command: "CREATE INDEX [index_name] ON [table_name] (column1, ...)", Description: "Creates an index on a table."},
		},
		"Joins": {
			{Command: "INNER JOIN", Description: "Returns records that have matching values in both tables."},
			{Command: "LEFT JOIN", Description: "Returns all records from the left table, and the matched records from the right table."},
			{Command: "RIGHT JOIN", Description: "Returns all records from the right table, and the matched records from the left table."},
			{Command: "FULL OUTER JOIN", Description: "Returns all records when there is a match in either left or right table."},
		},
		"Aggregate Functions": {
			{Command: "COUNT()", Description: "Returns the number of rows."},
			{Command: "SUM()", Description: "Returns the total sum of a numeric column."},
			{Command: "AVG()", Description: "Returns the average value of a numeric column."},
			{Command: "MIN()", Description: "Returns the smallest value of the selected column."},
			{Command: "MAX()", Description: "Returns the largest value of the selected column."},
		},
	}

	data := &view.PageData{
		Title:            "SQL Cheatsheet",
		ToolSpecificData: cheatsheetData,
	}
	view.Render(w, r, "sql-cheatsheet.html", data)
}
