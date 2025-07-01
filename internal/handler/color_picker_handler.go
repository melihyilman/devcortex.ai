package handler

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"devcortex.ai/internal/view"
)

func ColorPickerTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Color Picker / Converter",
		ToolSpecificData: make(map[string]interface{}),
	}

	if r.Method == http.MethodPost {
		color := r.FormValue("color")
		hex := r.FormValue("hex")

		if hex != "" {
			color = hex
		}

		// Remove # if present
		color = strings.TrimPrefix(color, "#")

		if len(color) == 6 {
			r, _ := strconv.ParseInt(color[0:2], 16, 0)
			g, _ := strconv.ParseInt(color[2:4], 16, 0)
			b, _ := strconv.ParseInt(color[4:6], 16, 0)

			data.ToolSpecificData.(map[string]interface{})["Hex"] = "#" + color
			data.ToolSpecificData.(map[string]interface{})["RGB"] = fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
			data.ToolSpecificData.(map[string]interface{})["HSL"] = rgbToHsl(int(r), int(g), int(b))
		}
	}

	view.Render(w, r, "color-picker.html", data)
}

func rgbToHsl(r, g, b int) string {
	rf := float64(r) / 255
	gf := float64(g) / 255
	bf := float64(b) / 255

	max := max(rf, gf, bf)
	min := min(rf, gf, bf)

	h := 0.0
	s := 0.0
	l := (max + min) / 2

	if max == min {
		h = 0
		s = 0
	} else {
		d := max - min
		s = 0.0
		if l > 0.5 {
			s = d / (2 - max - min)
		} else {
			s = d / (max + min)
		}

		switch max {
		case rf:
			h = (gf - bf) / d
			if gf < bf {
				h += 6
			}
		case gf:
			h = (bf - rf) / d + 2
		case bf:
			h = (rf - gf) / d + 4
		}
		h /= 6
	}

	return fmt.Sprintf("hsl(%.0f, %.0f%%, %.0f%%)", h*360, s*100, l*100)
}

func max(a, b, c float64) float64 {
	return math.Max(a, math.Max(b, c))
}

func min(a, b, c float64) float64 {
	return math.Min(a, math.Min(b, c))
}
