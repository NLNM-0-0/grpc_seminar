package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"product/repo"
	"product/server"
	"product/service"
	"strconv"
	"time"
)

type appConfig struct {
	Env  string
	Port int

	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBDatabase string
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalln("Error when loading config:", err)
	}

	fmt.Println("Connecting to database...")
	db, err := connectDatabaseWithRetryIn60s(cfg)
	if err != nil {
		log.Fatalln("Error when connecting to database:", err)
	}
	fmt.Println("Connected to database")

	if cfg.Env == "dev" {
		db = db.Debug()
	}

	productRepo := repo.NewProductRepo(db)
	productService := service.NewService(productRepo)

	log.Fatal(server.ListenGRPC(productService, cfg.Port))
}

func loadConfig() (*appConfig, error) {
	portString := os.Getenv("GO_PORT")
	portNumber, err := strconv.Atoi(portString)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return &appConfig{
		Env:        os.Getenv("GO_ENV"),
		Port:       portNumber,
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBDatabase: os.Getenv("DB_DATABASE"),
	}, nil
}

func connectDatabaseWithRetryIn60s(cfg *appConfig) (*gorm.DB, error) {
	const timeRetry = 60 * time.Second
	var connectDatabase = func(cfg *appConfig) (*gorm.DB, error) {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBDatabase)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		return db.Debug(), nil
	}

	var db *gorm.DB
	var err error

	deadline := time.Now().Add(timeRetry)

	for time.Now().Before(deadline) {
		log.Println("Connecting to database...")
		db, err = connectDatabase(cfg)
		if err == nil {
			return db, nil
		}
		time.Sleep(time.Second)
	}

	return nil, fmt.Errorf("failed to connect to database after retrying for 10 seconds: %w", err)
}
