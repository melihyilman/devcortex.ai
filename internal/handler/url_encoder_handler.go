package handler

import (
	"net/http"
	"net/url"

	"devcortex.ai/internal/view"
)

func URLEncoderTool(w http.ResponseWriter, r *http.Request) {
	pageData := &view.PageData{
		Title: "URL Encoder/Decoder",
		ToolSpecificData: make(map[string]interface{}),
	}

	pageData.ToolSpecificData.(map[string]interface{})["Operation"] = "encode"

	if r.Method == http.MethodPost {
		inputText := r.FormValue("inputText")
		operation := r.FormValue("operation")
		var result string

		if operation == "encode" {
			result = url.QueryEscape(inputText)
		} else {
			decoded, err := url.QueryUnescape(inputText)
			if err != nil {
				result = "Error: Invalid URL encoded string!"
			} else {
				result = decoded
			}
		}

		pageData.ToolSpecificData.(map[string]interface{})["InputText"] = inputText
		pageData.ToolSpecificData.(map[string]interface{})["Operation"] = operation
		pageData.ToolSpecificData.(map[string]interface{})["Result"] = result
	}

	view.Render(w, r, "url-encoder.html", pageData)
}