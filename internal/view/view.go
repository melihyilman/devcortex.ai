package view

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"reflect"
	"sort"

	"devcortex.ai/internal/model"
)

var funcMap = template.FuncMap{
	"add": func(a, b int) int {
		return a + b
	},
	"mul": func(a, b int) int {
		return a * b
	},
	"keys": func(m interface{}) []string {
		v := reflect.ValueOf(m)
		if v.Kind() != reflect.Map {
			return nil
		}
		keys := v.MapKeys()
		strkeys := make([]string, len(keys))
		for i := 0; i < len(keys); i++ {
			strkeys[i] = keys[i].String()
		}
		sort.Strings(strkeys)
		return strkeys
	},
}

type PageData struct {
	Title            string
	FeaturedTools    []model.Tool
	OtherTools       []model.Tool
	ToolSpecificData interface{}
}

func Render(w http.ResponseWriter, r *http.Request, tmpl string, data *PageData) {
	data.FeaturedTools = model.FeaturedTools
	data.OtherTools = model.OtherTools

	templates, err := template.New("").Funcs(funcMap).ParseFiles(
		filepath.Join("web", "template", "layout.html"),
		filepath.Join("web", "template", tmpl),
	)
	if err != nil {
		log.Printf("Error parsing templates: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = templates.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
