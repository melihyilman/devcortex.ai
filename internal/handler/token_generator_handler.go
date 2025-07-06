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
		
		b := make([]byte, 32)
		if _, err := rand.Read(b); err != nil {
			
			
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
