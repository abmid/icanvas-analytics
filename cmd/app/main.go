package main

import (
	"database/sql"
	"fmt"
	"os"

	analytics_handler "github.com/abmid/icanvas-analytics/internal/analytics/delivery/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/spf13/viper"
)

func dbSetup(host, username, dbname, password string) *sql.DB {
	parse, err := pgx.ParseURI("postgres://" + username + ":" + password + "@" + host + ":5432/" + dbname + "?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	db := stdlib.OpenDB(parse)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(4)
	return db
}

func readConfig() *viper.Viper {
	v := viper.New()
	getEnv := os.Getenv("APP_ENV")
	if getEnv == "production" {
		v.SetConfigName("prod")
	} else {
		v.SetConfigName("dev") // name of config file (without extension)
	}
	v.SetConfigType("yaml")          // REQUIRED if the config file does not have the extension in the name
	v.AddConfigPath("../../configs") // path to look for the config file in
	err := v.ReadInConfig()          // Find and read the config file
	if err != nil {                  // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return v
}

func main() {
	// Init Config
	config := readConfig()
	// InitDB
	dbHost := config.GetString("database.host")
	dbUsername := config.GetString("database.username")
	dbName := config.GetString("database.name")
	dbPassword := config.GetString("database.password")
	db := dbSetup(dbHost, dbUsername, dbName, dbPassword)
	defer db.Close()
	// Init Config LMS
	canvasUrl := config.GetString("canvas.url")
	canvasAccessToken := config.GetString("canvas.access_token")
	gin := gin.Default()
	aUC := analytics_handler.SetupUseCase(db, canvasUrl, canvasAccessToken)
	r1 := gin.Group("v1")
	analytics_handler.NewHandler("analytics", r1, aUC)
	gin.Run(config.GetString("domain.port"))
}
