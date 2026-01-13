package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"prabogo-todo/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// 1. Initialize Database
	if err := database.InitDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// 2. Initialize Fiber (Prabogo Style)
	app := fiber.New(fiber.Config{
		AppName: "Prabogo To-Do",
	})
	app.Use(logger.New())

	// 3. Serve Static Files (JokoUI)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./views/index.html")
	})

	// 4. API Routes
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		todos, err := database.GetAllTodos()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(todos)
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		var body struct {
			Title string `json:"title"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
		}

		id := generateID()
		timeStr := time.Now().Format("15:04")
		
		if err := database.InsertTodo(id, body.Title, timeStr); err != nil {
			log.Printf("Error saving: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": "Failed to save"})
		}

		return c.JSON(fiber.Map{
			"id":        id,
			"title":     body.Title,
			"completed": false,
			"time":      timeStr,
		})
	})

	app.Put("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		if err := database.ToggleTodo(id); err != nil {
			log.Printf("Error toggling: %v", err)
		}
		return c.JSON(fiber.Map{"status": "ok"})
	})

	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		if err := database.DeleteTodo(id); err != nil {
			log.Printf("Error deleting: %v", err)
		}
		return c.SendStatus(200)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Prabogo App listening on :" + port)
	log.Fatal(app.Listen(":" + port))
}

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
