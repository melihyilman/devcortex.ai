package handler

import (
	"bytes"
	"net/http"

	"devcortex.ai/internal/view"
	"github.com/yuin/goldmark"
)

func MarkdownPreviewerTool(w http.ResponseWriter, r *http.Request) {
	pageData := &view.PageData{
		Title: "Markdown Previewer",
		ToolSpecificData: make(map[string]interface{}),
	}

	if r.Method == http.MethodPost {
		markdownInput := r.FormValue("markdown")
		var buf bytes.Buffer
		if err := goldmark.Convert([]byte(markdownInput), &buf); err != nil {
			pageData.ToolSpecificData.(map[string]interface{})["Result"] = "Error converting markdown: " + err.Error()
		} else {
			pageData.ToolSpecificData.(map[string]interface{})["Result"] = buf.String()
		}
		pageData.ToolSpecificData.(map[string]interface{})["Markdown"] = markdownInput
	}

	view.Render(w, r, "markdown-previewer.html", pageData)
}
