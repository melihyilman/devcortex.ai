package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func MimeTypesCheatsheetTool(w http.ResponseWriter, r *http.Request) {
	cheatsheetData := map[string][]Command{
		"Application": {
			{Command: "application/json", Description: "JSON format"},
			{Command: "application/xml", Description: "XML format"},
			{Command: "application/pdf", Description: "Adobe Portable Document Format (PDF)"},
			{Command: "application/zip", Description: "ZIP archive"},
			{Command: "application/octet-stream", Description: "Any kind of binary data"},
		},
		"Text": {
			{Command: "text/plain", Description: "Plain text"},
			{Command: "text/html", Description: "HTML"},
			{Command: "text/css", Description: "Cascading Style Sheets (CSS)"},
			{Command: "text/csv", Description: "Comma-separated values (CSV)"},
			{Command: "text/javascript", Description: "JavaScript"},
		},
		"Image": {
			{Command: "image/jpeg", Description: "JPEG images"},
			{Command: "image/png", Description: "Portable Network Graphics"},
			{Command: "image/gif", Description: "Graphics Interchange Format (GIF)"},
			{Command: "image/svg+xml", Description: "Scalable Vector Graphics (SVG)"},
			{Command: "image/webp", Description: "WEBP image"},
		},
		"Audio": {
			{Command: "audio/mpeg", Description: "MP3 audio"},
			{Command: "audio/wav", Description: "WAV audio"},
			{Command: "audio/ogg", Description: "OGG audio"},
		},
		"Video": {
			{Command: "video/mp4", Description: "MP4 video"},
			{Command: "video/webm", Description: "WEBM video"},
			{Command: "video/quicktime", Description: "QuickTime video"},
		},
	}

	data := &view.PageData{
		Title:            "File Types & MIME Cheatsheet",
		ToolSpecificData: cheatsheetData,
	}
	view.Render(w, r, "mime-types-cheatsheet.html", data)
}
