package handler

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"devcortex.ai/internal/view"
)

type HashData struct {
	InputText string
	Hashes    map[string]string
}

func HashTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Hash Generator",
		ToolSpecificData: make(map[string]interface{}),
	}

	if r.Method == http.MethodPost {
		inputText := []byte(r.FormValue("inputText"))

		hashes := make(map[string]string)

		md5Hasher := md5.New()
		md5Hasher.Write(inputText)
		hashes["MD5"] = hex.EncodeToString(md5Hasher.Sum(nil))

		sha1Hasher := sha1.New()
		sha1Hasher.Write(inputText)
		hashes["SHA-1"] = hex.EncodeToString(sha1Hasher.Sum(nil))

		sha256Hasher := sha256.New()
		sha256Hasher.Write(inputText)
		hashes["SHA-256"] = hex.EncodeToString(sha256Hasher.Sum(nil))

		data.ToolSpecificData = map[string]interface{}{
			"InputText": string(inputText),
			"Hashes":    hashes,
		}
	}

	view.Render(w, r, "hash.html", data)
}
