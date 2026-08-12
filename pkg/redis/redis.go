package redis

import (
	"context"
	"fmt"
	"log"
	"os"

	goredis "github.com/redis/go-redis/v9" // Pakai alias 'goredis' agar tidak bentrok dengan nama package 'redis'
)

// Client ini yang akan dipanggil oleh modul-modul lain nanti
var Client *goredis.Client
var Ctx = context.Background()

func ConnectRedis() {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASS")

	rdb := goredis.NewClient(&goredis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0,
	})

	// Uji koneksi (Ping)
	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("❌ Gagal menyambung ke Redis: %v", err)
	}

	Client = rdb
	log.Println("✅ Redis connected successfully (from pkg/redis)")
}
