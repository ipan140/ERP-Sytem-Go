.PHONY: all run build swagger tidy clean

APP_NAME=erp-system

all: tidy swagger build run

tidy:
	@echo "Merapikan Go Modules..."
	go mod tidy

swagger:
	@echo "Membuat ulang dokumen Swagger API..."
	swag init

build:
	@echo "Membangun aplikasi (Compile)..."
	go build -o $(APP_NAME).exe

run:
	@echo "Menjalankan ERP System..."
	go run main.go

clean:
	@echo "Membersihkan file sisa..."
	del $(APP_NAME).exe

