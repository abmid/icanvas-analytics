package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/abmid/icanvas-analytics/internal/validation"
	analytics_handler "github.com/abmid/icanvas-analytics/pkg/analytics/delivery/http"
	auth_login_handler "github.com/abmid/icanvas-analytics/pkg/auth/login/delivery/http"
	auth_logout_handler "github.com/abmid/icanvas-analytics/pkg/auth/logout/delivery/http"
	auth_register_handler "github.com/abmid/icanvas-analytics/pkg/auth/register/delivery/http"
	canvas_account_handler "github.com/abmid/icanvas-analytics/pkg/canvas/account/delivery/http"
	setting_handler "github.com/abmid/icanvas-analytics/pkg/setting/delivery/http"
	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/spf13/viper"
)

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

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

/**
* Get Root Path Project
* @return string
 */
func GetRootPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// this process to get root project, because binary file main in this location today
	rootPath := filepath.Join(dir, "../..")
	return rootPath
}

func readConfig(rootPath string) *viper.Viper {
	v := viper.New()
	getEnv := os.Getenv("APP_ENV")
	if getEnv == "production" {
		v.SetConfigName("prod")
	} else {
		v.SetConfigName("dev") // name of config file (without extension)
	}
	v.SetConfigType("yaml")                // REQUIRED if the config file does not have the extension in the name
	v.AddConfigPath(rootPath + "/configs") // path to look for the config file in
	err := v.ReadInConfig()                // Find and read the config file
	if err != nil {                        // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return v
}

func main() {
	// Get Root Path
	rootPath := GetRootPath()
	// Init Config
	config := readConfig(rootPath)
	// InitDB
	dbHost := config.GetString("database.host")
	dbUsername := config.GetString("database.username")
	dbName := config.GetString("database.name")
	dbPassword := config.GetString("database.password")
	db := dbSetup(dbHost, dbUsername, dbName, dbPassword)
	defer db.Close()
	// Init JWT Key
	JWTKey := config.GetString("security.secret_key")
	// Init Route (echo)
	e := echo.New()
	// Remove trailing slash (/)
	e.Pre(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{RedirectCode: http.StatusPermanentRedirect}))
	// Add Custom Validation
	validation.AlphaValidation(e)
	// Use Logger
	e.Use(middleware.Logger())
	// Use Cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("../../web/app/dist/*.html")),
	}
	e.Renderer = renderer

	e.Static("/", "../../web/app/dist")
	// Named route "foobar"
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
	})
	/**
	* Route v1
	 */
	r1 := e.Group("/v1")
	// Auth
	// login
	loginUC := auth_login_handler.SetupUseCase(db)
	auth_login_handler.NewHandler("/auth", r1, JWTKey, loginUC)
	// logout
	auth_logout_handler.NewHandler("/auth", r1)
	// register
	registerUC := auth_register_handler.SetupUseCase(db)
	auth_register_handler.NewHandler("/auth", r1, registerUC)

	// Analytics Course
	aUC := analytics_handler.SetupUseCase(db)
	analytics_handler.NewHandler("/analytics", r1, JWTKey, aUC)

	// Setting
	settingUC := setting_handler.SetupUseCase(db)
	setting_handler.NewHandler("/settings", r1, JWTKey, settingUC)

	// Canvas
	canvasAccountUC := canvas_account_handler.SetupUseCase(settingUC)
	canvas_account_handler.NewHandler("/canvas", r1, JWTKey, canvasAccountUC)

	// Start Server
	e.Start(config.GetString("domain.port"))
}
