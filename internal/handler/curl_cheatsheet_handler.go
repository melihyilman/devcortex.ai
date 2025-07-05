package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func CurlCheatsheetTool(w http.ResponseWriter, r *http.Request) {
	cheatsheetData := map[string][]Command{
		"Basic Usage": {
			{Command: "curl [url]", Description: "Make a simple GET request to the specified URL."},
			{Command: "curl -o [filename] [url]", Description: "Download a file and save it with the given filename."},
			{Command: "curl -O [url]", Description: "Download a file and save it with its original filename."},
		},
		"HTTP Methods": {
			{Command: "curl -X POST [url]", Description: "Make a POST request."},
			{Command: "curl -X PUT [url]", Description: "Make a PUT request."},
			{Command: "curl -X DELETE [url]", Description: "Make a DELETE request."},
		},
		"Headers & Data": {
			{Command: "curl -H \"Content-Type: application/json\" [url]", Description: "Send a request with a custom HTTP header."},
			{Command: "curl -d \"param1=value1&param2=value2\" [url]", Description: "Send data in a POST request (form-urlencoded)."},
			{Command: "curl -d @data.json -H \"Content-Type: application/json\" [url]", Description: "Send data from a file (e.g., JSON)."},
		},
		"Authentication": {
			{Command: "curl -u [user:password] [url]", Description: "Send a request with Basic Authentication."},
			{Command: "curl -H \"Authorization: Bearer [token]\" [url]", Description: "Send a request with a Bearer Token."},
		},
		"Other Useful Flags": {
			{Command: "curl -i [url]", Description: "Include the HTTP response headers in the output."},
			{Command: "curl -L [url]", Description: "Follow any redirects."},
			{Command: "curl -v [url]", Description: "Verbose output, useful for debugging."},
		},
	}

	data := &view.PageData{
		Title:            "cURL Cheatsheet",
		ToolSpecificData: cheatsheetData,
	}
	view.Render(w, r, "curl-cheatsheet.html", data)
}
