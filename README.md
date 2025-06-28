# DevCortex.ai

> 🚀 **Geliştiriciler için Modern Araç Kutusu**

---

## 📁 Proje Yapısı

```
go.mod
cmd/
  web/
    main.go
internal/
  handler/
    base64_handler.go
    hash_handler.go
    home_Handler.go
    json_handler.go
    jwt_handler.go
    routes.go
    timestamp_handler.go
    tools_handler.go
  model/
    tool.go
  view/
    view.go
templates/
  tools.html
web/
  static/
    css/
      style.css
  template/
    base64.html
    hash.html
    home.html
    json.html
    jwt.html
    layout.html
    timestamp.html
    tools.html
    ...
```

## 🛠️ Araçlar

- **Base64 Encoder/Decoder**
- **Hash Generator**
- **JSON Formatter/Validator**
- **JWT Decoder**
- **Timestamp Converter**
- **UUID Generator**
- **Regex Tester**
- **Color Picker & Converter**
- **URL Encoder/Decoder**
- **Case Converter**
- **Random String Generator**
- **Lorem Ipsum Generator**
- **Minifier (JS/CSS/HTML)**
- **Markdown Previewer**
- **Diff Checker**

## 🌐 Kullanılan Teknolojiler

- **Go (Golang)**
- **HTML5, Bootstrap 5, Bootstrap Icons**
- **Go Templates**

## 📦 Başlangıç

```sh
# Bağımlılıkları yükle
go mod tidy

# Uygulamayı başlat
cd cmd/web
go run main.go
```

Uygulama varsayılan olarak `http://localhost:8080` adresinde çalışır.

## 👨‍💻 Katkı Sağlama

1. Fork'la & clone'la
2. Yeni bir branch oluştur
3. Değişikliklerini yap
4. PR gönder

---

> 💡 Her türlü öneri ve katkı için PR ve issue açabilirsiniz!
