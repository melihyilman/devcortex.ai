# DevCortex.ai

> 🚀 **Geliştiriciler için Modern Araç Kutusu**
>
> DevCortex.ai, yazılım geliştiricilerinin günlük hayatta ihtiyaç duyduğu birçok aracı tek bir platformda bir araya getiren açık kaynaklı bir projedir. Go dili ile geliştirilmiş olup, modern ve kullanıcı dostu bir arayüze sahiptir.

---

## 🛠️ Araçlar ve Modüller

Proje, aşağıdakiler de dahil olmak üzere çeşitli araçlar sunmaktadır:

### Metin ve Veri Araçları
- **Base64 Encoder/Decoder**: Metin ve verileri Base64 formatına dönüştürün ve geri çözün.
- **URL Encoder/Decoder**: URL'lerde güvenli bir şekilde kullanılabilmesi için metinleri kodlayın ve kodlanmış metinleri geri çözün.
- **Case Converter**: Metinleri `UPPERCASE`, `lowercase`, `Title Case`, `camelCase`, `PascalCase`, `snake_case` gibi farklı formatlara dönüştürün.
- **Lorem Ipsum Generator**: Belirli sayıda paragraf veya kelime içeren anlamsız metinler oluşturun.
- **Random String Generator**: Belirtilen uzunlukta ve karakter setinde rastgele metin dizileri oluşturun.
- **Diff Checker**: İki metin arasındaki farkları satır satır veya kelime kelime karşılaştırın.
- **Markdown Previewer**: Markdown formatındaki metinlerin canlı önizlemesini görün.

### Kodlama ve Biçimlendirme
- **JSON Formatter/Validator**: JSON verilerini biçimlendirin, doğrulayın ve ağaç yapısında görüntüleyin.
- **JSON to Code**: JSON verilerini Go, Python, Java gibi farklı programlama dillerinde kullanılabilecek struct/class yapılarına dönüştürün.
- **Minifier (JS/CSS/HTML)**: JavaScript, CSS ve HTML kodlarını sıkıştırarak dosya boyutunu küçültün.
- **Regex Tester**: Düzenli ifadeleri (Regular Expressions) test edin ve metinler üzerinde eşleşmeleri kontrol edin.
- **Regex Deconstructor**: Karmaşık düzenli ifadeleri analiz ederek bileşenlerine ayırın ve anlamlandırın.
- **SQL Autopsy**: SQL sorgularını analiz ederek performans ve yapısal iyileştirmeler için öneriler alın.

### Güvenlik ve Kriptografi
- **Hash Generator**: MD5, SHA-1, SHA-256, SHA-512 gibi popüler algoritmalarla metinlerin özet (hash) değerlerini hesaplayın.
- **JWT Decoder**: JSON Web Token (JWT) içeriğini (header, payload) analiz edin ve doğrulayın.
- **UUID Generator**: Evrensel olarak benzersiz kimlikler (UUID) oluşturun.

### Zaman ve Renk Araçları
- **Timestamp Converter**: Unix zaman damgalarını insan tarafından okunabilir tarihlere dönüştürün ve tersini yapın.
- **Color Picker & Converter**: Renk seçin ve HEX, RGB, HSL gibi farklı formatlar arasında dönüşüm yapın.
- **Cron Job Explorer**: Cron zamanlamalarını analiz edin ve bir sonraki çalışma zamanlarını görün.

---

## 🌐 Kullanılan Teknolojiler

- **Backend**: Go (Golang)
- **Frontend**: HTML5, Bootstrap 5, Bootstrap Icons
- **Templating**: Go Templates

---

## 📦 Başlangıç

Projeyi yerel makinenizde çalıştırmak için aşağıdaki adımları izleyin:

```sh
# 1. Projeyi klonlayın
git clone https://github.com/your-username/devcortex.ai.git
cd devcortex.ai

# 2. Gerekli Go modüllerini indirin
go mod tidy

# 3. Uygulamayı başlatın
go run ./cmd/web/main.go
```

Uygulama varsayılan olarak `http://localhost:8080` adresinde çalışacaktır.

---

## 👨‍💻 Katkı Sağlama

Projeye katkıda bulunmak isterseniz, lütfen aşağıdaki adımları izleyin:

1.  Projeyi fork'layın ve klonlayın.
2.  Yeni bir özellik veya düzeltme için yeni bir branch oluşturun (`git checkout -b feature/yeni-arac`).
3.  Değişikliklerinizi yapın ve commit'leyin.
4.  Değişikliklerinizi kendi fork'unuza push'layın (`git push origin feature/yeni-arac`).
5.  Bir Pull Request (PR) oluşturun.

Her türlü öneri ve katkı için PR ve issue açmaktan çekinmeyin!