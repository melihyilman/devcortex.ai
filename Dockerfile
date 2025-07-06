# === AŞAMA 1: Derleyici (Builder) ===
# Go'nun kurulu olduğu bir imajı temel alarak başlıyoruz.
# go.mod dosyanızla uyumlu bir sürüm seçin (ör: golang:1.22-alpine).
FROM golang:1.22-alpine AS builder

# Çalışma dizinini /app olarak ayarlıyoruz.
WORKDIR /app

# Önce sadece bağımlılık dosyalarını kopyalayıp indiriyoruz.
# Bu sayede kod değişmediği sürece Docker bu adımı atlayarak derlemeyi hızlandırır.
COPY go.mod go.sum ./
RUN go mod download

# Projenin geri kalan tüm dosyalarını kopyalıyoruz.
COPY . .

# Uygulamayı cmd/web dizininden derliyoruz.
# CGO_ENABLED=0 statik bir binary oluşturmak için önemlidir.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/server ./cmd/web

# === AŞAMA 2: Nihai İmaj (Final Image) ===
# Çok küçük ve güvenli bir temel imaj olan Alpine'ı kullanıyoruz.
FROM alpine:latest

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/web ./web

EXPOSE 8080
CMD ["./server"]