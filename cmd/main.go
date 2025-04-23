package main

import (
	"TODO/internal/handlers"
	"TODO/internal/storage"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	URL := os.Getenv("DATABASE_URL")
	if URL == "" {
		panic("DATABASE_URL not found")
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		panic("PORT not found")
	}

	db, err := storage.Open(URL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	h := handlers.New(db)
	app := fiber.New()

	taskGroup := app.Group("/task")
	{
		taskGroup.Get("/", h.GetTasks)
		taskGroup.Post("/", h.PostTask)
		taskGroup.Put("/:id", h.PutTask)
		taskGroup.Delete("/:id", h.DeleteTask)
	}

	log.Println("Starting server on port", PORT)
	if err := app.Listen(":" + PORT); err != nil {
		panic("Server startup error: " + err.Error())
	}
}
