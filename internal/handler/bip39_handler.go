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
		ToolSpecificData: map[string]interface{}{
			"GeneratedMnemonic": r.URL.Query().Get("GeneratedMnemonic"),
			"WordCount":         r.URL.Query().Get("WordCount"),
		},
	}

	if r.Method == http.MethodPost {
		wordCountStr := r.FormValue("word_count")
		wordCount, err := strconv.Atoi(wordCountStr)
		if err != nil || !isValidWordCount(wordCount) {
			wordCount = 12
		}

		var mnemonic string
		entropy, err := bip39.NewEntropy((wordCount / 3) * 32)
		if err != nil {
			mnemonic = "Error generating mnemonic"
		} else {
			mnemonic, err = bip39.NewMnemonic(entropy)
			if err != nil {
				mnemonic = "Error generating mnemonic"
			}
		}

		redirectData := map[string]string{
			"GeneratedMnemonic": mnemonic,
			"WordCount":         strconv.Itoa(wordCount),
		}
		redirectToPageWithData(w, r, redirectData)
		return
	}

	view.Render(w, r, "bip39-generator.html", data)
}

func isValidWordCount(count int) bool {
	return count == 12 || count == 15 || count == 18 || count == 21 || count == 24
}
