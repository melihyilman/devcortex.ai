package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"devcortex.ai/internal/view"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/svg"
)

func SVGOptimizerTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "SVG Optimizer",
	}

	if r.Method == http.MethodPost {
		file, header, err := r.FormFile("svg_file")
		if err != nil {
			data.ToolSpecificData = map[string]interface{}{"Error": "Could not read uploaded file."}
			view.Render(w, r, "svg-optimizer.html", data)
			return
		}
		defer file.Close()

		// Ensure the uploaded file is an SVG
		if !strings.HasSuffix(strings.ToLower(header.Filename), ".svg") {
			data.ToolSpecificData = map[string]interface{}{"Error": "Invalid file type. Please upload an SVG file."}
			view.Render(w, r, "svg-optimizer.html", data)
			return
		}

		m := minify.New()
		m.AddFunc("image/svg+xml", svg.Minify)

		var outputBuffer bytes.Buffer
		if err := m.Minify("image/svg+xml", &outputBuffer, file); err != nil {
			data.ToolSpecificData = map[string]interface{}{"Error": "Failed to optimize SVG: " + err.Error()}
			view.Render(w, r, "svg-optimizer.html", data)
			return
		}

		originalSize := header.Size
		optimizedSize := int64(outputBuffer.Len())
		reduction := 100 - (float64(optimizedSize) / float64(originalSize) * 100)

		data.ToolSpecificData = map[string]interface{}{
			"OriginalSVG":   "", // We don't show the original
			"OptimizedSVG":  outputBuffer.String(),
			"OriginalSize":  originalSize,
			"OptimizedSize": optimizedSize,
			"Reduction":     fmt.Sprintf("%.2f", reduction),
		}

	}

	view.Render(w, r, "svg-optimizer.html", data)
}
