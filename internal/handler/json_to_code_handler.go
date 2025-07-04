package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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

type Field struct {
	Name    string
	Type    string
	IsArray bool
}

type TypeDefinition struct {
	Name   string
	Fields []Field
}

func JsonToCode(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req JsonToCodeRequest
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(body, &req)
		if err != nil {
			http.Error(w, "Error parsing request body", http.StatusBadRequest)
			return
		}

		var jsonData interface{}
		err = json.Unmarshal([]byte(req.Json), &jsonData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(JsonToCodeResponse{Error: fmt.Sprintf("Invalid JSON: %v", err)})
			return
		}

		typeDefs := make(map[string]TypeDefinition)
		rootTypeName := "Root"
		isRootArray := false
		rootElementType := ""

		switch v := jsonData.(type) {
		case []interface{}:
			isRootArray = true
			if len(v) > 0 {
				switch elem := v[0].(type) {
				case map[string]interface{}:
					extractTypeDefinitions(elem, rootTypeName, typeDefs)
					rootElementType = rootTypeName
				default:
					rootElementType = getBaseType(elem)
				}
			} else {
				rootElementType = "interface{}"
			}
		case map[string]interface{}:
			extractTypeDefinitions(v, rootTypeName, typeDefs)
		case string, float64, bool:
			generatedCode := ""
			switch req.Language {
			case "csharp":
				generatedCode = getCSharpType(jsonData)
			case "go":
				generatedCode = getGoType(jsonData)
			case "typescript":
				generatedCode = getTypeScriptType(jsonData)
			}
			json.NewEncoder(w).Encode(JsonToCodeResponse{Code: generatedCode})
			return
		default:
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(JsonToCodeResponse{Error: "Unsupported JSON root type."})
			return
		}

		generatedCode := ""
		switch req.Language {
		case "csharp":
			generatedCode = generateCSharpCode(typeDefs, isRootArray, rootElementType)
		case "go":
			generatedCode = generateGoCode(typeDefs, isRootArray, rootElementType)
		case "typescript":
			generatedCode = generateTypeScriptCode(typeDefs, isRootArray, rootElementType)
		default:
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(JsonToCodeResponse{Error: "Unsupported language"})
			return
		}

		json.NewEncoder(w).Encode(JsonToCodeResponse{Code: generatedCode})
		return
	}

	data := &view.PageData{
		Title: "JSON to Code",
	}
	view.Render(w, r, "json-to-code.html", data)
}

func extractTypeDefinitions(data interface{}, currentTypeName string, typeDefs map[string]TypeDefinition) {
	if m, ok := data.(map[string]interface{}); ok {
		td := TypeDefinition{Name: currentTypeName}
		for key, value := range m {
			fieldName := key
			fieldType := ""
			isArray := false

			switch v := value.(type) {
			case string:
				fieldType = "string"
			case float64:
				if v == float64(int(v)) {
					fieldType = "int"
				} else {
					fieldType = "float64"
				}
			case bool:
				fieldType = "bool"
			case []interface{}:
				isArray = true
				if len(v) > 0 {
					nestedTypeName := currentTypeName + strings.Title(key) + "Element"
					extractTypeDefinitions(v[0], nestedTypeName, typeDefs)
					if _, exists := typeDefs[nestedTypeName]; exists {
						fieldType = nestedTypeName
					} else {
						fieldType = getBaseType(v[0])
					}
				} else {
					fieldType = "interface{}"
				}
			case map[string]interface{}:
				nestedTypeName := currentTypeName + strings.Title(key)
				extractTypeDefinitions(v, nestedTypeName, typeDefs)
				fieldType = nestedTypeName
			case nil:
				fieldType = "object"
			default:
				fieldType = "interface{}"
			}
			td.Fields = append(td.Fields, Field{Name: fieldName, Type: fieldType, IsArray: isArray})
		}
		typeDefs[currentTypeName] = td
	} else if arr, ok := data.([]interface{}); ok && len(arr) > 0 {
		extractTypeDefinitions(arr[0], currentTypeName, typeDefs)
	} else {
	}
}

