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
		ToolSpecificData: map[string]interface{}{
			"CURLCommand": r.URL.Query().Get("CURLCommand"),
			"Input": map[string]string{
				"url":     r.URL.Query().Get("url"),
				"method":  r.URL.Query().Get("method"),
				"headers": r.URL.Query().Get("headers"),
				"body":    r.URL.Query().Get("body"),
			},
		},
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

		redirectData := map[string]string{
			"CURLCommand": curlCmd.String(),
			"url":         url,
			"method":      method,
			"headers":     headers,
			"body":        body,
		}
		redirectToPageWithData(w, r, redirectData)
		return
	}

	view.Render(w, r, "curl-builder.html", data)
}
