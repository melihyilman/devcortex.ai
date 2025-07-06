package handler

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"net/http"

	"devcortex.ai/internal/view"
)

func HMACGeneratorTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "HMAC Generator",
		ToolSpecificData: map[string]interface{}{
			"GeneratedHMAC": r.URL.Query().Get("GeneratedHMAC"),
			"Algorithm":     r.URL.Query().Get("Algorithm"),
			"Message":       r.URL.Query().Get("Message"),
			"Secret":        r.URL.Query().Get("Secret"),
		},
	}

	if r.Method == http.MethodPost {
		message := r.FormValue("message")
		secret := r.FormValue("secret")
		algorithm := r.FormValue("algorithm")

		var h func() hash.Hash
		switch algorithm {
		case "sha1":
			h = sha1.New
		case "sha512":
			h = sha512.New
		default:
			h = sha256.New
		}

		mac := hmac.New(h, []byte(secret))
		mac.Write([]byte(message))
		generatedHMAC := hex.EncodeToString(mac.Sum(nil))

		redirectData := map[string]string{
			"GeneratedHMAC": generatedHMAC,
			"Algorithm":     algorithm,
			"Message":       message,
			"Secret":        secret,
		}
		redirectToPageWithData(w, r, redirectData)
		return
	}

	view.Render(w, r, "hmac-generator.html", data)
}
