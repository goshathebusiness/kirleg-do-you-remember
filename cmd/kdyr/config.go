package main

import (
	"github.com/goshathebusiness/kirleg-do-you-remember/pkg/cfg"
	"github.com/goshathebusiness/kirleg-do-you-remember/pkg/db"
)

type config struct {
	WebServer *cfg.WebServer `yaml:"webServer"`
	DB        *db.Config     `yaml:"db"`
}
