package tool

import (
	"encoding/json"
	"fmt"
	"strings"
)

// JsonToCodeConverter defines the interface for the JSON to Code conversion service.
type JsonToCodeConverter interface {
	Convert(jsonString, language string) (string, error)
}

// jsonToCodeConverter is the implementation of the JsonToCodeConverter interface.
type jsonToCodeConverter struct{}

// NewJsonToCodeConverter creates a new JsonToCodeConverter service.
func NewJsonToCodeConverter() JsonToCodeConverter {
	return &jsonToCodeConverter{}
}

// Convert handles the main logic for converting a JSON string to code.
func (s *jsonToCodeConverter) Convert(jsonString, language string) (string, error) {
	var jsonData interface{}
	err := json.Unmarshal([]byte(jsonString), &jsonData)
	if err != nil {
		return "", fmt.Errorf("invalid JSON: %w", err)
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
	default:
		return "", fmt.Errorf("unsupported JSON root type")
	}

	switch language {
	case "csharp":
		return generateCSharpCode(typeDefs, isRootArray, rootElementType), nil
	case "go":
		return generateGoCode(typeDefs, isRootArray, rootElementType), nil
	case "typescript":
		return generateTypeScriptCode(typeDefs, isRootArray, rootElementType), nil
	default:
		return "", fmt.Errorf("unsupported language: %s", language)
	}
}

// TypeDefinition represents a class/struct/interface definition
type TypeDefinition struct {
	Name   string
	Fields []Field
}

// Field represents a field within a type definition
type Field struct {
	Name    string
	Type    string
	IsArray bool
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
			default:
				fieldType = "interface{}"
			}
			td.Fields = append(td.Fields, Field{Name: fieldName, Type: fieldType, IsArray: isArray})
		}
		typeDefs[currentTypeName] = td
	} else if arr, ok := data.([]interface{}); ok && len(arr) > 0 {
		extractTypeDefinitions(arr[0], currentTypeName, typeDefs)
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
	default:
		return "interface{}"
	}
}

func generateCSharpCode(typeDefs map[string]TypeDefinition, isRootArray bool, rootElementType string) string {
	var sb strings.Builder
	if isRootArray {
		sb.WriteString(fmt.Sprintf("public class Root : List<%s> {}\n\n", rootElementType))
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

func generateGoCode(typeDefs map[string]TypeDefinition, isRootArray bool, rootElementType string) string {
	var sb strings.Builder
	if isRootArray {
		sb.WriteString(fmt.Sprintf("type Root []%s\n\n", rootElementType))
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

func generateTypeScriptCode(typeDefs map[string]TypeDefinition, isRootArray bool, rootElementType string) string {
	var sb strings.Builder
	if isRootArray {
		sb.WriteString(fmt.Sprintf("type Root = %s[];\n\n", rootElementType))
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
