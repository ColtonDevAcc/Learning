package main

import (
	"context"
	"edgerestapi/internal/comment"
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
	fmt.Println(cmtService.GetComment(
		context.Background(),
		"5e9f8f8f-f9b9-4f7b-b8b8-f9b9f9b9f9b9",
	))

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
