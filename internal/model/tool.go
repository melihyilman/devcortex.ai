package model

type Tool struct {
	Name        string
	Description string
	URL         string
	Icon        string // Bootstrap icon class'ı (örn: "bi bi-braces me-2")
}

var FeaturedTools = []Tool{
	{
		Name:        "Zaman Gezgini (Cron)",
		Description: "Cron ifadelerini analiz edin, görselleştirin ve hatalarını ayıklayın.",
		URL:         "/tools/cron-explorer",
		Icon:        "bi bi-clock-history me-2",
	},
	{
		Name:        "Regex Dekonstrüktörü",
		Description: "Regex ifadelerini görselleştirin, öğrenin ve optimize edin.",
		URL:         "/tools/regex-deconstructor",
		Icon:        "bi bi-diagram-3 me-2",
	},
	{
		Name:        "SQL Sorgu Otopsisi",
		Description: "SQL sorgularınızın performansını analiz edin ve iyileştirin.",
		URL:         "/tools/sql-autopsy",
		Icon:        "bi bi-activity me-2",
	},
}

var OtherTools = []Tool{
	{
		Name:        "Base64 Encoder/Decoder",
		Description: "Metinleri Base64 formatına çevirin veya Base64 formatından çözün.",
		URL:         "/tools/base64",
		Icon:        "bi bi-shield-lock me-2",
	},
	{
		Name:        "JWT Decoder",
		Description: "JSON Web Token'ların içeriğini (Header ve Payload) güvenli bir şekilde decode edin.",
		URL:         "/tools/jwt-decoder",
		Icon:        "bi bi-key me-2",
	},
	{
		Name:        "Hash Generator",
		Description: "Metinlerinizin MD5, SHA-1, ve SHA-256 hash değerlerini anında hesaplayın.",
		URL:         "/tools/hash-generator",
		Icon:        "bi bi-fingerprint me-2",
	},
	{
		Name:        "Unix Timestamp Converter",
		Description: "Unix zaman damgasını okunabilir tarihe veya tam tersine çevirin.",
		URL:         "/tools/timestamp-converter",
		Icon:        "bi bi-clock-history me-2",
	},
	{
		Name:        "UUID Generator",
		Description: "Rastgele ve benzersiz UUID’ler oluşturun.",
		URL:         "/tools/uuid-generator",
		Icon:        "bi bi-123 me-2",
	},
	{
		Name:        "Regex Tester",
		Description: "Düzenli ifadeleri test edin ve eşleşmeleri anında görün.",
		URL:         "/tools/regex-tester",
		Icon:        "bi bi-slash-square me-2",
	},
	{
		Name:        "Color Picker / Converter",
		Description: "Renk kodlarını (HEX, RGB, HSL) dönüştürün ve seçin.",
		URL:         "/tools/color-picker",
		Icon:        "bi bi-palette me-2",
	},
	{
		Name:        "Lorem Ipsum Generator",
		Description: "Hızlıca sahte metin (lorem ipsum) oluşturun.",
		URL:         "/tools/lorem-ipsum",
		Icon:        "bi bi-file-earmark-text me-2",
	},
	{
		Name:        "URL Encoder/Decoder",
		Description: "Metinleri URL formatına kodlayın veya çözün.",
		URL:         "/tools/url-encoder",
		Icon:        "bi bi-link-45deg me-2",
	},
	{
		Name:        "Case Converter",
		Description: "Metinlerinizi snake_case, camelCase, PascalCase gibi formatlara dönüştürün.",
		URL:         "/tools/case-converter",
		Icon:        "bi bi-textarea me-2",
	},
	{
		Name:        "HTML/CSS/JS Minifier",
		Description: "Kodunuzu küçültün ve optimize edin.",
		URL:         "/tools/minifier",
		Icon:        "bi bi-file-code me-2",
	},
	{
		Name:        "Markdown Previewer",
		Description: "Markdown dosyalarınızı anında önizleyin.",
		URL:         "/tools/markdown-previewer",
		Icon:        "bi bi-markdown me-2",
	},
	{
		Name:        "Diff Checker",
		Description: "İki metin veya kod bloğu arasındaki farkları karşılaştırın.",
		URL:         "/tools/diff-checker",
		Icon:        "bi bi-file-diff me-2",
	},
	{
		Name:        "Random String Generator",
		Description: "Güvenli rastgele stringler oluşturun.",
		URL:         "/tools/random-string",
		Icon:        "bi bi-shuffle me-2",
	},
}
