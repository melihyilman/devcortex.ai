package handler

import (
	"net/http"
	"strings"
	"unicode"

	"devcortex.ai/internal/view"
)

func CaseConverterTool(w http.ResponseWriter, r *http.Request) {
	pageData := &view.PageData{
		Title: "Case Converter",
		ToolSpecificData: map[string]interface{}{
			"InputText": r.URL.Query().Get("inputText"),
			"CaseType":  r.URL.Query().Get("caseType"),
			"Result":    r.URL.Query().Get("result"),
		},
	}

	if r.Method == http.MethodPost {
		inputText := r.FormValue("inputText")
		caseType := r.FormValue("caseType")
		var result string

		switch caseType {
		case "snake":
			result = toSnakeCase(inputText)
		case "camel":
			result = toCamelCase(inputText)
		case "pascal":
			result = toPascalCase(inputText)
		default:
			result = inputText
		}

		data := map[string]string{
			"inputText": inputText,
			"caseType":  caseType,
			"result":    result,
		}
		redirectToPageWithData(w, r, data)
		return
	}

	view.Render(w, r, "case-converter.html", pageData)
}

func toSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 && result[len(result)-1] != '_' {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

func toCamelCase(s string) string {
	s = toPascalCase(s)
	if len(s) == 0 {
		return ""
	}
	return strings.ToLower(string(s[0])) + s[1:]
}

func toPascalCase(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.Title(s)
	s = strings.ReplaceAll(s, " ", "")
	return s
}