package tool

import (
	"bytes"
	"encoding/json"
	"fmt"
)


type JSONFormatter interface {
	Format(rawJSON string) (string, bool, error)
}


type jsonFormatter struct{}


func NewJSONFormatter() JSONFormatter {
	return &jsonFormatter{}
}


func (s *jsonFormatter) Format(rawJSON string) (string, bool, error) {
	var js interface{}
	if err := json.Unmarshal([]byte(rawJSON), &js); err != nil {
		return "", false, fmt.Errorf("JSON parse error: %w", err)
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(rawJSON), "", "  "); err != nil {
		return "", false, fmt.Errorf("JSON formatting error: %w", err)
	}

	return prettyJSON.String(), true, nil
}
