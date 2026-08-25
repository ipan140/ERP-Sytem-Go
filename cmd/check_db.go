package main

import (
	"fmt"
	"ERP-System/config"
	"ERP-System/app/modules/auth"
	"ERP-System/app/modules/hr/employees"
	"ERP-System/app/modules/services/project"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	var userCount, empCount, projCount int64
	config.DB.Model(&auth.User{}).Count(&userCount)
	config.DB.Model(&employees.Employee{}).Count(&empCount)
	config.DB.Model(&project.Project{}).Count(&projCount)

	fmt.Println("\n=== BUKTI DATA MASUK ===")
	fmt.Printf("Total Data Users    : %d baris\n", userCount)
	fmt.Printf("Total Data Karyawan : %d baris\n", empCount)
	fmt.Printf("Total Data Proyek   : %d baris\n", projCount)
	fmt.Println("========================")
}
