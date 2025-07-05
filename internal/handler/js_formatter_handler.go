package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/js"
)

func JSFormatterTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "JavaScript Formatter",
	}

	if r.Method == http.MethodPost {
		jsInput := r.FormValue("js_input")
		var result string
		var success bool

		m := minify.New()
		m.AddFunc("application/javascript", js.Minify)

		// Note: This library is for minification. A true "pretty-printer" would require a different approach.
		// For this tool, "formatting" means creating a syntactically correct, minified version.
		minified, err := m.String("application/javascript", jsInput)
		if err != nil {
			result = "Error formatting JavaScript: " + err.Error()
		} else {
			result = minified
			success = true
		}

		data.ToolSpecificData = map[string]interface{}{
			"JSInput": jsInput,
			"Result":  result,
			"Success": success,
		}
	}

	view.Render(w, r, "js-formatter.html", data)
}
