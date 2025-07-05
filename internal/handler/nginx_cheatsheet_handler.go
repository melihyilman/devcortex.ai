package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func NginxCheatsheetTool(w http.ResponseWriter, r *http.Request) {
	cheatsheetData := map[string][]Command{
		"Basic Commands": {
			{Command: "sudo systemctl start nginx", Description: "Start the Nginx service."},
			{Command: "sudo systemctl stop nginx", Description: "Stop the Nginx service."},
			{Command: "sudo systemctl restart nginx", Description: "Restart the Nginx service."},
			{Command: "sudo systemctl reload nginx", Description: "Reload the configuration without dropping connections."},
			{Command: "sudo nginx -t", Description: "Test the Nginx configuration for syntax errors."},
		},
		"Server Block (Virtual Host)": {
			{Command: "server { ... }", Description: "Defines a virtual server to handle requests."},
			{Command: "listen 80;", Description: "Listen for connections on port 80."},
			{Command: "server_name example.com www.example.com;", Description: "Define the server names for this block."},
			{Command: "root /var/www/example.com;", Description: "Set the root directory for requests."},
			{Command: "index index.html index.htm;", Description: "Define the default file to serve."},
		},
		"Location Blocks": {
			{Command: "location / { ... }", Description: "Matches any request."},
			{Command: "location = /exact { ... }", Description: "Matches only the exact path /exact."},
			{Command: "location ~ \\.php$ { ... }", Description: "Matches any request ending in .php (case-sensitive regex)."},
			{Command: "location ~* \\.(gif|jpg|png)$ { ... }", Description: "Matches requests for common image formats (case-insensitive regex)."},
		},
		"Proxying": {
			{Command: "proxy_pass http://backend_server;", Description: "Pass the request to a backend server."},
			{Command: "proxy_set_header Host $host;", Description: "Pass the original host header to the backend."},
			{Command: "proxy_set_header X-Real-IP $remote_addr;", Description: "Pass the client's real IP to the backend."},
		},
		"SSL/TLS (HTTPS)": {
			{Command: "listen 443 ssl;", Description: "Listen for SSL connections on port 443."},
			{Command: "ssl_certificate /path/to/fullchain.pem;", Description: "Path to your SSL certificate."},
			{Command: "ssl_certificate_key /path/to/privkey.pem;", Description: "Path to your private key."},
		},
	}

	data := &view.PageData{
		Title:            "Nginx Cheatsheet",
		ToolSpecificData: cheatsheetData,
	}
	view.Render(w, r, "nginx-cheatsheet.html", data)
}
