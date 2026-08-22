package barcode

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	redisPkg "ERP-System/pkg/redis"
)

var ctx = context.Background()

func CreateBarcodeConfigService(data *BarcodeNomenclature) error {
	err := CreateBarcodeConfig(data)
	if err == nil && redisPkg.Client != nil {
		// Invalidate cache
		redisPkg.Client.Del(ctx, "supply_chain:barcode:all")
		fmt.Println("🧹 [Redis Clear] Invalidated 'supply_chain:barcode:all'")
	}
	return err
}

func GetAllBarcodeConfigService() ([]BarcodeNomenclature, error) {
	cacheKey := "supply_chain:barcode:all"

	if redisPkg.Client != nil {
		redisRepo := NewBarcodeRedis(redisPkg.Client)
		cachedData, err := redisRepo.GetBarcodeMapping(ctx, cacheKey)
		if err == nil && cachedData != "" {
			var configs []BarcodeNomenclature
			_ = json.Unmarshal([]byte(cachedData), &configs)
			fmt.Println("🚀 [Redis Hit] Fetching Barcode Configs from Cache!")
			return configs, nil
		}
	}

	fmt.Println("🐢 [DB Hit] Fetching Barcode Configs from PostgreSQL...")
	configs, err := GetAllBarcodeConfig()

	if err == nil && redisPkg.Client != nil {
		bytes, _ := json.Marshal(configs)
		_ = redisPkg.Client.Set(ctx, cacheKey, bytes, 24*time.Hour).Err()
	}

	return configs, err
}

func GetBarcodeConfigByIDService(id uint) (*BarcodeNomenclature, error) {
	return GetBarcodeConfigByID(id)
}

func UpdateBarcodeConfigService(data *BarcodeNomenclature) error {
	return UpdateBarcodeConfig(data)
}

func DeleteBarcodeConfigService(id uint) error {
	return DeleteBarcodeConfig(id)
}
