package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/abmid/icanvas-analytics/internal/analyticsjob/delivery/job"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
)

func dbSetup(host, username, dbname, password string) *sql.DB {
	parse, err := pgx.ParseURI("postgres://" + username + ":" + password + "@" + host + ":5432/" + dbname + "?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	db := stdlib.OpenDB(parse)
	db.SetMaxOpenConns(90)
	db.SetMaxIdleConns(4)
	return db
}

func CronSet() {
	fmt.Println("1. Get Config Information")
	getEnv := os.Getenv("APP_ENV")
	if getEnv == "production" {
		viper.SetConfigName("prod")
	} else {
		viper.SetConfigName("dev") // name of config file (without extension)
	}
	viper.SetConfigType("yaml")          // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("../../configs") // path to look for the config file in
	err := viper.ReadInConfig()          // Find and read the config file
	if err != nil {                      // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	dbHost := viper.GetString("database.host")
	dbUsername := viper.GetString("database.username")
	dbName := viper.GetString("database.name")
	dbPassword := viper.GetString("database.password")
	db := dbSetup(dbHost, dbUsername, dbName, dbPassword)
	c := cron.New(
		cron.WithLogger(
			cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))
	job.RunScheduling(c, db, viper.GetString("canvas.url"), viper.GetString("canvas.access_token"))
	c.Run()
}

func main() {
	CronSet()
}
