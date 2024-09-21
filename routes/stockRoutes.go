package routes

import (
    "github.com/gofiber/fiber/v2"
    "stock-api/database"
    "stock-api/models"
    "net/http"

)

func SetupRoutes(app *fiber.App) {
    app.Get("/stocks", GetAllStocks)
    app.Post("/stocks", CreateStock)
    app.Get("/stocks/:id", GetStock)
    app.Put("/stocks/:id", UpdateStock)
    app.Delete("/stocks/:id", DeleteStock)
}

func GetAllStocks(c *fiber.Ctx) error {
    var stocks []models.Stock
    database.DB.Find(&stocks)
    return c.JSON(stocks)
}

func CreateStock(c *fiber.Ctx) error {
    var stock models.Stock

    // Decode the JSON body into the stock variable
    if err := c.BodyParser(&stock); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Failed to parse request body"+ err.Error(),
        })
    }

    // Save to database (replace with your database logic)
    if result := database.DB.Create(&stock); result.Error != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create stock",
        })
    }

    return c.Status(http.StatusCreated).JSON(stock)
}

func GetStock(c *fiber.Ctx) error {
    id := c.Params("id")
    var stock models.Stock
    if err := database.DB.First(&stock, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Stock not found"})
    }
    return c.JSON(stock)
}

func UpdateStock(c *fiber.Ctx) error {
    id := c.Params("id")
    var stock models.Stock

    // Decode the JSON body into the stock variable
    if err := c.BodyParser(&stock); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Failed to parse request body",
        })
    }

    // Update the stock in the database (replace with your database logic)
    if result := database.DB.Model(&stock).Where("id = ?", id).Updates(stock); result.Error != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update stock",
        })
    }

    return c.JSON(stock)
}

func DeleteStock(c *fiber.Ctx) error {
    id := c.Params("id")
    var stock models.Stock
    if err := database.DB.Delete(&stock, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Stock not found"})
    }
    return c.SendString("Stock deleted")
}
