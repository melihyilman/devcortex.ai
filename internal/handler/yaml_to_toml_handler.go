package handler

import (
	"bytes"
	"net/http"

	"devcortex.ai/internal/view"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

func YAMLToTOMLTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "YAML to TOML Converter",
	}

	if r.Method == http.MethodPost {
		yamlInput := r.FormValue("yaml_input")
		var result string
		var success bool

		var yamlData interface{}
		err := yaml.Unmarshal([]byte(yamlInput), &yamlData)
		if err != nil {
			result = "Error parsing YAML: " + err.Error()
		} else {
			// The yaml package unmarshals to map[interface{}]interface{}.
			// We need to convert it to map[string]interface{} for TOML marshaling.
			convertedData := convertYAMLMap(yamlData)

			var tomlBuffer bytes.Buffer
			encoder := toml.NewEncoder(&tomlBuffer)
			err := encoder.Encode(convertedData)
			if err != nil {
				result = "Error converting to TOML: " + err.Error()
			} else {
				result = tomlBuffer.String()
				success = true
			}
		}

		data.ToolSpecificData = map[string]interface{}{
			"YAMLInput": yamlInput,
			"Result":    result,
			"Success":   success,
		}
	}

	view.Render(w, r, "yaml-to-toml.html", data)
}
