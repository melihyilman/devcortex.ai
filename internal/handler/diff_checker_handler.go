package handler

import (
	"net/http"
	"strings"

	"devcortex.ai/internal/view"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func DiffCheckerTool(w http.ResponseWriter, r *http.Request) {
	pageData := &view.PageData{
		Title: "Diff Checker",
		ToolSpecificData: make(map[string]interface{}),
	}

	if r.Method == http.MethodPost {
		text1 := r.FormValue("text1")
		text2 := r.FormValue("text2")

		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(text1, text2, false)

		var diffResult strings.Builder
		for _, diff := range diffs {
			switch diff.Type {
			case diffmatchpatch.DiffInsert:
				diffResult.WriteString("<span style=\"background-color:#e6ffe6;\">" + diff.Text + "</span>")
			case diffmatchpatch.DiffDelete:
				diffResult.WriteString("<span style=\"background-color:#ffe6e6;text-decoration:line-through;\">" + diff.Text + "</span>")
			case diffmatchpatch.DiffEqual:
				diffResult.WriteString("<span>" + diff.Text + "</span>")
			}
		}

		pageData.ToolSpecificData.(map[string]interface{})["Text1"] = text1
		pageData.ToolSpecificData.(map[string]interface{})["Text2"] = text2
		pageData.ToolSpecificData.(map[string]interface{})["Diff"] = diffResult.String()
	}

	view.Render(w, r, "diff-checker.html", pageData)
}
