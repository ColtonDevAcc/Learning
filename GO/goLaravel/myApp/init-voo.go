package main

import (
	"log"
	"myapp/data"
	"myapp/handlers"
	"os"

	voo "github.com/VooDooStack/Voo"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init voo
	voo := &voo.Voo{}
	err = voo.New(path)
	if err != nil {
		log.Fatal(err)
	}

	voo.AppName = "myapp"

	myHandlers := &handlers.Handlers{
		App: voo,
	}

	app := &application{
		App:      voo,
		Handlers: myHandlers,
	}

	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)
	myHandlers.Models = app.Models

	return app
}
