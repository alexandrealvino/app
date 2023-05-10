package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	// Servir a página HTML quando a rota "/" for acessada
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("/home/alexandre/Desktop/index.html")
	})

	// Iniciar o servidor na porta 3000
	app.Listen("0.0.0.0:4321")
}

//08007210026
