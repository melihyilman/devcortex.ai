package main

import (
	"log"
	"net/http"
	"os"

	"devcortex.ai/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	handler.RegisterRoutes(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Bilgi: PORT ortam değişkeni bulunamadı, varsayılan olarak %s kullanılıyor.", port)
	}

	finalMux := http.Handler(handler.AnalyticsMiddleware(mux))

	log.Printf("Sunucu http://localhost:%s adresinde başlatılıyor...", port)
	if err := http.ListenAndServe(":"+port, finalMux); err != nil {
		log.Fatal("Could not start server: ", err)
	}
}
