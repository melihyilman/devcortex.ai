package model

type Tool struct {
	Name        string
	Description string
	URL         string
	Icon        string
	Category    string
}

var FeaturedTools = []Tool{
	{
		Name:        "SQL Performance Analyzer",
		Description: "Analyze SQL query execution plans and get performance recommendations.",
		URL:         "/tools/sql-performance",
		Icon:        "bi bi-activity",
		Category:    "featured",
	},
	{
		Name:        "Cron Explainer",
		Description: "Analyze, visualize, and debug cron expressions.",
		URL:         "/tools/cron-explainer",
		Icon:        "bi bi-calendar-event",
		Category:    "featured",
	},
	{
		Name:        "Regex Deconstructor",
		Description: "Visualize, learn, and optimize regex expressions.",
		URL:         "/tools/regex-deconstructor",
		Icon:        "bi bi-braces-asterisk",
		Category:    "featured",
	},
}

var OtherTools = []Tool{
	// Cryptography & Security
	{Name: "Token Generator", Description: "Generate secure tokens.", URL: "/tools/token-generator", Icon: "bi bi-key", Category: "Cryptography & Security"},
	{Name: "Bcrypt Hash & Compare", Description: "Hash and compare passwords with Bcrypt.", URL: "/tools/bcrypt-hash", Icon: "bi bi-shield-lock", Category: "Cryptography & Security"},
	{Name: "Encrypt/Decrypt Text", Description: "Encrypt and decrypt text with AES.", URL: "/tools/encrypt-decrypt", Icon: "bi bi-lock", Category: "Cryptography & Security"},
	{Name: "HMAC Generator", Description: "Generate HMAC hashes.", URL: "/tools/hmac-generator", Icon: "bi bi-hash", Category: "Cryptography & Security"},
	{Name: "RSA Key Pair Generator", Description: "Generate public/private RSA key pairs.", URL: "/tools/rsa-generator", Icon: "bi bi-diagram-2", Category: "Cryptography & Security"},
	{Name: "Password Strength Analyser", Description: "Analyze the strength of your password.", URL: "/tools/password-strength", Icon: "bi bi-bar-chart", Category: "Cryptography & Security"},
	{Name: "Base64 Encoder/Decoder", Description: "Encode or decode texts to/from Base64 format.", URL: "/tools/base64", Icon: "bi bi-arrow-repeat", Category: "Cryptography & Security"},
	{Name: "JWT Encoder/Decoder", Description: "Encode or decode JWT tokens.", URL: "/tools/jwt-decoder", Icon: "bi bi-box-arrow-in-right", Category: "Cryptography & Security"},
	{Name: "Hash Generator", Description: "Generate hashes from text.", URL: "/tools/hash-generator", Icon: "bi bi-fingerprint", Category: "Cryptography & Security"},

	// Generators
	{Name: "BIP39 Passphrase Generator", Description: "Generate BIP39 mnemonic passphrases.", URL: "/tools/bip39-generator", Icon: "bi bi-wallet2", Category: "Generators"},
	{Name: "UUID Generator", Description: "Generate UUIDs.", URL: "/tools/uuid-generator", Icon: "bi bi-node-plus", Category: "Generators"},
	{Name: "ULID Generator", Description: "Generate ULIDs.", URL: "/tools/ulid-generator", Icon: "bi bi-123", Category: "Generators"},
	{Name: "Random String Generator", Description: "Generate random strings of a specified length.", URL: "/tools/random-string", Icon: "bi bi-shuffle", Category: "Generators"},
	{Name: "Favicon Generator", Description: "Create favicons from images.", URL: "/tools/favicon-generator", Icon: "bi bi-gem", Category: "Generators"},
	{Name: "QR Code Generator", Description: "Generate QR codes from text or URLs.", URL: "/tools/qr-code-generator", Icon: "bi bi-qr-code", Category: "Generators"},
	{Name: "Lorem Ipsum Generator", Description: "Generate Lorem Ipsum placeholder text.", URL: "/tools/lorem-ipsum", Icon: "bi bi-file-text", Category: "Generators"},
	{Name: "Crontab Generator", Description: "Generate crontab schedules.", URL: "/tools/crontab-generator", Icon: "bi bi-calendar-plus", Category: "Generators"},

	// Converters
	{Name: "Image Converter", Description: "Convert images between formats (PNG, JPG, WebP).", URL: "/tools/image-converter", Icon: "bi bi-image-alt", Category: "Converters"},
	{Name: "YAML to JSON Converter", Description: "Convert YAML data to JSON.", URL: "/tools/yaml-to-json", Icon: "bi bi-file-earmark-code", Category: "Converters"},
	{Name: "JSON to XML Converter", Description: "Convert JSON data to XML.", URL: "/tools/json-to-xml", Icon: "bi bi-file-earmark-code", Category: "Converters"},
	{Name: "JSON to CSV Converter", Description: "Convert JSON data to CSV.", URL: "/tools/json-to-csv", Icon: "bi bi-file-earmark-spreadsheet", Category: "Converters"},
	{Name: "JSON to TOML Converter", Description: "Convert JSON data to TOML.", URL: "/tools/json-to-toml", Icon: "bi bi-file-earmark-code", Category: "Converters"},
	{Name: "YAML to TOML Converter", Description: "Convert YAML data to TOML.", URL: "/tools/yaml-to-toml", Icon: "bi bi-file-earmark-code", Category: "Converters"},
	{Name: "JSON to Go Struct Converter", Description: "Convert JSON data to Go struct definitions.", URL: "/tools/json-to-code", Icon: "bi bi-file-code", Category: "Converters"},
	{Name: "Color Format Converter", Description: "Convert colors between formats (Hex, RGB, HSL).", URL: "/tools/color-converter", Icon: "bi bi-palette", Category: "Converters"},
	{Name: "Date/Time Converter", Description: "Convert dates and times between timezones and formats.", URL: "/tools/datetime-converter", Icon: "bi bi-clock-history", Category: "Converters"},
	{Name: "URL Encoder/Decoder", Description: "Encode or decode URL components.", URL: "/tools/url-encoder", Icon: "bi bi-link-45deg", Category: "Converters"},
	{Name: "Text Case Converter", Description: "Convert text to different cases (uppercase, lowercase, etc.).", URL: "/tools/case-converter", Icon: "bi bi-type", Category: "Converters"},

	// Formatters & Validators
	{Name: "HTML Formatter", Description: "Format and beautify HTML code.", URL: "/tools/html-formatter", Icon: "bi bi-code-slash", Category: "Formatters & Validators"},
	{Name: "JavaScript Formatter", Description: "Format and beautify JavaScript code.", URL: "/tools/js-formatter", Icon: "bi bi-file-earmark-code", Category: "Formatters & Validators"},
	{Name: "CSS Formatter", Description: "Format and beautify CSS code.", URL: "/tools/css-formatter", Icon: "bi bi-file-earmark-code", Category: "Formatters & Validators"},
	{Name: "JSON Formatter", Description: "Format and validate JSON data.", URL: "/tools/json-formatter", Icon: "bi bi-braces", Category: "Formatters & Validators"},
	{Name: "Minifier", Description: "Minify CSS, JS, JSON, HTML, and SVG files.", URL: "/tools/minifier", Icon: "bi bi-file-zip", Category: "Formatters & Validators"},

	// Web & Network
	{Name: "cURL Builder", Description: "Build cURL commands with an easy interface.", URL: "/tools/curl-builder", Icon: "bi bi-terminal", Category: "Web & Network"},
	{Name: "SVG Optimizer", Description: "Optimize SVG files for the web.", URL: "/tools/svg-optimizer", Icon: "bi bi-vector-pen", Category: "Web & Network"},
	{Name: "Markdown Previewer", Description: "Instantly preview your Markdown files.", URL: "/tools/markdown-previewer", Icon: "bi bi-markdown", Category: "Web & Network"},
	{Name: "Diff Checker", Description: "Compare two texts to see the differences.", URL: "/tools/diff-checker", Icon: "bi bi-distribute-vertical", Category: "Web & Network"},
	{Name: "Text Statistics", Description: "Analyze text for word count, character count, etc.", URL: "/tools/text-stats", Icon: "bi bi-card-text", Category: "Web & Network"},

	// Cheatsheets
	{Name: "Git Cheatsheet", Description: "Quick reference for Git commands.", URL: "/tools/git-cheatsheet", Icon: "bi bi-git", Category: "Cheatsheets"},
	{Name: "Linux Cheatsheet", Description: "Quick reference for Linux commands.", URL: "/tools/linux-cheatsheet", Icon: "bi bi-terminal-fill", Category: "Cheatsheets"},
	{Name: "Regex Cheatsheet", Description: "Quick reference for regular expressions.", URL: "/tools/regex-cheatsheet", Icon: "bi bi-braces-asterisk", Category: "Cheatsheets"},
	{Name: "Docker Cheatsheet", Description: "Quick reference for Docker commands.", URL: "/tools/docker-cheatsheet", Icon: "bi bi-box-seam", Category: "Cheatsheets"},
	{Name: "Kubernetes Cheatsheet", Description: "Quick reference for Kubernetes (kubectl) commands.", URL: "/tools/kubernetes-cheatsheet", Icon: "bi bi-bounding-box", Category: "Cheatsheets"},
	{Name: "FFmpeg Cheatsheet", Description: "Quick reference for FFmpeg commands.", URL: "/tools/ffmpeg-cheatsheet", Icon: "bi bi-camera-reels", Category: "Cheatsheets"},
	{Name: "SSH Cheatsheet", Description: "Quick reference for SSH commands.", URL: "/tools/ssh-cheatsheet", Icon: "bi bi-terminal", Category: "Cheatsheets"},
	{Name: "Linux 'find' Cheatsheet", Description: "Quick reference for the 'find' command.", URL: "/tools/find-cheatsheet", Icon: "bi bi-search", Category: "Cheatsheets"},
	{Name: "Nginx Cheatsheet", Description: "Quick reference for Nginx configuration.", URL: "/tools/nginx-cheatsheet", Icon: "bi bi-server", Category: "Cheatsheets"},
	{Name: "HTTP Status Codes", Description: "Quick reference for HTTP status codes.", URL: "/tools/http-status-codes", Icon: "bi bi-cloud-slash", Category: "Cheatsheets"},
	{Name: "File Types & MIME", Description: "Quick reference for file types and MIME types.", URL: "/tools/mime-types-cheatsheet", Icon: "bi bi-file-earmark", Category: "Cheatsheets"},
	{Name: "Vim Cheatsheet", Description: "Quick reference for Vim commands.", URL: "/tools/vim-cheatsheet", Icon: "bi bi-keyboard", Category: "Cheatsheets"},
	{Name: "cURL Cheatsheet", Description: "Quick reference for cURL commands.", URL: "/tools/curl-cheatsheet", Icon: "bi bi-terminal", Category: "Cheatsheets"},
	{Name: "SQL Cheatsheet", Description: "Quick reference for SQL commands.", URL: "/tools/sql-cheatsheet", Icon: "bi bi-database", Category: "Cheatsheets"},
	{Name: "Network Commands Cheatsheet", Description: "Quick reference for network commands.", URL: "/tools/net-cheatsheet", Icon: "bi bi-reception-4", Category: "Cheatsheets"},
}
