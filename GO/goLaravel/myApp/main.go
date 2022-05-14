package main

import (
	"myapp/data"
	"myapp/handlers"

	voo "github.com/VooDooStack/Voo"
)

type application struct {
	App *voo.Voo
	Handlers *handlers.Handlers
	Models data.Models
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
