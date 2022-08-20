package main

import (
	"fmt"
	"os"
	"log"

	"codebrains.io/todo-list/database"
	todo "codebrains.io/todo-list/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func initDatabase() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Some error occured. Err: %s", envErr)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to Database")
	}
	fmt.Println("Database connection successful")
	database.DBConn.AutoMigrate(&todo.Todo{})
	fmt.Println("Database migrated")
}
func setupRoutes(app *fiber.App) {
	app.Get("/todos", todo.GetTodos)
	app.Get("/todos/:id", todo.GetTodoById)
	app.Post("/todos", todo.CreateTodo)
	app.Put("/todos/:id", todo.UpdateTodo)
	app.Delete("/todos/:id", todo.DeleteTodo)
}

func main() {
	app := fiber.New()
	initDatabase()
	app.Get("/", helloWorld)
	setupRoutes(app)
	app.Listen(":8000")
}
