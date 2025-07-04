package main

import (
	"log"
	"net/http"

	"devcortex.ai/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./web/static/"))

	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	handler.RegisterRoutes(mux)

	log.Println("Starting server... Listening on http://localhost:9090")
	if err := http.ListenAndServe(":9090", mux); err != nil {
		log.Fatal("Could not start server: ", err)
	}
}