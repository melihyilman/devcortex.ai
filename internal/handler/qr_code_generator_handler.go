package handler

import (
	"encoding/base64"
	"net/http"

	"devcortex.ai/internal/view"
	"github.com/skip2/go-qrcode"
)

func QRCodeGeneratorTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "QR Code Generator",
	}

	if r.Method == http.MethodPost {
		text := r.FormValue("text")
		if text != "" {
			var png []byte
			png, err := qrcode.Encode(text, qrcode.Medium, 256)
			if err != nil {
				data.ToolSpecificData = map[string]interface{}{
					"Error": "Could not generate QR code.",
				}
			} else {
				data.ToolSpecificData = map[string]interface{}{
					"QRCode": base64.StdEncoding.EncodeToString(png),
					"QRText": text,
				}
			}
		}
	}

	view.Render(w, r, "qr-code-generator.html", data)
}
