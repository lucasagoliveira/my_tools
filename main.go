package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"fmt"
)

func main() {
    app := fiber.New()

    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    log.Fatal(app.Listen(":3000"))

    fmt.Print("hi")
}