package constants

// Role adalah tipe data khusus untuk memastikan konsistensi hak akses (RBAC)
type Role string

const (
	// Level Administrator & Eksekutif
	RoleSuperadmin Role = "SUPERADMIN"
	RoleDirector   Role = "DIRECTOR"

	// Divisi Finance
	RoleFinanceManager Role = "FINANCE_MANAGER"
	RoleFinanceBilling Role = "FINANCE_BILLING"

	// Divisi Supply Chain (Gudang & Pembelian)
	RoleWarehouseManager Role = "WAREHOUSE_MANAGER"
	RoleWarehouseWorker  Role = "WAREHOUSE_WORKER"
	RolePurchasing       Role = "PURCHASING"

	// Divisi Produksi & Kualitas (Pabrik)
	RoleManufacturingManager Role = "MANUFACTURING_MANAGER"
	RoleQualityManager       Role = "QUALITY_MANAGER"
	RoleMaintenanceManager   Role = "MAINTENANCE_MANAGER"

	// Divisi Sales (Penjualan)
	RoleSalesManager Role = "SALES_MANAGER"
	RoleSalesStaff   Role = "SALES_STAFF"

	// Divisi HR (Sumber Daya Manusia)
	RoleHRManager    Role = "HR_MANAGER"
	RoleFleetManager Role = "FLEET_MANAGER" // Mengurus aset kendaraan
	RoleEmployee     Role = "EMPLOYEE"      // Role default untuk semua pegawai internal

	// Divisi Marketing & Website
	RoleMarketingManager Role = "MARKETING_MANAGER"

	// Divisi Services (Proyek & Bantuan Pelanggan)
	RoleProjectManager Role = "PROJECT_MANAGER"
	RoleSupportAgent   Role = "SUPPORT_AGENT"

	// Eksternal (Publik / Pelanggan)
	RoleCustomer Role = "CUSTOMER"
	RoleGuest    Role = "GUEST"
)
