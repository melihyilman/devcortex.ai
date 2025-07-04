package handler

import (
	"encoding/json"
	"net/http"

	"devcortex.ai/internal/tool"
	"devcortex.ai/internal/view"
)

type JsonToCodeRequest struct {
	Json     string `json:"json"`
	Language string `json:"language"`
}

type JsonToCodeResponse struct {
	Code  string `json:"code"`
	Error string `json:"error,omitempty"`
}

func JsonToCode(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data := &view.PageData{
			Title: "JSON to Code",
		}
		view.Render(w, r, "json-to-code.html", data)
		return
	}

	if r.Method == http.MethodPost {
		var req JsonToCodeRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Error parsing request body", http.StatusBadRequest)
			return
		}

		converter := tool.NewJsonToCodeConverter()
		generatedCode, err := converter.Convert(req.Json, req.Language)

		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(JsonToCodeResponse{Error: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(JsonToCodeResponse{Code: generatedCode})
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}