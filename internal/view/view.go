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
	Description      string
	CanonicalURL     string
	FeaturedTools    []model.Tool
	OtherTools       []model.Tool
	ToolSpecificData interface{}
}

func Render(w http.ResponseWriter, r *http.Request, tmpl string, data *PageData) {
	// Populate common data for all pages
	data.FeaturedTools = model.FeaturedTools
	data.OtherTools = model.OtherTools

	// Construct Canonical URL. In a real production environment, you'd get the
	// scheme and host from a config file to avoid host header injection vulnerabilities.
	scheme := "http"
	if r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https" {
		scheme = "https"
	}
	data.CanonicalURL = scheme + "://" + r.Host + r.URL.Path

	// Set a default description if one isn't provided by the handler.
	if data.Description == "" {
		data.Description = "A collection of smart, fast, and free developer tools, including JSON Formatter, JWT Debugger, Base64 Encoder, Hash Generator, and P2P Scrum Poker."
	}

	// Parse templates
	templates, err := template.New("").Funcs(funcMap).ParseFiles(
		filepath.Join("web", "template", "layout.html"),
		filepath.Join("web", "template", tmpl),
	)
	if err != nil {
		log.Printf("Error parsing templates: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the template
	err = templates.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
