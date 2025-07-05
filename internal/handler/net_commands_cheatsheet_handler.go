package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func NetCommandsCheatsheetTool(w http.ResponseWriter, r *http.Request) {
	cheatsheetData := map[string][]Command{
		"Connectivity & Reachability": {
			{Command: "ping [host]", Description: "Checks if a host is reachable and measures round-trip time."},
			{Command: "traceroute [host]", Description: "Traces the path packets take to a network host."},
		},
		"DNS Information": {
			{Command: "nslookup [domain]", Description: "Queries DNS servers to find DNS details (e.g., IP address)."},
			{Command: "dig [domain]", Description: "A more advanced DNS lookup utility."},
		},
		"Network Configuration & Statistics": {
			{Command: "ifconfig", Description: "Displays or configures network interfaces (deprecated on modern Linux)."},
			{Command: "ip addr", Description: "The modern replacement for ifconfig; shows IP addresses and network interface information."},
			{Command: "netstat -tulpn", Description: "Shows active network connections, listening ports, and PIDs."},
			{Command: "ss -tulpn", Description: "The modern replacement for netstat; shows socket statistics."},
		},
		"File Transfer": {
			{Command: "wget [url]", Description: "Downloads files from the network."},
			{Command: "curl -O [url]", Description: "A versatile tool to transfer data with URLs."},
		},
	}

	data := &view.PageData{
		Title:            "Network Commands Cheatsheet",
		ToolSpecificData: cheatsheetData,
	}
	view.Render(w, r, "net-commands-cheatsheet.html", data)
}
