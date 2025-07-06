package handler

import (
	"net/http"
	"strconv"

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
	isValid, _ := strconv.ParseBool(r.URL.Query().Get("IsValid"))
	data := JSONData{
		InputJSON: r.URL.Query().Get("InputJSON"),
		Result:    r.URL.Query().Get("Result"),
		IsValid:   isValid,
		Error:     r.URL.Query().Get("Error"),
	}
	jsonService := tool.NewJSONFormatter()

	if r.Method == http.MethodPost {
		inputJSON := r.FormValue("jsonInput")
		formatted, isValid, err := jsonService.Format(inputJSON)

		redirectData := map[string]string{
			"InputJSON": inputJSON,
			"IsValid":   strconv.FormatBool(isValid),
		}

		if err != nil {
			redirectData["Error"] = err.Error()
			redirectData["Result"] = inputJSON
		} else {
			redirectData["Result"] = formatted
		}
		redirectToPageWithData(w, r, redirectData)
		return
	}

	view.Render(w, r, "json.html", &view.PageData{
		Title:            "JSON Validator & Formatter",
		ToolSpecificData: data,
	})
}
