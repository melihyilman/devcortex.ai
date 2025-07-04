package handler

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/tools", Tools)

	mux.HandleFunc("/tools/regex-deconstructor", RegexTool)
	mux.HandleFunc("/tools/sql-autopsy", SQLAutopsyTool)
	mux.HandleFunc("/tools/base64", Base64Tool)
	mux.HandleFunc("/tools/jwt-decoder", JWTTool)
	mux.HandleFunc("/tools/hash-generator", HashTool)
	mux.HandleFunc("/tools/timestamp-converter", TimestampTool)
		mux.HandleFunc("/tools/json-formatter", JsonFormatterPageHandler)
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
	mux.HandleFunc("/tools/cron-explainer", CronTool)
	// mux.HandleFunc("/tools/sql-formatter", SQLTool) // Temporarily disabled due to module issues

}
