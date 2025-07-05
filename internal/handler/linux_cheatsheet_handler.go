package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func LinuxCheatsheetTool(w http.ResponseWriter, r *http.Request) {
	cheatsheetData := map[string][]Command{
		"File Commands": {
			{Command: "ls", Description: "List directory contents."},
			{Command: "ls -la", Description: "List all files (including hidden) in long format."},
			{Command: "cd [dir]", Description: "Change the current directory."},
			{Command: "pwd", Description: "Show the present working directory."},
			{Command: "mkdir [dir]", Description: "Create a new directory."},
			{Command: "rm [file]", Description: "Remove a file."},
			{Command: "rm -r [dir]", Description: "Remove a directory and its contents recursively."},
			{Command: "cp [file1] [file2]", Description: "Copy a file."},
			{Command: "mv [file1] [file2]", Description: "Move or rename a file."},
			{Command: "touch [file]", Description: "Create a new empty file or update the timestamp of an existing file."},
			{Command: "cat [file]", Description: "Display the contents of a file."},
			{Command: "head [file]", Description: "Display the first 10 lines of a file."},
			{Command: "tail [file]", Description: "Display the last 10 lines of a file."},
		},
		"Process Management": {
			{Command: "ps", Description: "Display currently active processes."},
			{Command: "top", Description: "Display all running processes in real-time."},
			{Command: "kill [pid]", Description: "Terminate a process with the given process ID (pid)."},
			{Command: "killall [proc]", Description: "Kill all processes named [proc]."},
		},
		"File Permissions": {
			{Command: "chmod [permissions] [file]", Description: "Change the permissions of a file (e.g., chmod 755 myfile)."},
			{Command: "chown [user:group] [file]", Description: "Change the ownership of a file."},
		},
		"Searching": {
			{Command: "grep [pattern] [file]", Description: "Search for a pattern within a file."},
			{Command: "find [dir] -name [filename]", Description: "Search for files in a directory."},
		},
		"System Information": {
			{Command: "df", Description: "Display disk space usage."},
			{Command: "du", Description: "Display directory space usage."},
			{Command: "free", Description: "Display memory and swap usage."},
			{Command: "uname -a", Description: "Show kernel and system information."},
		},
		"Networking": {
			{Command: "ping [host]", Description: "Check network connectivity to a host."},
			{Command: "ifconfig", Description: "Display network interface information (deprecated, use 'ip addr')."},
			{Command: "ip addr", Description: "Display network interface information."},
			{Command: "netstat -tulpn", Description: "Show all listening ports and services."},
		},
	}

	data := &view.PageData{
		Title:            "Linux Cheatsheet",
		ToolSpecificData: cheatsheetData,
	}
	view.Render(w, r, "linux-cheatsheet.html", data)
}
