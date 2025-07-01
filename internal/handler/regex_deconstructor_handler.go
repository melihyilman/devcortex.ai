package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func RegexDeconstructorTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Regex Dekonstrüktörü",
		ToolSpecificData: make(map[string]interface{}),
	}

	// Placeholder for regex deconstruction logic
	// In a real application, you would parse the regex and extract its components.
	data.ToolSpecificData.(map[string]interface{})["RegexPattern"] = `^([a-zA-Z0-9._%+-]+)@([a-zA-Z0-9.-]+\.[a-zA-Z]{2,})$`
	data.ToolSpecificData.(map[string]interface{})["Explanation"] = "Bu regex, bir e-posta adresini doğrular."

	view.Render(w, r, "regex-deconstructor.html", data)
}