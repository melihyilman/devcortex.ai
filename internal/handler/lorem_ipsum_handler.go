package handler

import (
	"net/http"
	"strconv"
	"strings"

	"devcortex.ai/internal/view"
	"github.com/drhodes/golorem"
)

func LoremIpsumTool(w http.ResponseWriter, r *http.Request) {
	pageData := &view.PageData{
		Title: "Lorem Ipsum Generator",
		ToolSpecificData: make(map[string]interface{}),
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		paragraphsStr := r.FormValue("paragraphs")
		paragraphs, err := strconv.Atoi(paragraphsStr)
		if err != nil || paragraphs <= 0 {
			paragraphs = 3
		}
		if paragraphs > 20 {
			paragraphs = 20
		}

		var loremText []string
		for i := 0; i < paragraphs; i++ {
			loremText = append(loremText, lorem.Paragraph(3, 5))
		}

		pageData.ToolSpecificData.(map[string]interface{})["Lorem"] = strings.Join(loremText, "\n\n")
		pageData.ToolSpecificData.(map[string]interface{})["Paragraphs"] = paragraphs
	}

	view.Render(w, r, "lorem-ipsum.html", pageData)
}