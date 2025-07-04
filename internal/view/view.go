package view

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"devcortex.ai/internal/model"
)

type PageData struct {
	Title            string
	FeaturedTools    []model.Tool
	OtherTools       []model.Tool
	ToolSpecificData interface{}
}

func Render(w http.ResponseWriter, r *http.Request, tmpl string, data *PageData) {
	data.FeaturedTools = model.FeaturedTools
	data.OtherTools = model.OtherTools

	files := []string{
		"web/template/layout.html",
		filepath.Join("web/template", tmpl),
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Printf("Template parsing error: %v", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}