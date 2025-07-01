package handler

import (
	"net/http"
	"regexp"

	"devcortex.ai/internal/view"
)

func RegexTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Regex Tester",
		ToolSpecificData: make(map[string]interface{}),
	}

	if r.Method == http.MethodPost {
		pattern := r.FormValue("pattern")
		inputText := r.FormValue("inputText")

		data.ToolSpecificData.(map[string]interface{})["Pattern"] = pattern
		data.ToolSpecificData.(map[string]interface{})["InputText"] = inputText

		re, err := regexp.Compile(pattern)
		if err != nil {
			data.ToolSpecificData.(map[string]interface{})["Error"] = "Invalid regex pattern: " + err.Error()
		} else {
			matches := re.FindAllString(inputText, -1)
			data.ToolSpecificData.(map[string]interface{})["Matches"] = matches
		}
	}

	view.Render(w, r, "regex-tester.html", data)
}
