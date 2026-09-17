package mailer

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
	"strings"
	"time"
)

func CreateEmailLogService(data *EmailLog) error {
	return CreateEmailLog(data)
}

func GetAllEmailLogService() ([]EmailLog, error) {
	return GetAllEmailLog()
}

func GetPaginatedEmailLogService(offset, limit int, search string) ([]EmailLog, int64, error) {
	return GetPaginatedEmailLog(offset, limit, search)
}

func GetEmailLogByIDService(id uint) (*EmailLog, error) {
	return GetEmailLogByID(id)
}

func UpdateEmailLogService(data *EmailLog) error {
	return UpdateEmailLog(data)
}

func DeleteEmailLogService(id uint) error {
	return DeleteEmailLog(id)
}

func GetSmtpConfigService() (*SmtpConfig, error) {
	return GetActiveSmtpConfig()
}

func SaveSmtpConfigService(data *SmtpConfig) error {
	return SaveSmtpConfig(data)
}

// TestSendEmailService melakukan koneksi riil ke SMTP server dan mengirimkan email tes
func TestSendEmailService(recipient string) (string, error) {
	cfg, err := GetActiveSmtpConfig()
	if err != nil {
		return "", fmt.Errorf("gagal mengambil konfigurasi SMTP: %w", err)
	}

	if strings.TrimSpace(cfg.Host) == "" {
		return "", fmt.Errorf("SMTP Host belum dikonfigurasi")
	}

	if strings.TrimSpace(recipient) == "" {
		recipient = cfg.Username
	}

	subject := "Uji Coba Pengiriman Email SMTP ERP Enterprise"
	body := fmt.Sprintf("Halo,\n\nIni adalah email uji coba dari sistem ERP Enterprise untuk memastikan konfigurasi SMTP Host %s:%d berjalan dengan baik.\n\nWaktu Kirim: %s\nPengirim: %s\n\nSalam,\nERP System Administrator",
		cfg.Host, cfg.Port, time.Now().Format("02 Jan 2006 15:04:05 MST"), cfg.SenderName)

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	// Persiapkan pesan RFC 822
	fromHeader := fmt.Sprintf("%s <%s>", cfg.SenderName, cfg.Username)
	if cfg.SenderName == "" {
		fromHeader = cfg.Username
	}

	msg := []byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n%s",
		fromHeader, recipient, subject, body))

	var sendErr error

	// Dial TCP dengan timeout 7 detik
	timeout := 7 * time.Second
	conn, dialErr := net.DialTimeout("tcp", addr, timeout)
	if dialErr != nil {
		sendErr = fmt.Errorf("gagal terhubung ke host %s:%d (Connection timeout/refused): %v", cfg.Host, cfg.Port, dialErr)
	} else {
		defer conn.Close()

		if strings.EqualFold(cfg.Encryption, "SSL_TLS") || cfg.Port == 465 {
			// TLS connection langsung
			tlsConfig := &tls.Config{
				InsecureSkipVerify: true,
				ServerName:         cfg.Host,
			}
			tlsConn := tls.Client(conn, tlsConfig)
			c, clientErr := smtp.NewClient(tlsConn, cfg.Host)
			if clientErr != nil {
				sendErr = fmt.Errorf("gagal inisialisasi client SMTP TLS: %v", clientErr)
			} else {
				defer c.Quit()
				if cfg.Username != "" && cfg.Password != "" {
					auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host)
					if authErr := c.Auth(auth); authErr != nil {
						sendErr = fmt.Errorf("autentikasi SMTP gagal: %v", authErr)
					}
				}
				if sendErr == nil {
					if err := c.Mail(cfg.Username); err != nil {
						sendErr = fmt.Errorf("MAIL command gagal: %v", err)
					} else if err := c.Rcpt(recipient); err != nil {
						sendErr = fmt.Errorf("RCPT command gagal: %v", err)
					} else {
						w, err := c.Data()
						if err != nil {
							sendErr = fmt.Errorf("DATA command gagal: %v", err)
						} else {
							_, _ = w.Write(msg)
							_ = w.Close()
						}
					}
				}
			}
		} else {
			// Plain / STARTTLS
			c, clientErr := smtp.NewClient(conn, cfg.Host)
			if clientErr != nil {
				sendErr = fmt.Errorf("gagal inisialisasi client SMTP: %v", clientErr)
			} else {
				defer c.Quit()
				if strings.EqualFold(cfg.Encryption, "STARTTLS") || cfg.Port == 587 {
					if ok, _ := c.Extension("STARTTLS"); ok {
						tlsConfig := &tls.Config{
							InsecureSkipVerify: true,
							ServerName:         cfg.Host,
						}
						if startTlsErr := c.StartTLS(tlsConfig); startTlsErr != nil {
							sendErr = fmt.Errorf("gagal STARTTLS: %v", startTlsErr)
						}
					}
				}

				if sendErr == nil && cfg.Username != "" && cfg.Password != "" {
					auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host)
					if authErr := c.Auth(auth); authErr != nil {
						sendErr = fmt.Errorf("autentikasi SMTP gagal: %v", authErr)
					}
				}

				if sendErr == nil {
					if err := c.Mail(cfg.Username); err != nil {
						sendErr = fmt.Errorf("MAIL command gagal: %v", err)
					} else if err := c.Rcpt(recipient); err != nil {
						sendErr = fmt.Errorf("RCPT command gagal: %v", err)
					} else {
						w, err := c.Data()
						if err != nil {
							sendErr = fmt.Errorf("DATA command gagal: %v", err)
						} else {
							_, _ = w.Write(msg)
							_ = w.Close()
						}
					}
				}
			}
		}
	}

	// Catat riwayat ke setting.email_logs
	now := time.Now()
	status := "SENT"
	if sendErr != nil {
		status = "FAILED"
	}
	emailLog := EmailLog{
		Recipient: recipient,
		Subject:   subject,
		Body:      body,
		Status:    status,
		SentAt:    &now,
	}
	_ = CreateEmailLog(&emailLog)

	if sendErr != nil {
		return "", sendErr
	}

	return fmt.Sprintf("Email uji coba berhasil dikirim ke %s via SMTP Host %s:%d", recipient, cfg.Host, cfg.Port), nil
}
