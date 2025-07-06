package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
)

func HTMLFormatterTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "HTML Formatter",
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

		data.ToolSpecificData = map[string]interface{}{
			"HTMLInput": htmlInput,
			"Result":    result,
			"Success":   success,
		}
	}

	view.Render(w, r, "html-formatter.html", data)
}
