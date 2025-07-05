package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
	"github.com/nbutton23/zxcvbn-go"
)

func PasswordStrengthTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Password Strength Analyser",
	}

	if r.Method == http.MethodPost {
		password := r.FormValue("password")
		if password != "" {
			result := zxcvbn.PasswordStrength(password, nil)
			data.ToolSpecificData = map[string]interface{}{
				"Result":   result,
				"Password": password,
			}
		}
	}

	view.Render(w, r, "password-strength.html", data)
}
