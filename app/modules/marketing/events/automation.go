package events

import "fmt"

func GenerateTicketBarcode(ticketID uint) {
	fmt.Printf("[ENGINE] Otomatis menciptakan tiket Barcode/QR Code untuk ID Tiket %d...\n", ticketID)
}
