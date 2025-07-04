package handler

import (
	"bytes"
	"net/http"

	"devcortex.ai/internal/view"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
)

func MinifierTool(w http.ResponseWriter, r *http.Request) {
	pageData := &view.PageData{
		Title: "HTML/CSS/JS Minifier",
		ToolSpecificData: make(map[string]interface{}),
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		inputText := r.FormValue("inputText")
		minifyType := r.FormValue("minifyType")

		m := minify.New()
		var resultBuffer bytes.Buffer
		var err error

		switch minifyType {
		case "html":
			m.AddFunc("text/html", html.Minify)
			err = m.Minify("text/html", &resultBuffer, bytes.NewBufferString(inputText))
		case "css":
			m.AddFunc("text/css", css.Minify)
			err = m.Minify("text/css", &resultBuffer, bytes.NewBufferString(inputText))
		case "js":
			m.AddFunc("application/javascript", js.Minify)
			err = m.Minify("application/javascript", &resultBuffer, bytes.NewBufferString(inputText))
		}

		if err != nil {
			pageData.ToolSpecificData.(map[string]interface{})["Result"] = "Error: " + err.Error()
		} else {
			pageData.ToolSpecificData.(map[string]interface{})["Result"] = resultBuffer.String()
		}

		pageData.ToolSpecificData.(map[string]interface{})["InputText"] = inputText
		pageData.ToolSpecificData.(map[string]interface{})["MinifyType"] = minifyType
	}

	view.Render(w, r, "minifier.html", pageData)
}