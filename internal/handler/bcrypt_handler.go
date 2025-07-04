package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
	"golang.org/x/crypto/bcrypt"
)

func BcryptTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Bcrypt Hash & Compare",
	}

	if r.Method == http.MethodPost {
		action := r.FormValue("action")
		var result string
		var success bool

		if action == "hash" {
			textToHash := r.FormValue("text_to_hash")
			hashed, err := bcrypt.GenerateFromPassword([]byte(textToHash), bcrypt.DefaultCost)
			if err != nil {
				result = "Error hashing text."
			} else {
				result = string(hashed)
				success = true
			}
		} else if action == "compare" {
			textToCompare := r.FormValue("text_to_compare")
			hashToCompare := r.FormValue("hash_to_compare")
			err := bcrypt.CompareHashAndPassword([]byte(hashToCompare), []byte(textToCompare))
			if err == nil {
				result = "The text and hash match."
				success = true
			} else {
				result = "The text and hash do not match."
			}
		}

		data.ToolSpecificData = map[string]interface{}{
			"Result":  result,
			"Success": success,
		}
	}

	view.Render(w, r, "bcrypt-hash.html", data)
}
