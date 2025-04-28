package main

import (
	"log"
	"os"
	"test-task-TODO/internal/handlers"
	cfg "test-task-TODO/pkg/config"
	psql "test-task-TODO/storage"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg.LoadConfig() //Подключаю конфиг

	url := cfg.PsqlCfg() //Формирую url

	conn, err := psql.Open(url) //открываю коннект
	if err != nil {
		log.Fatalf("Unable to connect to db: %v\n", err)
	}
	defer conn.Close() //закрываю коннект

	h := handlers.New(conn)
	app := fiber.New()

	taskGroup := app.Group("/tasks")
	{
		taskGroup.Get("/", h.GetTasks)
		taskGroup.Post("/", h.PostTask)
		taskGroup.Put("/:id", h.PutTask)
		taskGroup.Delete("/:id", h.DeleteTask)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Printf("port not found")
	}

	log.Println("starting server on port", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal("server startup error: " + err.Error())
	}
}
