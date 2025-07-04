package tool

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// JSONFormatter defines the interface for the JSON formatting service.
type JSONFormatter interface {
	Format(rawJSON string) (string, bool, error)
}

// jsonFormatter is the implementation of the JSONFormatter interface.
type jsonFormatter struct{}

// NewJSONFormatter creates a new JSONFormatter service.
func NewJSONFormatter() JSONFormatter {
	return &jsonFormatter{}
}

// Format validates and pretty-prints a raw JSON string.
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
