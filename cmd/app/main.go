package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/abmid/icanvas-analytics/internal/validation"
	analytics_handler "github.com/abmid/icanvas-analytics/pkg/analytics/delivery/http"
	auth_login_handler "github.com/abmid/icanvas-analytics/pkg/auth/login/delivery/http"
	auth_register_handler "github.com/abmid/icanvas-analytics/pkg/auth/register/delivery/http"
	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"

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
	JWTKey := config.GetString("security.secret_key")
	// Init Route (echo)
	e := echo.New()
	validation.AlphaValidation(e)
	e.Use(middleware.Logger())
	/**
	* Route v1
	 */
	r1 := e.Group("/v1")
	// Auth
	loginUC := auth_login_handler.SetupUseCase(db)
	auth_login_handler.NewHandler("/auth", r1, JWTKey, loginUC)
	registerUC := auth_register_handler.SetupUseCase(db)
	auth_register_handler.NewHandler("/auth", r1, registerUC)
	// Analytics Course
	aUC := analytics_handler.SetupUseCase(db, canvasUrl, canvasAccessToken)
	analytics_handler.NewHandler("/analytics", r1, JWTKey, aUC)

	// Start Server
	e.Start(config.GetString("domain.port"))
}
