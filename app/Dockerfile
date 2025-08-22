# -------------------------
# Build stage
# -------------------------
FROM golang:1.21-alpine AS builder

# çalışma dizini
WORKDIR /app

# Go mod ve bağımlılıkları kopyala
COPY go.mod go.sum ./
RUN go mod download

# kaynak kodu kopyala
COPY *.go ./

# statik olarak binary oluştur
RUN CGO_ENABLED=0 GOOS=linux go build -o pod-watcher main.go

# -------------------------
# Final stage
# -------------------------
FROM alpine:3.18

# gerekli CA sertifikaları
RUN apk add --no-cache ca-certificates

WORKDIR /app

# binary'yi kopyala
COPY --from=builder /app/pod-watcher .

# uygulamayı çalıştır
ENTRYPOINT ["./pod-watcher"]
