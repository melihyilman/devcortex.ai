package handler

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"devcortex.ai/internal/view"
)

func JSONToCSVTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "JSON to CSV Converter",
	}

	if r.Method == http.MethodPost {
		jsonInput := r.FormValue("json_input")
		var result string
		var success bool

		var jsonData []map[string]interface{}
		err := json.Unmarshal([]byte(jsonInput), &jsonData)
		if err != nil {
			result = "Error parsing JSON: " + err.Error() + ". Please provide an array of objects."
		} else {
			if len(jsonData) == 0 {
				result = "The JSON array is empty."
			} else {
				// Get headers and sort them for consistent order
				headers := []string{}
				for key := range jsonData[0] {
					headers = append(headers, key)
				}
				sort.Strings(headers)

				var csvBuffer bytes.Buffer
				csvWriter := csv.NewWriter(&csvBuffer)

				// Write header
				if err := csvWriter.Write(headers); err != nil {
					result = "Error writing CSV header: " + err.Error()
				} else {
					// Write rows
					for _, row := range jsonData {
						record := []string{}
						for _, header := range headers {
							value := ""
							if val, ok := row[header]; ok {
								value = fmt.Sprintf("%v", val)
							}
							record = append(record, value)
						}
						if err := csvWriter.Write(record); err != nil {
							result = "Error writing CSV row: " + err.Error()
							break
						}
					}
				}

				if result == "" {
					csvWriter.Flush()
					result = csvBuffer.String()
					success = true
				}
			}
		}

		data.ToolSpecificData = map[string]interface{}{
			"JSONInput": jsonInput,
			"Result":    result,
			"Success":   success,
		}
	}

	view.Render(w, r, "json-to-csv.html", data)
}
