package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
	"github.com/clbanning/mxj"
)

func JSONToXMLTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "JSON to XML Converter",
	}

	if r.Method == http.MethodPost {
		jsonInput := r.FormValue("json_input")
		var result string
		var success bool

		mv, err := mxj.NewMapJson([]byte(jsonInput))
		if err != nil {
			result = "Error parsing JSON: " + err.Error()
		} else {
			xmlData, err := mv.XmlIndent("", "  ")
			if err != nil {
				result = "Error converting to XML: " + err.Error()
			} else {
				result = string(xmlData)
				success = true
			}
		}

		data.ToolSpecificData = map[string]interface{}{
			"JSONInput": jsonInput,
			"Result":    result,
			"Success":   success,
		}
	}

	view.Render(w, r, "json-to-xml.html", data)
}
