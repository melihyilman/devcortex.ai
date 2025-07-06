package tool

import (
	"fmt"
	"html/template"
	"regexp"
)


type RegexTester interface {
	Test(pattern, testString string) (RegexTestResult, error)
}


type RegexTestResult struct {
	HighlightedText template.HTML
	Matches         [][]string
}


type regexTester struct{}


func NewRegexTester() RegexTester {
	return &regexTester{}
}


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
