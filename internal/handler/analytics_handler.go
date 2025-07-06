package handler

import (
	"net/http"
	"sort"
	"sync"
	"html/template"
)

var (
	visitCounts = make(map[string]int)
	mutex       = &sync.Mutex{}
)

func AnalyticsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/favicon.ico" {
			mutex.Lock()
			visitCounts[r.URL.Path]++
			mutex.Unlock()
		}
		next.ServeHTTP(w, r)
	})
}

func AnalyticsHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	tmpl, err := template.ParseFiles("web/template/layout.html", "web/template/analytics.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	
	type pageVisit struct {
		Path  string
		Count int
	}
	var sortedVisits []pageVisit
	for path, count := range visitCounts {
		sortedVisits = append(sortedVisits, pageVisit{path, count})
	}

	
	sort.Slice(sortedVisits, func(i, j int) bool {
		return sortedVisits[i].Count > sortedVisits[j].Count
	})

	if err := tmpl.ExecuteTemplate(w, "layout", sortedVisits); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
