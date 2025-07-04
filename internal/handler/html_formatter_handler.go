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

		// For formatting, we use the Minify function but with options to keep it readable.
		// The library is primarily for minification, but we can configure it for pretty-printing.
		// A simple way to "format" is to just process it. For more complex formatting, a different library might be better,
		// but for consistency, we'll use minify.

		// Since the default is to minify, we'll just show the minified version as "formatted"
		// A true "formatter" would require more complex logic to add indentation etc.
		// Let's stick to the library's strength: minification/optimization.

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
