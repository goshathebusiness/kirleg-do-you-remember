package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/goshathebusiness/kirleg-do-you-remember/pkg/cfg"
	pg "github.com/goshathebusiness/kirleg-do-you-remember/pkg/db/postgres"
	"github.com/goshathebusiness/kirleg-do-you-remember/pkg/kdyr/router"
	"github.com/goshathebusiness/kirleg-do-you-remember/pkg/kdyr/services"
)

const defaultConfigPath = "./cmd/kdyr/config.yaml"

func main() {
	var configPath string
	flag.StringVar(
		&configPath,
		"config-path",
		defaultConfigPath,
		"provides a path to configuration file with extension .yaml")

	flag.Parse()

	var config config

	err := cfg.UnmarshalYAML(configPath, &config)
	if err != nil {
		log.Fatalf("Failed to load config file. Err: %s", err.Error())
	}

	db, err := pg.NewDB(config.DB.URL(), config.DB.IsolationLevel)
	if err != nil {
		log.Fatalf("Failed to connect to db. Err: %s", err.Error())
	}
	log.Println("Successfully connected to db")

	svc := services.NewServices(db)

	r := router.NewRouter(svc)

	log.Printf("Server is starting on %s", config.WebServer.Addr)
	log.Fatal(http.ListenAndServe(config.WebServer.Addr, r))

}
