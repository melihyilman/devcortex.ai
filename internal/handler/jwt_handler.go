package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
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
		ToolSpecificData: make(map[string]interface{}),
	}

	if r.Method == http.MethodPost {
		tokenString := r.FormValue("jwtToken")

		token, err := jwt.Parse(tokenString, nil)
		if err != nil {
			data.ToolSpecificData = map[string]interface{}{"Error": "Token parse error: " + err.Error()}
		} else if token != nil {
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok {
				headerJSON, _ := json.MarshalIndent(token.Header, "", "  ")
				claimsJSON, _ := json.MarshalIndent(claims, "", "  ")
				data.ToolSpecificData = map[string]interface{}{"Header": string(headerJSON), "Payload": string(claimsJSON)}
			} else {
				data.ToolSpecificData = map[string]interface{}{"Error": "Could not read claims."}
			}
		} else {
			data.ToolSpecificData = map[string]interface{}{"Error": "Invalid token."}
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