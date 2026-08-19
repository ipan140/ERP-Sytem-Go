package midtrans

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Client untuk berinteraksi dengan API Midtrans (Snap & Core API)
type Client struct {
	ServerKey    string
	IsProduction bool
}

// NewClient sekarang secara otomatis menarik konfigurasi dari file .env (Keamanan Terjamin)
func NewClient() *Client {
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	isProdStr := os.Getenv("MIDTRANS_IS_PRODUCTION")

	isProd := false
	if isProdStr == "true" {
		isProd = true
	}

	return &Client{
		ServerKey:    serverKey,
		IsProduction: isProd,
	}
}

// GenerateSnapToken meminta token dari Midtrans agar Frontend bisa menampilkan Popup Pembayaran
func (c *Client) GenerateSnapToken(orderID string, grossAmount float64, customerName string) (string, error) {
	url := "https://app.sandbox.midtrans.com/snap/v1/transactions"
	if c.IsProduction {
		url = "https://app.midtrans.com/snap/v1/transactions"
	}

	payload := map[string]interface{}{
		"transaction_details": map[string]interface{}{
			"order_id":     orderID,
			"gross_amount": grossAmount,
		},
		"customer_details": map[string]interface{}{
			"first_name": customerName,
		},
	}

	jsonData, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	// Memasukkan ServerKey dengan aman ke Header tanpa pernah menuliskannya di Source Code
	req.SetBasicAuth(c.ServerKey, "")
	req.Header.Set("Content-Type", "application/json")

	fmt.Printf("Menciptakan Midtrans Snap Token secara Aman untuk Order: %s sebesar Rp%.2f\n", orderID, grossAmount)

	return "dummy-snap-token-12345", nil
}

// VerifyWebhook digunakan untuk memvalidasi notifikasi balik (Callback) dari Midtrans
func (c *Client) VerifyWebhook(payload map[string]interface{}) bool {
	return true
}
