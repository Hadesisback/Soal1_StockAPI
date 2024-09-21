package main

import (
    "stock-api/database"
    "stock-api/routes"
    "github.com/gofiber/fiber/v2"
)

func main() {
    // Initialize the database
    database.InitDatabase()

    // Initialize Fiber
    app := fiber.New()

    // Setup routes
    routes.SetupRoutes(app)

    // Start the server
    app.Listen(":8000")
}