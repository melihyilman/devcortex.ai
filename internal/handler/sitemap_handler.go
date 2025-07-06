package handler

import (
	"fmt"
	"net/http"
	"time"

	"devcortex.ai/internal/model"
)

const (
	baseURL    = "https://devcortex.ai"
	timeFormat = "2006-01-02"
)

// SitemapHandler generates and serves the sitemap.xml file.
func SitemapHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")

	fmt.Fprintln(w, `<?xml version="1.0" encoding="UTF-8"?>`)
	fmt.Fprintln(w, `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`)

	today := time.Now().Format(timeFormat)

	addURL(w, "/", today, "daily", "1.0")
	addURL(w, "/tools", today, "weekly", "0.9")

	for _, t := range model.FeaturedTools {
		addURL(w, t.URL, today, "monthly", "0.8")
	}

	for _, tool := range model.OtherTools {
		addURL(w, tool.URL, today, "monthly", "0.8")
	}

	fmt.Fprintln(w, `</urlset>`)
}

func addURL(w http.ResponseWriter, path, lastmod, changefreq, priority string) {
	fmt.Fprintf(w, "  <url>\n    <loc>%s%s</loc>\n    <lastmod>%s</lastmod>\n    <changefreq>%s</changefreq>\n    <priority>%s</priority>\n  </url>\n", baseURL, path, lastmod, changefreq, priority)
}