func getBaseType(value interface{}) string {
	switch v := value.(type) {
	case string:
		return "string"
	case float64:
		if v == float64(int(v)) {
			return "int"
		}
		return "float64"
	case bool:
		return "bool"
	case nil:
		return "interface{}"
	default:
		return "interface{}"
	}
}

func getCSharpType(value interface{}) string {
	switch v := value.(type) {
	case string:
		return "string"
	case float64:
		if v == float64(int(v)) {
			return "int"
		}
		return "double"
	case bool:
		return "bool"
	case nil:
		return "object"
	default:
		return "object"
	}
}

func generateCSharpCode(typeDefs map[string]TypeDefinition, isRootArray bool, rootElementType string) string {
	var sb strings.Builder

	if isRootArray {
		sb.WriteString(fmt.Sprintf("using System.Collections.Generic;\n\n"))
		if rootElementType != "" && rootElementType != "interface{}" {
			sb.WriteString(fmt.Sprintf("public class Root : List<%s> {}\n\n", rootElementType))
		} else {
			sb.WriteString("public class Root : List<object> {}\n\n")
		}
	}

	for _, td := range typeDefs {
		sb.WriteString(fmt.Sprintf("public class %s\n{\n", td.Name))
		for _, field := range td.Fields {
			csharpType := field.Type
			if field.IsArray {
				csharpType = fmt.Sprintf("List<%s>", field.Type)
			}
			pascalKey := strings.ToUpper(field.Name[:1]) + field.Name[1:]
			sb.WriteString(fmt.Sprintf("    public %s %s { get; set; }\n", csharpType, pascalKey))
		}
		sb.WriteString("}\n\n")
	}
	return sb.String()
}

func getGoType(value interface{}) string {
	switch v := value.(type) {
	case string:
		return "string"
	case float64:
		if v == float64(int(v)) {
			return "int"
		}
		return "float64"
	case bool:
		return "bool"
	case nil:
		return "interface{}"
	default:
		return "interface{}"
	}
}

func generateGoCode(typeDefs map[string]TypeDefinition, isRootArray bool, rootElementType string) string {
	var sb strings.Builder

	if isRootArray {
		if rootElementType != "" && rootElementType != "interface{}" {
			sb.WriteString(fmt.Sprintf("type Root []%s\n\n", rootElementType))
		} else {
			sb.WriteString("type Root []interface{}\n\n")
		}
	}

	for _, td := range typeDefs {
		sb.WriteString(fmt.Sprintf("type %s struct {\n", td.Name))
		for _, field := range td.Fields {
			goType := field.Type
			if field.IsArray {
				goType = fmt.Sprintf("[]%s", field.Type)
			}
			pascalKey := strings.ToUpper(field.Name[:1]) + field.Name[1:]
			sb.WriteString(fmt.Sprintf("    %s %s `json:\"%s\"`\n", pascalKey, goType, field.Name))
		}
		sb.WriteString("}\n\n")
	}
	return sb.String()
}

func getTypeScriptType(value interface{}) string {
	switch value.(type) {
	case string:
		return "string"
	case float64:
		return "number"
	case bool:
		return "boolean"
	case nil:
		return "any"
	default:
		return "any"
	}
}

func generateTypeScriptCode(typeDefs map[string]TypeDefinition, isRootArray bool, rootElementType string) string {
	var sb strings.Builder

	if isRootArray {
		if rootElementType != "" && rootElementType != "interface{}" {
			sb.WriteString(fmt.Sprintf("type Root = %s[];\n\n", rootElementType))
		} else {
			sb.WriteString("type Root = any[];\n\n")
		}
	}

	for _, td := range typeDefs {
		sb.WriteString(fmt.Sprintf("interface %s {\n", td.Name))
		for _, field := range td.Fields {
			typeScriptType := field.Type
			if field.IsArray {
				typeScriptType = fmt.Sprintf("%s[]", field.Type)
			}
			sb.WriteString(fmt.Sprintf("    %s: %s;\n", field.Name, typeScriptType))
		}
		sb.WriteString("}\n\n")
	}
	return sb.String()
}
