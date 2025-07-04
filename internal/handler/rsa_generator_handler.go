package handler

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"net/http"
	"strconv"

	"devcortex.ai/internal/view"
)

func RSAGeneratorTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "RSA Key Pair Generator",
	}

	if r.Method == http.MethodPost {
		bitSizeStr := r.FormValue("bit_size")
		bitSize, err := strconv.Atoi(bitSizeStr)
		if err != nil || (bitSize != 2048 && bitSize != 4096) {
			bitSize = 2048 // Default to 2048 bits
		}

		privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
		if err != nil {
			// Handle error
		}

		// Encode private key to PEM format
		privateKeyPEM := &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		}
		privateKeyStr := string(pem.EncodeToMemory(privateKeyPEM))

		// Encode public key to PEM format
		publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
		if err != nil {
			// Handle error
		}
		publicKeyPEM := &pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: publicKeyBytes,
		}
		publicKeyStr := string(pem.EncodeToMemory(publicKeyPEM))

		data.ToolSpecificData = map[string]interface{}{
			"PrivateKey": privateKeyStr,
			"PublicKey":  publicKeyStr,
			"BitSize":    bitSize,
		}
	}

	view.Render(w, r, "rsa-generator.html", data)
}
