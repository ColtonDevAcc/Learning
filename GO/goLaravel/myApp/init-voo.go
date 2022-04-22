package main

import (
	"log"
	"os"

	voo "github.com/VooDooStack/Voo"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//init voo
	voo := &voo.Voo{}
	err = voo.New(path)
	if err != nil {
		log.Fatal(err)
	}

	voo.AppName = "myapp"
	voo.Debug = true

	app := &application{
		App: voo,
	}

	return app
}
