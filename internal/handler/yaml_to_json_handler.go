package handler

import (
	"encoding/json"
	"net/http"

	"devcortex.ai/internal/view"
	"gopkg.in/yaml.v2"
)

func YAMLToJSONTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "YAML to JSON Converter",
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
			// We need to convert it to map[string]interface{} for JSON marshaling.
			convertedData := convertYAMLMap(yamlData)
			jsonData, err := json.MarshalIndent(convertedData, "", "  ")
			if err != nil {
				result = "Error converting to JSON: " + err.Error()
			} else {
				result = string(jsonData)
				success = true
			}
		}

		data.ToolSpecificData = map[string]interface{}{
			"YAMLInput": yamlInput,
			"Result":    result,
			"Success":   success,
		}
	}

	view.Render(w, r, "yaml-to-json.html", data)
}

// convertYAMLMap recursively converts map[interface{}]interface{} to map[string]interface{}
func convertYAMLMap(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convertYAMLMap(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convertYAMLMap(v)
		}
	}
	return i
}
