package handler

import (
	"html/template"
	"net/http"
	"os"

	"devcortex.ai/internal/view"
	"github.com/gomarkdown/markdown"
)

func ExtensionsHandler(w http.ResponseWriter, r *http.Request) {
	md, err := os.ReadFile("vscode-extension/TOOLS.md")
	if err != nil {
		http.Error(w, "Failed to read Markdown file", http.StatusInternalServerError)
		return
	}

	html := markdown.ToHTML(md, nil, nil)

	data := &view.PageData{
		Title:   "Extensions",
		Content: template.HTML(html),
	}

	view.Render(w, r, "extensions.html", data)
}
