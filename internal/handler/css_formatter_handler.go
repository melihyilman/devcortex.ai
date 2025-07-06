package handler

import (
	"net/http"
	"fmt"

	"devcortex.ai/internal/view"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
)

func CSSFormatterTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "CSS Formatter",
		ToolSpecificData: map[string]interface{}{
			"CSSInput": r.URL.Query().Get("CSSInput"),
			"Result":   r.URL.Query().Get("Result"),
			"Success":  r.URL.Query().Get("Success") == "true",
		},
	}

	if r.Method == http.MethodPost {
		cssInput := r.FormValue("css_input")
		var result string
		var success bool

		m := minify.New()
		m.AddFunc("text/css", css.Minify)

		formatted, err := m.String("text/css", cssInput)
		if err != nil {
			result = "Error formatting CSS: " + err.Error()
		} else {
			result = formatted
			success = true
		}

		redirectData := map[string]string{
			"CSSInput": cssInput,
			"Result":   result,
			"Success":  fmt.Sprintf("%t", success),
		}
		redirectToPageWithData(w, r, redirectData)
		return
	}

	view.Render(w, r, "css-formatter.html", data)
}
