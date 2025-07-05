package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func RegexCheatsheetTool(w http.ResponseWriter, r *http.Request) {
	cheatsheetData := map[string][]Command{
		"Character Classes": {
			{Command: ".", Description: "Matches any single character except newline."},
			{Command: "\\d", Description: "Matches any digit (0-9)."},
			{Command: "\\D", Description: "Matches any non-digit."},
			{Command: "\\w", Description: "Matches any word character (alphanumeric & underscore)."},
			{Command: "\\W", Description: "Matches any non-word character."},
			{Command: "\\s", Description: "Matches any whitespace character (space, tab, newline)."},
			{Command: "\\S", Description: "Matches any non-whitespace character."},
			{Command: "[abc]", Description: "Matches any character in the set (a, b, or c)."},
			{Command: "[^abc]", Description: "Matches any character not in the set."},
			{Command: "[a-z]", Description: "Matches any character in the range a to z."},
		},
		"Anchors": {
			{Command: "^", Description: "Matches the beginning of the string."},
			{Command: "$", Description: "Matches the end of the string."},
			{Command: "\\b", Description: "Matches a word boundary."},
			{Command: "\\B", Description: "Matches a non-word boundary."},
		},
		"Quantifiers": {
			{Command: "*", Description: "Matches 0 or more repetitions of the preceding character."},
			{Command: "+", Description: "Matches 1 or more repetitions of the preceding character."},
			{Command: "?", Description: "Matches 0 or 1 repetition of the preceding character."},
			{Command: "{n}", Description: "Matches exactly n repetitions."},
			{Command: "{n,}", Description: "Matches n or more repetitions."},
			{Command: "{n,m}", Description: "Matches between n and m repetitions."},
		},
		"Grouping and Capturing": {
			{Command: "(...)", Description: "Captures the matched group."},
			{Command: "(?:...)", Description: "Non-capturing group."},
			{Command: "|", Description: "Acts like a boolean OR. Matches the expression before or after the |."},
		},
		"Flags / Modifiers": {
			{Command: "g", Description: "Global search (don't return after the first match)."},
			{Command: "i", Description: "Case-insensitive search."},
			{Command: "m", Description: "Multi-line search."},
		},
	}

	data := &view.PageData{
		Title:            "Regex Cheatsheet",
		ToolSpecificData: cheatsheetData,
	}
	view.Render(w, r, "regex-cheatsheet.html", data)
}
