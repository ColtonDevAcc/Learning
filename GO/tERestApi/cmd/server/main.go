package main

import (
	"edgerestapi/internal/comment"
	transportHttp "edgerestapi/internal/transport/http"

	"edgerestapi/internal/db"
	"fmt"
)

// run - responsible for instantiating the server and starting it.

func Run() error {
	fmt.Println("starting application ....")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to database: ")
		return err

	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database: ")
		return err
	}

	fmt.Println("successfully connected to database")

	cmtService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Server(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
