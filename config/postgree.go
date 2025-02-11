package config

import (
	"fmt"
	"log"
	"os"

	migrations "github.com/Andesson/marketplace-auth-service/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgree() (*gorm.DB, error) {
	logger := GetLogger("postgree")
	host := os.Getenv("host")
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")
	port := os.Getenv("port")
	sslmode := os.Getenv("sslmode")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)
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
