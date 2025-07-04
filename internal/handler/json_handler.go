package handler

import (
	"net/http"

	"devcortex.ai/internal/tool"
	"devcortex.ai/internal/view"
)

type JSONData struct {
	InputJSON string
	Result    string
	IsValid   bool
	Error     string
}

func JSONTool(w http.ResponseWriter, r *http.Request) {
	data := JSONData{}
	jsonService := tool.NewJSONFormatter()

	if r.Method == http.MethodPost {
		inputJSON := r.FormValue("jsonInput")
		data.InputJSON = inputJSON

		formatted, isValid, err := jsonService.Format(inputJSON)
		if err != nil {
			data.Error = err.Error()
			data.IsValid = false
			data.Result = inputJSON // Show the original input on error
		} else {
			data.Result = formatted
			data.IsValid = isValid
		}
	}

	view.Render(w, r, "json.html", &view.PageData{
		Title:            "JSON Validator & Formatter",
		ToolSpecificData: data,
	})
}
