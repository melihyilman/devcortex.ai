package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
)

func CSSFormatterTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "CSS Formatter",
	}

	if r.Method == http.MethodPost {
		cssInput := r.FormValue("css_input")
		var result string
		var success bool

		m := minify.New()
		m.AddFunc("text/css", css.Minify)

		// Using the minify library to format the CSS.
		formatted, err := m.String("text/css", cssInput)
		if err != nil {
			result = "Error formatting CSS: " + err.Error()
		} else {
			result = formatted
			success = true
		}

		data.ToolSpecificData = map[string]interface{}{
			"CSSInput": cssInput,
			"Result":   result,
			"Success":  success,
		}
	}

	view.Render(w, r, "css-formatter.html", data)
}
