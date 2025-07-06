package handler

import (
	"net/http"
	"fmt"

	"devcortex.ai/internal/view"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/js"
)

func JSFormatterTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "JavaScript Formatter",
		ToolSpecificData: map[string]interface{}{
			"JSInput": r.URL.Query().Get("JSInput"),
			"Result":  r.URL.Query().Get("Result"),
			"Success": r.URL.Query().Get("Success") == "true",
		},
	}

	if r.Method == http.MethodPost {
		jsInput := r.FormValue("js_input")
		var result string
		var success bool

		m := minify.New()
		m.AddFunc("application/javascript", js.Minify)

		minified, err := m.String("application/javascript", jsInput)
		if err != nil {
			result = "Error formatting JavaScript: " + err.Error()
		} else {
			result = minified
			success = true
		}

		redirectData := map[string]string{
			"JSInput": jsInput,
			"Result":  result,
			"Success": fmt.Sprintf("%t", success),
		}
		redirectToPageWithData(w, r, redirectData)
		return
	}

	view.Render(w, r, "js-formatter.html", data)
}
