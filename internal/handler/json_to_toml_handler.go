package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"devcortex.ai/internal/view"
	"github.com/BurntSushi/toml"
)

func JSONToTOMLTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "JSON to TOML Converter",
	}

	if r.Method == http.MethodPost {
		jsonInput := r.FormValue("json_input")
		var result string
		var success bool

		var jsonData interface{}
		err := json.Unmarshal([]byte(jsonInput), &jsonData)
		if err != nil {
			result = "Error parsing JSON: " + err.Error()
		} else {
			var tomlBuffer bytes.Buffer
			encoder := toml.NewEncoder(&tomlBuffer)
			err := encoder.Encode(jsonData)
			if err != nil {
				result = "Error converting to TOML: " + err.Error()
			} else {
				result = tomlBuffer.String()
				success = true
			}
		}

		data.ToolSpecificData = map[string]interface{}{
			"JSONInput": jsonInput,
			"Result":    result,
			"Success":   success,
		}
	}

	view.Render(w, r, "json-to-toml.html", data)
}
