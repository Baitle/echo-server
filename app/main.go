package main


import (
    "log"
    "github.com/gofiber/fiber/v2"

)

func main() {
    // Create new Fiber instance
    app := fiber.New()

    // Create route for POST requests to /echo
    app.Post("/echo", func(c *fiber.Ctx) error {
        // Return the request body as-is
        return c.Send(c.Body())
    })

    // Create route for GET requests to /echo
    app.Get("/echo", func(c *fiber.Ctx) error {
        // Return any query parameter named 'message'
        message := c.Query("message", "")
        return c.SendString(message)
    })

    // Start server on port 3000
    log.Fatal(app.Listen(":3000"))
}