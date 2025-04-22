package main

import (
	"TODO/internal/handlers"
	"TODO/storage"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load("../config/.env"); err != nil {
		panic("Ошибка загрузки файла .env: " + err.Error())
	}

	URL := os.Getenv("DATABASE_URL")
	if URL == "" {
		panic("DATABASE_URL no found")
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		panic("Переменная PORT не задана")
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

	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.SendFile("../swagger.yaml")
	})

	if err := app.Listen(":" + PORT); err != nil {
		panic("Ошибка запуска сервера: " + err.Error())
	}
}
