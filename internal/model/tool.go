package model

type Tool struct {
	Name        string
	Description string
	URL         string
	Icon        string
}

var FeaturedTools = []Tool{
	{
		Name:        "Cron Explainer",
		Description: "Analyze, visualize, and debug cron expressions.",
		URL:         "/tools/cron-explainer",
		Icon:        "bi bi-calendar-event",
	},
	{
		Name:        "Regex Deconstructor",
		Description: "Visualize, learn, and optimize regex expressions.",
		URL:         "/tools/regex-deconstructor",
		Icon:        "bi bi-braces-asterisk",
	},
	
}

var OtherTools = []Tool{
	{
		Name:        "JSON Formatter",
		Description: "Format and validate JSON data.",
		URL:         "/tools/json-formatter",
		Icon:        "bi bi-braces",
	},
	{
		Name:        "JSON to Code",
		Description: "Convert JSON data to your selected programming language.",
		URL:         "/tools/json-to-code",
		Icon:        "bi bi-file-code",
	},
	{
		Name:        "Base64 Encoder/Decoder",
		Description: "Encode or decode texts to/from Base64 format.",
		URL:         "/tools/base64",
		Icon:        "bi bi-arrow-repeat",
	},
	{
		Name:        "JWT Decoder",
		Description: "Securely decode the content (Header and Payload) of JSON Web Tokens.",
		URL:         "/tools/jwt-decoder",
		Icon:        "bi bi-key-fill",
	},
	{
		Name:        "Hash Generator",
		Description: "Instantly calculate MD5, SHA-1, and SHA-256 hash values of your texts.",
		URL:         "/tools/hash-generator",
		Icon:        "bi bi-fingerprint",
	},
	{
		Name:        "Unix Timestamp Converter",
		Description: "Convert Unix timestamp to human-readable date and vice versa.",
		URL:         "/tools/timestamp-converter",
		Icon:        "bi bi-hourglass-split",
	},
	{
		Name:        "UUID Generator",
		Description: "Generate random and unique UUIDs.",
		URL:         "/tools/uuid-generator",
		Icon:        "bi bi-grip-horizontal",
	},
	{
		Name:        "Regex Tester",
		Description: "Test regular expressions and see matches instantly.",
		URL:         "/tools/regex-tester",
		Icon:        "bi bi-search",
	},
	{
		Name:        "Color Picker / Converter",
		Description: "Convert and select color codes (HEX, RGB, HSL).",
		URL:         "/tools/color-picker",
		Icon:        "bi bi-palette-fill",
	},
	{
		Name:        "Lorem Ipsum Generator",
		Description: "Quickly generate fake text (lorem ipsum).",
		URL:         "/tools/lorem-ipsum",
		Icon:        "bi bi-file-earmark-text-fill",
	},
	{
		Name:        "URL Encoder/Decoder",
		Description: "Encode or decode texts to/from URL format.",
		URL:         "/tools/url-encoder",
		Icon:        "bi bi-link-45deg",
	},
	{
		Name:        "Case Converter",
		Description: "Convert your texts to formats like snake_case, camelCase, PascalCase.",
		URL:         "/tools/case-converter",
		Icon:        "bi bi-letter-case",
	},
	{
		Name:        "HTML/CSS/JS Minifier",
		Description: "Minify and optimize your code.",
		URL:         "/tools/minifier",
		Icon:        "bi bi-file-zip-fill",
	},
	{
		Name:        "Markdown Previewer",
		Description: "Instantly preview your Markdown files.",
		URL:         "/tools/markdown-previewer",
		Icon:        "bi bi-markdown-fill",
	},
	{
		Name:        "Diff Checker",
		Description: "Compare differences between two text or code blocks.",
		URL:         "/tools/diff-checker",
		Icon:        "bi bi-file-diff-fill",
	},
	{
		Name:        "Random String Generator",
		Description: "Generate secure random strings.",
		URL:         "/tools/random-string",
		Icon:        "bi bi-shuffle",
	},
}