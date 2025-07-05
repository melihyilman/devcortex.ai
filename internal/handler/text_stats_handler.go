package handler

import (
	"net/http"
	"strings"
	"unicode/utf8"

	"devcortex.ai/internal/view"
)

func TextStatsTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Text Statistics",
	}

	if r.Method == http.MethodPost {
		textInput := r.FormValue("text_input")

		if textInput != "" {
			wordCount := len(strings.Fields(textInput))
			charCount := utf8.RuneCountInString(textInput)
			lineCount := len(strings.Split(textInput, "\n"))

			data.ToolSpecificData = map[string]interface{}{
				"TextInput": textInput,
				"WordCount": wordCount,
				"CharCount": charCount,
				"LineCount": lineCount,
				"HasResult": true,
			}
		}
	}

	view.Render(w, r, "text-stats.html", data)
}
