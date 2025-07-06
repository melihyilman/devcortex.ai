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
	hashes := make(map[string]string)
	hashes["MD5"] = r.URL.Query().Get("md5")
	hashes["SHA-1"] = r.URL.Query().Get("sha1")
	hashes["SHA-256"] = r.URL.Query().Get("sha256")

	data := &view.PageData{
		Title: "Hash Generator",
		ToolSpecificData: map[string]interface{}{
			"InputText": r.URL.Query().Get("inputText"),
			"Hashes":    hashes,
		},
	}

	if r.Method == http.MethodPost {
		inputText := r.FormValue("inputText")
		inputTextBytes := []byte(inputText)

		md5Hasher := md5.New()
		md5Hasher.Write(inputTextBytes)
		md5Hash := hex.EncodeToString(md5Hasher.Sum(nil))

		sha1Hasher := sha1.New()
		sha1Hasher.Write(inputTextBytes)
		sha1Hash := hex.EncodeToString(sha1Hasher.Sum(nil))

		sha256Hasher := sha256.New()
		sha256Hasher.Write(inputTextBytes)
		sha256Hash := hex.EncodeToString(sha256Hasher.Sum(nil))

		redirectData := map[string]string{
			"inputText": inputText,
			"md5":       md5Hash,
			"sha1":      sha1Hash,
			"sha256":    sha256Hash,
		}
		redirectToPageWithData(w, r, redirectData)
		return
	}

	view.Render(w, r, "hash.html", data)
}