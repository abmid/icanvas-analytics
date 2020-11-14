package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)

func main() {
	// Init Config
	fmt.Println("1. Get Config Information")
	// getEnv := os.Getenv("APP_ENV")
	viper.SetConfigType("env") // REQUIRED if the config file does not have the extension in the name
	viper.SetConfigName(".env")
	viper.AddConfigPath("../")  // path to look for the config file in
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	dbHost := viper.GetString("PG_HOST")
	dbUsername := viper.GetString("PG_USER")
	dbName := viper.GetString("PG_DBNAME")
	dbPassword := viper.GetString("PG_PASSWORD")
	dbPort := viper.GetString("PG_PORT")
	m, err := migrate.New(
		"file://migrations",
		"postgres://"+dbUsername+":"+dbPassword+"@"+dbHost+":"+dbPort+"/"+dbName+"?sslmode=disable")
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
		if err := m.Up(); err != nil && m.Up().Error() != "no change" {
			log.Fatal(err)
		}
	}
	fmt.Println("4. Finish")
}
