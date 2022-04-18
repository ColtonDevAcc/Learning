package main

import (
	"log"
	"os"

	"github.com/voodoostack/celeritas"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//init celeritas
	cel := &celeritas.Celeritas{}

	err := cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "app"
	cel.Debug = true

	app := &application{
		app: cel,
	}

	return app
}
