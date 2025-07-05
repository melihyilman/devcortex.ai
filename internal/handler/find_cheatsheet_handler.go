package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func FindCheatsheetTool(w http.ResponseWriter, r *http.Request) {
	cheatsheetData := map[string][]Command{
		"Basic Find": {
			{Command: "find . -name \"myfile.txt\"", Description: "Find files named 'myfile.txt' in the current directory and subdirectories."},
			{Command: "find /home -name \"*.log\"", Description: "Find all files ending with '.log' in the /home directory."},
			{Command: "find . -iname \"myfile.txt\"", Description: "Case-insensitive search for file names."},
		},
		"Find by Type": {
			{Command: "find . -type d -name \"mydir\"", Description: "Find directories named 'mydir'."},
			{Command: "find . -type f -name \"*.conf\"", Description: "Find files (not directories) ending with '.conf'."},
		},
		"Find by Time": {
			{Command: "find . -mtime -7", Description: "Find files modified in the last 7 days."},
			{Command: "find . -mtime +7", Description: "Find files modified more than 7 days ago."},
			{Command: "find . -cmin -60", Description: "Find files whose status was changed in the last 60 minutes."},
		},
		"Find by Size": {
			{Command: "find . -size +100M", Description: "Find files larger than 100 Megabytes."},
			{Command: "find . -size -10k", Description: "Find files smaller than 10 Kilobytes."},
			{Command: "find . -size 0 -delete", Description: "Find and delete empty files."},
		},
		"Find and Execute": {
			{Command: "find . -name \"*.tmp\" -exec rm {} \\;", Description: "Find all .tmp files and execute the 'rm' command on them."},
			{Command: "find . -type f -name \"*.sh\" -exec chmod +x {} \\;", Description: "Find all .sh files and make them executable."},
		},
	}

	data := &view.PageData{
		Title:            "Linux 'find' Cheatsheet",
		ToolSpecificData: cheatsheetData,
	}
	view.Render(w, r, "find-cheatsheet.html", data)
}
