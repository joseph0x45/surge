package main

import (
	"daemon/db"
	"daemon/handler"
	"embed"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joseph0x45/goutils"
	"github.com/joseph0x45/sad"
)

var version = "debug"

//go:embed templates
var templatesFS embed.FS

var templates *template.Template

func init() {
	var err error
	templates, err = template.ParseFS(
		templatesFS,
    "*.html",
	)
	if err != nil {
		panic(err)
	}
}

func main() {
	goutils.SetAppName("daemon")
	versionFlag := flag.Bool("version", false, "Display the current version")
	generateServiceFileFlag := flag.Bool("generate-service-file", false, "Generate a service file")
	flag.Parse()
	if *versionFlag {
		fmt.Printf("Vis %s\n", version)
		return
	}
	if *generateServiceFileFlag {
		goutils.GenerateServiceFile("My awesome daemon")
		return
	}
	dbPath := goutils.Setup()
	conn := db.Connect(sad.DBConnectionOptions{
		EnableForeignKeys: true,
		DatabasePath:      dbPath,
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := chi.NewRouter()
	handler := handler.NewHandler(templates, conn, version)
	server := http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}
	handler.RegisterRoutes(r)
	log.Printf("Starting server on  http://0.0.0.0:%s", port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
