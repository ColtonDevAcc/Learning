package main

import (
	"myapp/handlers"

	voo "github.com/VooDooStack/Voo"
)

type application struct {
	App      *voo.Voo
	Handlers *handlers.Handlers
}

func main() {
	v := initApplication()
	v.App.ListenAndServe()

}
