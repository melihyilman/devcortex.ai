package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func HTTPStatusCodesCheatsheetTool(w http.ResponseWriter, r *http.Request) {
	cheatsheetData := map[string][]Command{
		"1xx Informational": {
			{Command: "100 Continue", Description: "The server has received the request headers and the client should proceed to send the request body."},
			{Command: "101 Switching Protocols", Description: "The server is switching protocols as requested by the client."},
		},
		"2xx Success": {
			{Command: "200 OK", Description: "The request has succeeded."},
			{Command: "201 Created", Description: "The request has been fulfilled and has resulted in one or more new resources being created."},
			{Command: "202 Accepted", Description: "The request has been accepted for processing, but the processing has not been completed."},
			{Command: "204 No Content", Description: "The server successfully processed the request but is not returning any content."},
		},
		"3xx Redirection": {
			{Command: "301 Moved Permanently", Description: "The requested resource has been assigned a new permanent URI."},
			{Command: "302 Found", Description: "The requested resource resides temporarily under a different URI (HTTP 1.0)."},
			{Command: "304 Not Modified", Description: "Indicates that the resource has not been modified since the version specified by the request headers."},
		},
		"4xx Client Error": {
			{Command: "400 Bad Request", Description: "The server cannot or will not process the request due to a client error."},
			{Command: "401 Unauthorized", Description: "Authentication is required and has failed or has not yet been provided."},
			{Command: "403 Forbidden", Description: "The server understood the request but refuses to authorize it."},
			{Command: "404 Not Found", Description: "The server has not found anything matching the Request-URI."},
			{Command: "429 Too Many Requests", Description: "The user has sent too many requests in a given amount of time."},
		},
		"5xx Server Error": {
			{Command: "500 Internal Server Error", Description: "The server encountered an unexpected condition that prevented it from fulfilling the request."},
			{Command: "502 Bad Gateway", Description: "The server, while acting as a gateway or proxy, received an invalid response from an inbound server."},
			{Command: "503 Service Unavailable", Description: "The server is currently unable to handle the request due to a temporary overload or maintenance."},
			{Command: "504 Gateway Timeout", Description: "The server, while acting as a gateway or proxy, did not receive a timely response from an upstream server."},
		},
	}

	data := &view.PageData{
		Title:            "HTTP Status Codes Cheatsheet",
		ToolSpecificData: cheatsheetData,
	}
	view.Render(w, r, "http-status-codes.html", data)
}
