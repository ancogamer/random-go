package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Host struct {
	Fiber *fiber.App
}

func main() {
	// Hosts
	hosts := map[string]*Host{}
	//-----
	// API
	//-----
	api := fiber.New()
	api.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	hosts["api.localhost:3000"] = &Host{api}
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API")
	})
	//------
	// Blog
	//------
	blog := fiber.New()
	blog.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	hosts["blog.localhost:3000"] = &Host{blog}
	blog.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Blog")
	})
	//---------
	// Website
	//---------
	site := fiber.New()
	site.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	hosts["localhost:3000"] = &Host{site}
	site.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Website")
	})
	// Server
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		host := hosts[c.Hostname()]
		if host == nil {
			return c.SendStatus(fiber.StatusNotFound)
		} else {
			host.Fiber.Handler()(c.Context())
			return nil
		}
	})
	log.Fatal(app.Listen(":3000"))
}
