package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

var ModelsToMigrate []interface{}

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		AppConfig.DBHost, AppConfig.DBUser, AppConfig.DBPass, AppConfig.DBName, AppConfig.DBPort, AppConfig.DBSSLMode, AppConfig.DBTimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             500 * time.Millisecond,
				LogLevel:                  logger.Warn,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
	})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	schemas := []string{"setting", "hrd", "sales", "marketing", "services", "supply_chain", "finance", "website_portal"}
	for _, schema := range schemas {
		db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s;", schema))
	}

	// Clean orphaned records if table exists so AutoMigrate foreign keys succeed
	db.Exec("DELETE FROM sales.sale_order_lines WHERE order_id NOT IN (SELECT id FROM sales.sale_orders);")

	if len(ModelsToMigrate) > 0 {
		err = db.AutoMigrate(ModelsToMigrate...)
		AddManualForeignKeys(db)
		if err != nil {
			log.Fatal("AutoMigrate error:", err)
		}
	}

	DB = db
	log.Println("PostgreSQL connected successfully")
}

func AddManualForeignKeys(db *gorm.DB) {
	queries := []string{
		"ALTER TABLE finance.payments ADD CONSTRAINT fk_finance_payments_invoice FOREIGN KEY (invoice_id) REFERENCES finance.invoices(id) ON DELETE SET NULL;",
		"ALTER TABLE finance.payment_transactions ADD CONSTRAINT fk_finance_payment_transactions_invoice FOREIGN KEY (invoice_id) REFERENCES finance.invoices(id) ON DELETE SET NULL;",
		"ALTER TABLE finance.payment_transactions ADD CONSTRAINT fk_finance_payment_acquirer FOREIGN KEY (acquirer_id) REFERENCES finance.payment_acquirers(id) ON DELETE SET NULL;",
		"ALTER TABLE finance.tax_repartition_lines ADD CONSTRAINT fk_finance_tax_repartition_account FOREIGN KEY (account_id) REFERENCES finance.accounts(id) ON DELETE SET NULL;",
		"ALTER TABLE services.projects ADD CONSTRAINT fk_services_projects_sale_order FOREIGN KEY (sale_order_id) REFERENCES sales.sale_orders(id) ON DELETE SET NULL;",

		"ALTER TABLE finance.fiscal_positions ADD CONSTRAINT fk_finance_fiscal_positions_tax_src FOREIGN KEY (tax_src_id) REFERENCES finance.taxes(id) ON DELETE SET NULL;",
		"ALTER TABLE finance.fiscal_positions ADD CONSTRAINT fk_finance_fiscal_positions_tax_dest FOREIGN KEY (tax_dest_id) REFERENCES finance.taxes(id) ON DELETE SET NULL;",
		"ALTER TABLE finance.invoices ADD CONSTRAINT fk_finance_invoices_incoterm FOREIGN KEY (incoterm_id) REFERENCES finance.account_incotermses(id) ON DELETE SET NULL;",
	}
	for _, q := range queries {
		db.Exec(q)
	}
}
