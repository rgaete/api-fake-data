package models

import (
	"github.com/gofiber/fiber"
)

func GetScanItems(c *fiber.Ctx) {
	c.Send("All Books")
}

func GetScanItem(c *fiber.Ctx) {
	c.Send("Single Book")
}
