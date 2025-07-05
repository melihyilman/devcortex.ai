package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func SSHCheatsheetTool(w http.ResponseWriter, r *http.Request) {
	cheatsheetData := map[string][]Command{
		"Basic Connection": {
			{Command: "ssh [user]@[host]", Description: "Connect to a remote host as a specific user."},
			{Command: "ssh -p [port] [user]@[host]", Description: "Connect to a host on a specific port."},
		},
		"Authentication": {
			{Command: "ssh-keygen -t rsa -b 4096", Description: "Generate a new RSA SSH key pair."},
			{Command: "ssh-copy-id [user]@[host]", Description: "Copy your public key to a remote host for passwordless login."},
			{Command: "ssh -i /path/to/private_key [user]@[host]", Description: "Connect using a specific private key."},
		},
		"File Transfer (SCP & SFTP)": {
			{Command: "scp /path/to/local/file [user]@[host]:/path/to/remote/", Description: "Copy a file from local to remote."},
			{Command: "scp [user]@[host]:/path/to/remote/file /path/to/local/", Description: "Copy a file from remote to local."},
			{Command: "scp -r /path/to/local/dir [user]@[host]:/path/to/remote/", Description: "Copy a directory recursively from local to remote."},
			{Command: "sftp [user]@[host]", Description: "Start an interactive SFTP session."},
		},
		"Tunneling / Port Forwarding": {
			{Command: "ssh -L [local_port]:[destination_host]:[destination_port] [user]@[gateway_host]", Description: "Local port forwarding. Access a service on a remote network as if it were local."},
			{Command: "ssh -R [remote_port]:[local_host]:[local_port] [user]@[gateway_host]", Description: "Remote port forwarding. Expose a local service to a remote network."},
		},
	}

	data := &view.PageData{
		Title:            "SSH Cheatsheet",
		ToolSpecificData: cheatsheetData,
	}
	view.Render(w, r, "ssh-cheatsheet.html", data)
}
