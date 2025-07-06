# DevCortex.ai

[![Go Report Card](https://goreportcard.com/badge/github.com/melihyilman/devcortex.ai)](https://goreportcard.com/report/github.com/melihyilman/devcortex.ai)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

> 🚀 **A Modern, Privacy-First Toolbox for Developers**
>
> DevCortex.ai is an open-source collection of smart, fast, and free utilities designed to streamline daily development tasks. Built with Go, it offers a clean, responsive interface and prioritizes client-side processing to ensure your data remains private.

---

## ✨ Key Features

- **Privacy-Focused**: Most tools run entirely in your browser. Your data never leaves your machine.
- **Comprehensive Toolset**: From formatters and converters to security tools and cheatsheets, all in one place.
- **Fast & Lightweight**: Built with Go for a high-performance backend and a clean, no-nonsense frontend.
- **Open Source**: Transparent, community-driven, and free to use. Contributions are welcome!
- **PWA Ready**: Install DevCortex.ai as a desktop or mobile application for offline access.

---

## 💖 Support

If you find this project useful, please consider supporting its development. Every contribution is highly appreciated!

[![Buy Me A Coffee](https://cdn.buymeacoffee.com/buttons/lato-yellow.png)](https://www.buymeacoffee.com/melihyilman)

---

## 🛠️ Available Tools

The platform offers a wide range of tools, categorized for easy access:

### Featured Tools
- **SQL Performance Analyzer**: Analyze SQL query execution plans and get performance recommendations.
- **Cron Explainer**: Analyze, visualize, and debug cron expressions.
- **Regex Deconstructor**: Visualize, learn, and optimize regex expressions.
- **Peer-to-Peer Scrum Poker**: Real-time, serverless Scrum Poker for agile teams using WebRTC.

### Cryptography & Security
- **Token Generator**: Generate secure random tokens.
- **Bcrypt Hash & Compare**: Hash and compare passwords with Bcrypt.
- **Encrypt/Decrypt Text**: Encrypt and decrypt text with AES.
- **HMAC Generator**: Generate HMAC hashes with various algorithms.
- **RSA Key Pair Generator**: Generate public/private RSA key pairs.
- **Password Strength Analyser**: Analyze the strength of your password.
- **Base64 Encoder/Decoder**: Encode or decode texts to/from Base64 format.
- **JWT Encoder/Decoder**: Encode or decode JSON Web Tokens.
- **Hash Generator**: Generate hashes from text using MD5, SHA1, SHA256, etc.

### Generators
- **BIP39 Passphrase Generator**: Generate BIP39 mnemonic passphrases.
- **UUID/ULID Generator**: Generate universally unique identifiers.
- **Random String Generator**: Generate random strings of a specified length.
- **Favicon Generator**: Create a full set of favicons from a single image.
- **QR Code Generator**: Generate QR codes from text or URLs.
- **Lorem Ipsum Generator**: Generate placeholder text.
- **Crontab Generator**: Interactively generate crontab schedules.

### Converters
- **Image Converter**: Convert images between formats (PNG, JPG, WebP).
- **Data Converters**: Convert between YAML, JSON, XML, CSV, and TOML.
- **JSON to Code**: Convert JSON data to Go, C#, and TypeScript struct/class definitions.
- **Color Format Converter**: Convert colors between Hex, RGB, and HSL.
- **Date/Time Converter**: Convert dates and times between timezones and formats.
- **URL Encoder/Decoder**: Encode or decode URL components.
- **Text Case Converter**: Convert text to various cases (camelCase, snake_case, etc.).

### Formatters & Validators
- **Code Formatters**: Format and beautify HTML, CSS, and JavaScript.
- **JSON Formatter**: Format and validate JSON data.
- **Minifier**: Minify CSS, JS, JSON, HTML, and SVG files to reduce size.

### Web & Network
- **cURL Builder**: Build cURL commands with an easy-to-use interface.
- **SVG Optimizer**: Optimize SVG files for the web.
- **Markdown Previewer**: Instantly preview your Markdown files.
- **Diff Checker**: Compare two texts to see the differences.
- **Text Statistics**: Analyze text for word count, character count, etc.

### Cheatsheets
- A quick reference for **Git, Linux, Regex, Docker, Kubernetes, FFmpeg, SSH, `find`, Nginx, HTTP Status Codes, MIME Types, Vim, cURL, SQL, and Network Commands**.

---

## 🌐 Tech Stack

- **Backend**: Go (Golang)
- **Frontend**: HTML5, Bootstrap 5, Bootstrap Icons
- **Templating**: Go Templates

---

## 📦 Getting Started

To run the project on your local machine, follow these steps:

```sh
# 1. Clone the repository
git clone https://github.com/melihyilman/devcortex.ai.git
cd devcortex.ai

# 2. Tidy up Go modules
go mod tidy

# 3. Run the application
go run ./cmd/web
```

it'll be working on `http://localhost:8080`

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.