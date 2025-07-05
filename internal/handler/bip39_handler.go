package handler

import (
	"net/http"
	"strconv"

	"devcortex.ai/internal/view"
	"github.com/tyler-smith/go-bip39"
)

func Bip39GeneratorTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "BIP39 Passphrase Generator",
	}

	if r.Method == http.MethodPost {
		wordCountStr := r.FormValue("word_count")
		wordCount, err := strconv.Atoi(wordCountStr)
		if err != nil || !isValidWordCount(wordCount) {
			wordCount = 12 // Default to 12 words if invalid
		}

		// Entropy bits must be a multiple of 32
		// 12 words = 128 bits, 24 words = 256 bits
		entropy, err := bip39.NewEntropy((wordCount / 3) * 32)
		if err != nil {
			data.ToolSpecificData = map[string]interface{}{
				"GeneratedMnemonic": "Error generating mnemonic",
			}
			view.Render(w, r, "bip39-generator.html", data)
			return
		}

		mnemonic, err := bip39.NewMnemonic(entropy)
		if err != nil {
			data.ToolSpecificData = map[string]interface{}{
				"GeneratedMnemonic": "Error generating mnemonic",
			}
			view.Render(w, r, "bip39-generator.html", data)
			return
		}

		data.ToolSpecificData = map[string]interface{}{
			"GeneratedMnemonic": mnemonic,
			"WordCount":         wordCount,
		}
	}

	view.Render(w, r, "bip39-generator.html", data)
}

func isValidWordCount(count int) bool {
	return count == 12 || count == 15 || count == 18 || count == 21 || count == 24
}
