package main

import (
	"unico/api"

	"unico/internal/unico/database"
	"unico/internal/unico/migrations"

	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	for _, envVar := range []string{
		"DB_CONNECTION",
	} {
		if _, ok := os.LookupEnv(envVar); !ok {
			log.Fatalf("Required enviroment variable not set: %s", envVar)
		}
	}

	if stage, ok := os.LookupEnv("APP_ENV"); ok {
		if stage == "prod" {
			os.Setenv("GIN_MODE", "release")
		}
	}
}

func main() {

	// Start DB Connection
	connectiondb, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  os.Getenv("DB_CONNECTION"),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Panic("Failed to connect to database: ", err)
	}

	// SET POOL CONNECTIONS
	db, _ := connectiondb.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxLifetime(time.Minute * 5)

	// ping to database
	err = db.Ping()
	// error ping to domain
	if err != nil {
		log.Panic(err)
	}

	err = migrations.Run(connectiondb)
	if err != nil {
		log.Panic("Failed to run migration: ", err)
	}
	server := gin.Default()
	databasedomain := database.NewDatabaseHandle(connectiondb)
	api.ConfigureRoutes(server, databasedomain)
	server.Run(":3000")

}
