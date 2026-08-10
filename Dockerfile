# ==========================================
# BUILD STAGE
# ==========================================
FROM golang:1.25.4-alpine AS build

WORKDIR /app

# Install git (dibutuhkan untuk go mod)
RUN apk add --no-cache git

# Copy go mod & sum dulu (cache dependency)
COPY go.mod go.sum ./
RUN go mod download

# Copy semua source code
COPY . .

# Build binary 
# (CGO_ENABLED=0 ditambahkan agar binary bisa berjalan mandiri tanpa library C bawaan OS di runtime stage)
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# ==========================================
# RUNTIME STAGE
# ==========================================
FROM alpine:latest

WORKDIR /app

# Install CA cert (biar HTTPS jalan) dan tzdata (penting untuk zona waktu, terutama untuk database)
RUN apk add --no-cache ca-certificates tzdata

# Copy binary dari build stage
COPY --from=build /app/main .

# Expose port (sesuai server kamu)
EXPOSE 8000

# Run app
CMD ["./main"]