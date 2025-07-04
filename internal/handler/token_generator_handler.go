package handler

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"devcortex.ai/internal/view"
)

func TokenGeneratorTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Token Generator",
	}

	if r.Method == http.MethodPost {
		// Generate a 32-byte random token
		b := make([]byte, 32)
		if _, err := rand.Read(b); err != nil {
			// Handle error
			// For simplicity, we'll just show an error in the token field
			data.ToolSpecificData = map[string]interface{}{
				"GeneratedToken": "Error generating token",
			}
			view.Render(w, r, "token-generator.html", data)
			return
		}
		token := hex.EncodeToString(b)

		data.ToolSpecificData = map[string]interface{}{
			"GeneratedToken": token,
		}
	}

	view.Render(w, r, "token-generator.html", data)
}
