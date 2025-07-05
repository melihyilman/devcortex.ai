package handler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"net/http"

	"devcortex.ai/internal/view"
)

func EncryptDecryptTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Encrypt/Decrypt Text (AES-GCM)",
	}

	if r.Method == http.MethodPost {
		action := r.FormValue("action")
		text := r.FormValue("text")
		password := r.FormValue("password")
		var result string
		var success bool

		// Derive a 32-byte key from the password using SHA-256
		key := sha256.Sum256([]byte(password))

		if action == "encrypt" {
			encrypted, err := encrypt(text, key[:])
			if err != nil {
				result = "Error encrypting text: " + err.Error()
			} else {
				result = encrypted
				success = true
			}
		} else if action == "decrypt" {
			decrypted, err := decrypt(text, key[:])
			if err != nil {
				result = "Error decrypting text: " + err.Error()
			} else {
				result = decrypted
				success = true
			}
		}

		data.ToolSpecificData = map[string]interface{}{
			"Result":  result,
			"Success": success,
		}
	}

	view.Render(w, r, "encrypt-decrypt.html", data)
}

func encrypt(plaintext string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decrypt(ciphertextB64 string, key []byte) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextB64)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", err
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
