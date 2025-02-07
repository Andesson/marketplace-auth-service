package config

import (
	"fmt"
	"log"

	migrations "github.com/Andesson/marketplace-auth-service/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgree() (*gorm.DB, error) {
	logger := GetLogger("postgree")
	dsn := "host=localhost user=user password=123 dbname=auth_db port=5432 sslmode=disable"
	fmt.Println("🔄 Tentando conectar ao banco de dados...")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("postgree opening error: %v", err)
		log.Fatalf("❌ Falha na conexão com o banco: %v", err)
		return nil, err
	}
	fmt.Println("✅ Conexão bem-sucedida!")
	migrations.RunMigrations(db)
	return db, nil
}
