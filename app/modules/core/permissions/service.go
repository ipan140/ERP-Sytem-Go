package permissions

import (
	"ERP-System/config"
	"errors"
)

// Pastikan tabel role_permissions sudah ada (AutoMigrate)
func init() {
	// Menghindari panic saat DB belum terkoneksi penuh, biasanya AutoMigrate ditaruh di main
	// Tapi untuk simulasi cepat, kita taruh di fungsi helper atau abaikan sementara
}

func GetAllPermissionsService() ([]RolePermission, error) {
	var perms []RolePermission
	// Membaca seluruh status toggle dari database
	if config.DB == nil {
		return perms, errors.New("database belum siap")
	}
	err := config.DB.AutoMigrate(&RolePermission{}) // Pastikan tabel ada
	if err != nil {
		return nil, err
	}
	err = config.DB.Find(&perms).Error
	return perms, err
}

func TogglePermissionService(req TogglePermissionRequest) error {
	if config.DB == nil {
		return errors.New("database belum siap")
	}
	err := config.DB.AutoMigrate(&RolePermission{})
	if err != nil {
		return err
	}

	var perm RolePermission
	// Cari apakah aturan untuk Role dan Module ini sudah pernah dibuat sebelumnya
	err = config.DB.Where("role_name = ? AND module = ?", req.RoleName, req.Module).First(&perm).Error
	if err != nil {
		// Jika belum ada, buat baru (Default false semua)
		perm = RolePermission{
			RoleName:  req.RoleName,
			Module:    req.Module,
			CanRead:   false,
			CanWrite:  false,
			CanDelete: false,
		}
	}

	// Update field yang sesuai dengan toggle yang diklik
	switch req.Action {
	case "read":
		perm.CanRead = req.Value
	case "write":
		perm.CanWrite = req.Value
	case "delete":
		perm.CanDelete = req.Value
	case "export":
		perm.CanExport = req.Value
	case "approve":
		perm.CanApprove = req.Value
	default:
		return errors.New("action tidak dikenali (gunakan: read, write, delete, export, approve)")
	}

	// Simpan perubahan ke database
	return config.DB.Save(&perm).Error
}
