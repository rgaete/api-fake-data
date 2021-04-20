package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gofiber/fiber/v2"
)

// ScanItem struct
type ScanItem struct {
	Tienda        int    `json:"tienda"`
	CodigoDeBarra string `json:"codigo_de_barra"`
	Cantidad      int    `json:"cantidad"`
	Precio        int    `json:"precio"`
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app := fiber.New()
	api := app.Group("/api/v1") // /api

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Flink & Kafka Workshop")
	})

	api.Get("/products/:code", getProductByCode)

	app.Listen(":" + port)

}

type Product []struct {
	CodigoDeBarra string `json:"codigo_de_barra"`
	Nombre        string `json:"nombre"`
	Departamento  string `json:"departamento"`
}

func getProductByCode(c *fiber.Ctx) error {
	jsonFile, err := os.Open("products.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var products Product

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &products)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(products); i++ {
		if products[i].CodigoDeBarra == c.Params("code") {
			err := c.JSON(&fiber.Map{
				"producto": products[i],
			})
			return err
		}
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	return err
}
