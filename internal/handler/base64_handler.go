package handler

import (
	"encoding/base64"
	"net/http"

	"devcortex.ai/internal/view"
)

type Base64Data struct {
	InputText string
	Operation string
	Result    string
}

func Base64Tool(w http.ResponseWriter, r *http.Request) {
	pageData := &view.PageData{
		Title: "Base64 Tool",
		ToolSpecificData: map[string]interface{}{
			"Operation": "encode",
		},
	}

	if r.Method == http.MethodPost {
		inputText := r.FormValue("inputText")
		operation := r.FormValue("operation")
		var result string
		if operation == "encode" {
			result = base64.StdEncoding.EncodeToString([]byte(inputText))
		} else {
			decodedBytes, err := base64.StdEncoding.DecodeString(inputText)
			if err != nil {
				result = "Error: Invalid Base64 text!"
			} else {
				result = string(decodedBytes)
			}
		}
		pageData.ToolSpecificData = map[string]interface{}{
			"InputText": inputText,
			"Operation": operation,
			"Result":    result,
		}
	}

	view.Render(w, r, "base64.html", pageData)
}