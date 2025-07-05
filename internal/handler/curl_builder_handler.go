package handler

import (
	"fmt"
	"net/http"
	"strings"

	"devcortex.ai/internal/view"
)

func CURLBuilderTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "cURL Command Builder",
	}

	if r.Method == http.MethodPost {
		url := r.FormValue("url")
		method := r.FormValue("method")
		headers := r.FormValue("headers")
		body := r.FormValue("body")

		var curlCmd strings.Builder
		curlCmd.WriteString(fmt.Sprintf("curl -X %s '%s'", method, url))

		if headers != "" {
			headerLines := strings.Split(strings.TrimSpace(headers), "\n")
			for _, line := range headerLines {
				line = strings.TrimSpace(line)
				if line != "" {
					curlCmd.WriteString(fmt.Sprintf(" \\\n  -H '%s'", line))
				}
			}
		}

		if body != "" {
			curlCmd.WriteString(fmt.Sprintf(" \\\n  -d '%s'", body))
		}

		data.ToolSpecificData = map[string]interface{}{
			"CURLCommand": curlCmd.String(),
			"Input": map[string]string{
				"url":     url,
				"method":  method,
				"headers": headers,
				"body":    body,
			},
		}
	}

	view.Render(w, r, "curl-builder.html", data)
}
