package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)

func main() {
	// Init Config
	fmt.Println("1. Get Config Information")
	getEnv := os.Getenv("APP_ENV")
	if getEnv == "production" {
		viper.SetConfigName("prod")
	} else {
		viper.SetConfigName("dev") // name of config file (without extension)
	}
	viper.SetConfigType("yaml")       // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("../configs") // path to look for the config file in
	err := viper.ReadInConfig()       // Find and read the config file
	if err != nil {                   // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	dbHost := viper.GetString("database.host")
	dbUsername := viper.GetString("database.username")
	dbName := viper.GetString("database.name")
	dbPassword := viper.GetString("database.password")
	m, err := migrate.New(
		"file://migrations",
		"postgres://"+dbUsername+":"+dbPassword+"@"+dbHost+":5432/"+dbName+"?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("2. Database Connection Successfull")
	flagStatus := flag.String("status", "anonymous", "to get migration type")
	flagForce := flag.Int("force", 0, "Force Fix this")
	flag.Parse()
	fmt.Println("3. Migrate database status ", *flagStatus)
	if *flagForce != 0 {
		if err := m.Force(*flagForce); err != nil {
			log.Fatal(err)
		}
	}
	if *flagStatus == "down" {
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("4. Finish")
}
