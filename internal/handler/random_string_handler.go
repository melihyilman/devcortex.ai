package handler

import (
	"crypto/rand"
	"math/big"
	"net/http"
	"strconv"

	"devcortex.ai/internal/view"
)

const ( // Define character sets
	lowerChars  = "abcdefghijklmnopqrstuvwxyz"
	upperChars  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitChars  = "0123456789"
	symbolChars = "!@#$%^&*()_+-=[]{}|;':\",./<>?`~"
)

func RandomStringTool(w http.ResponseWriter, r *http.Request) {
	pageData := &view.PageData{
		Title: "Random String Generator",
		ToolSpecificData: make(map[string]interface{}),
	}

	// Default values for initial load
	pageData.ToolSpecificData.(map[string]interface{})["Length"] = 16
	pageData.ToolSpecificData.(map[string]interface{})["CharsetLower"] = true
	pageData.ToolSpecificData.(map[string]interface{})["CharsetUpper"] = true
	pageData.ToolSpecificData.(map[string]interface{})["CharsetDigit"] = true
	pageData.ToolSpecificData.(map[string]interface{})["CharsetSymbol"] = false

	if r.Method == http.MethodPost {
		lengthStr := r.FormValue("length")
		length, err := strconv.Atoi(lengthStr)
		if err != nil || length < 4 || length > 64 {
			length = 16 // Default to 16 if invalid
		}

		charsets := r.Form["charset"]
		var charset string
		for _, cs := range charsets {
			switch cs {
			case "lower":
				charset += lowerChars
				pageData.ToolSpecificData.(map[string]interface{})["CharsetLower"] = true
			case "upper":
				charset += upperChars
				pageData.ToolSpecificData.(map[string]interface{})["CharsetUpper"] = true
			case "digit":
				charset += digitChars
				pageData.ToolSpecificData.(map[string]interface{})["CharsetDigit"] = true
			case "symbol":
				charset += symbolChars
				pageData.ToolSpecificData.(map[string]interface{})["CharsetSymbol"] = true
			}
		}

		if charset == "" {
			charset = lowerChars + upperChars + digitChars // Default if no charset selected
			pageData.ToolSpecificData.(map[string]interface{})["CharsetLower"] = true
			pageData.ToolSpecificData.(map[string]interface{})["CharsetUpper"] = true
			pageData.ToolSpecificData.(map[string]interface{})["CharsetDigit"] = true
		}

		generatedString, err := generateRandomString(length, charset)
		if err != nil {
			pageData.ToolSpecificData.(map[string]interface{})["Result"] = "Error generating string"
		} else {
			pageData.ToolSpecificData.(map[string]interface{})["Result"] = generatedString
		}
		pageData.ToolSpecificData.(map[string]interface{})["Length"] = length
	}

	view.Render(w, r, "random-string.html", pageData)
}

func generateRandomString(length int, charset string) (string, error) {
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		b[i] = charset[num.Int64()]
	}
	return string(b), nil
}
