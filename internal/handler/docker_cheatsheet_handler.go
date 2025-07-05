package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func DockerCheatsheetTool(w http.ResponseWriter, r *http.Request) {
	cheatsheetData := map[string][]Command{
		"Image Management": {
			{Command: "docker build -t [name:tag] .", Description: "Build an image from a Dockerfile in the current directory."},
			{Command: "docker images", Description: "List all local images."},
			{Command: "docker rmi [image_id]", Description: "Remove a local image."},
			{Command: "docker pull [image]", Description: "Pull an image from a registry (e.g., Docker Hub)."},
			{Command: "docker push [username/image]", Description: "Push an image to a registry."},
		},
		"Container Management": {
			{Command: "docker run [image]", Description: "Create and start a new container from an image."},
			{Command: "docker run -d -p [host_port]:[container_port] [image]", Description: "Run a container in detached mode with port mapping."},
			{Command: "docker ps", Description: "List all running containers."},
			{Command: "docker ps -a", Description: "List all containers (running and stopped)."},
			{Command: "docker stop [container_id]", Description: "Stop a running container."},
			{Command: "docker start [container_id]", Description: "Start a stopped container."},
			{Command: "docker rm [container_id]", Description: "Remove a stopped container."},
			{Command: "docker logs [container_id]", Description: "Fetch the logs of a container."},
			{Command: "docker exec -it [container_id] [command]", Description: "Execute a command inside a running container (e.g., /bin/bash)."},
		},
		"Docker Compose": {
			{Command: "docker-compose up", Description: "Create and start containers defined in docker-compose.yml."},
			{Command: "docker-compose up -d", Description: "Create and start containers in detached mode."},
			{Command: "docker-compose down", Description: "Stop and remove containers, networks, and volumes."},
			{Command: "docker-compose ps", Description: "List containers for the current project."},
			{Command: "docker-compose logs -f [service_name]", Description: "Follow log output for a specific service."},
		},
		"System": {
			{Command: "docker system prune", Description: "Remove all unused containers, networks, images (both dangling and unreferenced)."},
			{Command: "docker system prune -a", Description: "Remove all unused images, not just dangling ones."},
			{Command: "docker info", Description: "Display system-wide information."},
		},
	}

	data := &view.PageData{
		Title:            "Docker Cheatsheet",
		ToolSpecificData: cheatsheetData,
	}
	view.Render(w, r, "docker-cheatsheet.html", data)
}
