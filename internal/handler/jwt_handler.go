package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	"devcortex.ai/internal/view"
)

type JWTData struct {
	InputToken string
	Header     string
	Payload    string
	Error      string
}

func JWTTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "JWT Decoder",
		Data:  JWTData{},
	}

	if r.Method == http.MethodPost {
		tokenStr := r.FormValue("jwtToken")

		parts := strings.Split(tokenStr, ".")
		if len(parts) != 3 {
			data.Data = JWTData{InputToken: tokenStr, Error: "Geçersiz JWT formatı. Token 3 bölümden oluşmalıdır."}
		} else {
			header, err1 := decodeJWTPart(parts[0])
			payload, err2 := decodeJWTPart(parts[1])

			if err1 != nil || err2 != nil {
				data.Data = JWTData{InputToken: tokenStr, Error: "Token'ın Header veya Payload kısmı decode edilemedi. Geçersiz Base64."}
			} else {
				data.Data = JWTData{InputToken: tokenStr, Header: header, Payload: payload}
			}
		}
	}

	view.Render(w, r, "jwt.html", data)
}

func decodeJWTPart(part string) (string, error) {
	if l := len(part) % 4; l > 0 {
		part += strings.Repeat("=", 4-l)
	}

	decoded, err := base64.URLEncoding.DecodeString(part)
	if err != nil {
		return "", err
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, decoded, "", "  ")
	if err != nil {
		return string(decoded), nil
	}

	return prettyJSON.String(), nil
}
