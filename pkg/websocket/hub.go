package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	// Upgrader untuk menaikkan koneksi HTTP menjadi WebSocket
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Di production, sesuaikan dengan Strict CORS
		},
	}
	
	// Clients menyimpan semua koneksi user yang sedang aktif
	// Key: UserID, Value: Koneksi Websocket
	Clients = make(map[string]*websocket.Conn)
	mutex   = &sync.Mutex{}
)

// ServeWS adalah endpoint yang diakses frontend saat pertama kali buka web (misal: /api/ws?user_id=123)
func ServeWS(c echo.Context) error {
	userID := c.QueryParam("user_id")
	if userID == "" {
		return c.String(http.StatusBadRequest, "user_id is required")
	}

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("Gagal upgrade ke WebSocket:", err)
		return err
	}

	// Daftarkan koneksi
	mutex.Lock()
	Clients[userID] = ws
	mutex.Unlock()

	log.Printf("? [WebSocket] User %s terhubung!", userID)

	// Dengarkan ping dari client agar koneksi tidak terputus
	defer func() {
		mutex.Lock()
		delete(Clients, userID)
		mutex.Unlock()
		ws.Close()
		log.Printf("? [WebSocket] User %s terputus.", userID)
	}()

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
	return nil
}

// SendNotification mendorong pop-up ke satu user spesifik secara Real-Time
func SendNotification(userID string, message string) {
	mutex.Lock()
	defer mutex.Unlock()

	ws, ok := Clients[userID]
	if ok {
		err := ws.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Printf("? Gagal mengirim notif WS ke user %s: %v", userID, err)
			ws.Close()
			delete(Clients, userID)
		} else {
			log.Printf("? [WebSocket] Berhasil mengirim pop-up ke User %s!", userID)
		}
	}
}
