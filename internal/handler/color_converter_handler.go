package handler

import (
	"fmt"
	"net/http"

	"devcortex.ai/internal/view"
	"github.com/lucasb-eyer/go-colorful"
)

func ColorConverterTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Color Format Converter",
	}

	if r.Method == http.MethodPost {
		colorInput := r.FormValue("color_input")
		var result map[string]string
		var success bool

		c, err := colorful.Hex(colorInput)
		if err != nil {
			// Try parsing as RGB
			var R, G, B float64
			_, err = fmt.Sscanf(colorInput, "rgb(%f, %f, %f)", &R, &G, &B)
			if err != nil {
				// Try parsing as HSL
				var H, S, L float64
				_, err = fmt.Sscanf(colorInput, "hsl(%f, %f, %f)", &H, &S, &L)
				if err != nil {
					result = map[string]string{"Error": "Invalid color format. Use Hex, rgb(r,g,b), or hsl(h,s,l)."}
				} else {
					c = colorful.Hsl(H, S, L)
					success = true
				}
			} else {
				c = colorful.Color{R: R / 255.0, G: G / 255.0, B: B / 255.0}
				success = true
			}
		} else {
			success = true
		}

		if success {
			h, s, l := c.Hsl()
			result = map[string]string{
				"Hex":   c.Hex(),
				"RGB":   fmt.Sprintf("rgb(%.0f, %.0f, %.0f)", c.R*255, c.G*255, c.B*255),
				"HSL":   fmt.Sprintf("hsl(%.0f, %.2f, %.2f)", h, s, l),
				"Color": c.Hex(),
				"Input": colorInput,
			}
		}

		data.ToolSpecificData = result
	}

	view.Render(w, r, "color-converter.html", data)
}
