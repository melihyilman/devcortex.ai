package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"devcortex.ai/internal/view"
)

type JSONData struct {
	InputJSON string
	Result    string
	IsValid   bool
}

func JSONTool(w http.ResponseWriter, r *http.Request) {
	pageData := &view.PageData{
		Title: "JSON Validator & Formatter",
		ToolSpecificData: nil,
	}

	if r.Method == http.MethodPost {
		inputJSON := r.FormValue("jsonInput")
		var result string
		isValid := false

		var js interface{}
		if err := json.Unmarshal([]byte(inputJSON), &js); err != nil {
			result = "JSON Parse Hatası: " + err.Error()
			isValid = false
		} else {
			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, []byte(inputJSON), "", "  "); err != nil {
				result = "JSON Formatlama Hatası: " + err.Error()
				isValid = false
			} else {
				result = prettyJSON.String()
				isValid = true
			}
		}

		pageData.ToolSpecificData = map[string]interface{}{
			"InputJSON": inputJSON,
			"Result":    result,
			"IsValid":   isValid,
		}
	}

	view.Render(w, r, "json.html", pageData)
}
