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
		ToolSpecificData: map[string]interface{}{
			"Input": r.URL.Query().Get("Input"),
			"Hex":   r.URL.Query().Get("Hex"),
			"RGB":   r.URL.Query().Get("RGB"),
			"HSL":   r.URL.Query().Get("HSL"),
			"Error": r.URL.Query().Get("Error"),
			"Color": r.URL.Query().Get("Hex"),
		},
	}

	if r.Method == http.MethodPost {
		colorInput := r.FormValue("color_input")
		redirectData := map[string]string{
			"Input": colorInput,
		}
		var success bool
		var c colorful.Color

		c, err := colorful.Hex(colorInput)
		if err != nil {
			var R, G, B float64
			_, err = fmt.Sscanf(colorInput, "rgb(%f, %f, %f)", &R, &G, &B)
			if err != nil {
				var H, S, L float64
				_, err = fmt.Sscanf(colorInput, "hsl(%f, %f, %f)", &H, &S, &L)
				if err != nil {
					redirectData["Error"] = "Invalid color format. Use Hex, rgb(r,g,b), or hsl(h,s,l)."
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
			redirectData["Hex"] = c.Hex()
			redirectData["RGB"] = fmt.Sprintf("rgb(%.0f, %.0f, %.0f)", c.R*255, c.G*255, c.B*255)
			redirectData["HSL"] = fmt.Sprintf("hsl(%.0f, %.2f, %.2f)", h, s, l)
		}

		redirectToPageWithData(w, r, redirectData)
		return
	}

	view.Render(w, r, "color-converter.html", data)
}
