package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func getFoods(c *fiber.Ctx) error {
	return c.JSON(foods)
}

func getFoodByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}
	var editFood EditFood
	if err := c.BodyParser(&editFood); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	for i, food := range foods {
		if food.ID == uint(id) {
			foods[i].Name = editFood.Name
			foods[i].Price = editFood.Price
			return c.JSON(foods[i])
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Food not found")
}

func createFood(c *fiber.Ctx) error {
	var food Food
	if err := c.BodyParser(&food); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	foods = append(foods, food)
	return c.JSON(food)
}

func updateFoodByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}
	var editFood EditFood
	if err := c.BodyParser(&editFood); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	for i, food := range foods {
		if food.ID == uint(id) {
			foods[i].Name = editFood.Name
			foods[i].Price = editFood.Price
			return c.JSON(foods[i])
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Food not found")
}

func deleteFoodByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}
	for i, food := range foods {
		if food.ID == uint(id) {
			foods = append(foods[:i], foods[i+1:]...)
			return c.SendString("Food deleted")
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Food not found")
}

func main() {
	fmt.Println("hello world")

	// fiber instance
	app := fiber.New()

	// routes
	app.Get("/", helloWorld)
	app.Get("/foods", getFoods)
	app.Get("/foods/:id", getFoodByID)
	app.Post("/foods", createFood)
	app.Put("/foods/:id", updateFoodByID)
	app.Delete("/foods/:id", deleteFoodByID)

	app.Get("/info", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg":     "hello world 🚀",
			"go":      "fiber 🥦",
			"boolean": true,
			"number":  1234,
		})
	})

	// app listen at port 3000
	app.Listen(":3000")
}
type Food struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

type EditFood struct {
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

var foods = []Food{
	{ID: 1, Name: "ต้มยำกุ้ง", Price: 140},
	{ID: 2, Name: "ไก่ทอด", Price: 100},
	{ID: 3, Name: "ก๋วยเตี๋ยว", Price: 30},
	{ID: 4, Name: "เบอร์เกอร์", Price: 149},
}