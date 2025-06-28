package view

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"devcortex.ai/internal/model"
)

type PageData struct {
	Title string
	Tools []model.Tool
	Data  interface{}
}

func Render(w http.ResponseWriter, r *http.Request, tmpl string, data *PageData) {
	files := []string{
		"web/template/layout.html",
		filepath.Join("web/template", tmpl),
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Printf("Şablon parse hatası: %v", err)
		http.Error(w, "Internal Server Error", 500)
		return // Hata sonrası fonksiyonu sonlandır
	}

	err = ts.Execute(w, data)
	if err != nil {
		log.Printf("Şablon execute hatası: %v", err)
		http.Error(w, "Internal Server Error", 500)
		return // Hata sonrası fonksiyonu sonlandır
	}
}
