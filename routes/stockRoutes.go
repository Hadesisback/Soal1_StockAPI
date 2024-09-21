package routes

import (
    "github.com/gofiber/fiber/v2"
    "stock-api/handler"

)

func SetupRoutes(app *fiber.App) {
    app.Get("/stocks", handler.GetAllStocks)
    app.Post("/stocks", handler.CreateStock)
    app.Get("/stocks/:id", handler.GetStock)
    app.Put("/stocks/:id", handler.UpdateStock)
    app.Delete("/stocks/:id", handler.DeleteStock)
}
