package handler

import (
	"net/http"
)

// tools artık model.Tools olarak global kullanılacak

// RegisterRoutes, projemizdeki tüm rotaları tanımlar ve verilen mux'a kaydeder.
func RegisterRoutes(mux *http.ServeMux) {
	// Ana sayfa ve statik rotalar
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/tools", Tools)

	// Araç rotaları
	mux.HandleFunc("/tools/base64", Base64Tool)
	mux.HandleFunc("/tools/jwt-decoder", JWTTool)
	mux.HandleFunc("/tools/hash-generator", HashTool)
	mux.HandleFunc("/tools/timestamp-converter", TimestampTool)
	mux.HandleFunc("/tools/json-formatter", JSONTool)
	mux.HandleFunc("/tools/uuid-generator", UUIDTool)
	mux.HandleFunc("/tools/regex-tester", RegexTool)
	mux.HandleFunc("/tools/color-picker", ColorPickerTool)
	mux.HandleFunc("/tools/lorem-ipsum", LoremIpsumTool)
	mux.HandleFunc("/tools/url-encoder", URLEncoderTool)
	mux.HandleFunc("/tools/case-converter", CaseConverterTool)
	mux.HandleFunc("/tools/minifier", MinifierTool)
	mux.HandleFunc("/tools/markdown-previewer", MarkdownPreviewerTool)
	mux.HandleFunc("/tools/diff-checker", DiffCheckerTool)
	mux.HandleFunc("/tools/random-string", RandomStringTool)

	// YENİ BİR ARAÇ EKLEDİĞİNDE, SADECE BU DOSYAYA YENİ BİR SATIR EKLEYECEKSİN.
}
