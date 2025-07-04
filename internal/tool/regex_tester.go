package tool

import (
	"fmt"
	"html/template"
	"regexp"
)

// RegexTester defines the interface for the regex testing service.
type RegexTester interface {
	Test(pattern, testString string) (RegexTestResult, error)
}

// RegexTestResult holds the results of a regex test.
type RegexTestResult struct {
	HighlightedText template.HTML
	Matches         [][]string
}

// regexTester is the implementation of the RegexTester interface.
type regexTester struct{}

// NewRegexTester creates a new RegexTester service.
func NewRegexTester() RegexTester {
	return &regexTester{}
}

// Test compiles a regex pattern and finds all matches in a test string.
func (s *regexTester) Test(pattern, testString string) (RegexTestResult, error) {
	if pattern == "" || testString == "" {
		return RegexTestResult{}, fmt.Errorf("please enter both a pattern and a test string")
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return RegexTestResult{}, fmt.Errorf("invalid regex pattern: %w", err)
	}

	matches := re.FindAllStringSubmatch(testString, -1)
	if len(matches) == 0 {
		return RegexTestResult{}, fmt.Errorf("no matches found")
	}

	highlighted := re.ReplaceAllString(testString, "<mark>${0}</mark>")

	return RegexTestResult{
		HighlightedText: template.HTML(highlighted),
		Matches:         matches,
	}, nil
}
