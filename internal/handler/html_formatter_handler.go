package handler

import (
	"net/http"
	"fmt"

	"devcortex.ai/internal/view"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
)

func HTMLFormatterTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "HTML Formatter",
		ToolSpecificData: map[string]interface{}{
			"HTMLInput": r.URL.Query().Get("HTMLInput"),
			"Result":    r.URL.Query().Get("Result"),
			"Success":   r.URL.Query().Get("Success") == "true",
		},
	}

	if r.Method == http.MethodPost {
		htmlInput := r.FormValue("html_input")
		var result string
		var success bool

		m := minify.New()
		m.AddFunc("text/html", html.Minify)

		minified, err := m.String("text/html", htmlInput)
		if err != nil {
			result = "Error formatting HTML: " + err.Error()
		} else {
			result = minified
			success = true
		}

		redirectData := map[string]string{
			"HTMLInput": htmlInput,
			"Result":    result,
			"Success":   fmt.Sprintf("%t", success),
		}
		redirectToPageWithData(w, r, redirectData)
		return
	}

	view.Render(w, r, "html-formatter.html", data)
}
