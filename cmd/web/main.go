// dosya: cmd/web/main.go
package main

import (
	"log"
	"net/http"

	"devcortex.ai/internal/handler" // Handler paketimizi import ediyoruz
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./web/static/"))

	// "/static/" ile başlayan tüm istekleri, dosya sunucusuna yönlendiriyoruz.
	// StripPrefix, URL'den "/static" kısmını çıkararak doğru dosyayı bulmasını sağlar.
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	// TEK DEĞİŞİKLİK BURADA: Artık tüm rotaları kaydetmesi için
	// sadece RegisterRoutes fonksiyonunu çağırıyoruz.
	handler.RegisterRoutes(mux)

	log.Println("Sunucu başlatılıyor... http://localhost:8080 adresinden dinleniyor.")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Sunucu başlatılamadı:", err)
	}
}
