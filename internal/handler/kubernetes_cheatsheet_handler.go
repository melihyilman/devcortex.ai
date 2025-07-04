package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func KubernetesCheatsheetTool(w http.ResponseWriter, r *http.Request) {
	cheatsheetData := map[string][]Command{
		"Cluster Management": {
			{Command: "kubectl config view", Description: "Display the current context."},
			{Command: "kubectl config use-context [context-name]", Description: "Switch to a different context."},
			{Command: "kubectl cluster-info", Description: "Display information about the cluster."},
			{Command: "kubectl get nodes", Description: "List all nodes in the cluster."},
		},
		"Deploying Applications": {
			{Command: "kubectl apply -f [file.yaml]", Description: "Create or update resources from a YAML file."},
			{Command: "kubectl create deployment [name] --image=[image]", Description: "Create a new deployment."},
			{Command: "kubectl expose deployment [name] --port=[port] --type=LoadBalancer", Description: "Expose a deployment as a new service."},
			{Command: "kubectl scale deployment [name] --replicas=[count]", Description: "Scale the number of pods in a deployment."},
		},
		"Viewing Resources": {
			{Command: "kubectl get pods", Description: "List all pods in the current namespace."},
			{Command: "kubectl get pods -n [namespace]", Description: "List all pods in a specific namespace."},
			{Command: "kubectl get all", Description: "List all resources in the current namespace."},
			{Command: "kubectl describe pod [pod-name]", Description: "Show detailed information about a pod."},
			{Command: "kubectl logs [pod-name]", Description: "Print the logs for a pod."},
			{Command: "kubectl logs -f [pod-name]", Description: "Stream the logs for a pod."},
		},
		"Debugging": {
			{Command: "kubectl exec -it [pod-name] -- /bin/sh", Description: "Execute a shell inside a running pod."},
			{Command: "kubectl port-forward [pod-name] [local-port]:[pod-port]", Description: "Forward a local port to a port on a pod."},
			{Command: "kubectl top pod", Description: "Display CPU and memory usage for pods."},
			{Command: "kubectl get events --sort-by='.lastTimestamp'", Description: "List recent events in the cluster."},
		},
		"Deleting Resources": {
			{Command: "kubectl delete -f [file.yaml]", Description: "Delete resources defined in a YAML file."},
			{Command: "kubectl delete pod [pod-name]", Description: "Delete a pod."},
			{Command: "kubectl delete deployment [deployment-name]", Description: "Delete a deployment."},
			{Command: "kubectl delete service [service-name]", Description: "Delete a service."},
		},
	}

	data := &view.PageData{
		Title:            "Kubernetes Cheatsheet",
		ToolSpecificData: cheatsheetData,
	}
	view.Render(w, r, "kubernetes-cheatsheet.html", data)
}
