package handler

import (
	"html/template"
	"net/http"

	"devcortex.ai/internal/tool"
	"devcortex.ai/internal/view"
)

type RegexData struct {
	Pattern         string
	TestString      string
	HighlightedText template.HTML
	Matches         [][]string
	Error           string
}

func RegexTool(w http.ResponseWriter, r *http.Request) {
	data := RegexData{}
	regexService := tool.NewRegexTester()

	if r.Method == http.MethodPost {
		pattern := r.FormValue("pattern")
		testString := r.FormValue("test_string")
		data.Pattern = pattern
		data.TestString = testString

		result, err := regexService.Test(pattern, testString)
		if err != nil {
			data.Error = err.Error()
		} else {
			data.HighlightedText = result.HighlightedText
			data.Matches = result.Matches
		}
	}

	view.Render(w, r, "regex.html", &view.PageData{
		Title:            "Regex Deconstructor",
		ToolSpecificData: data,
	})
}