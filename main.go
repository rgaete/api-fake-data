package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Shopify/sarama"
	"github.com/brianvoe/gofakeit"
	"github.com/gofiber/fiber/v2"
)

// ScanItem struct
type ScanItem struct {
	Tienda        int    `json:"tienda"`
	Departamento  int    `json:"departamento"`
	CodigoDeBarra string `json:"codigo_de_barra"`
	Cantidad      int    `json:"cantidad"`
	Precio        int    `json:"precio"`
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := fiber.New()
	api := app.Group("/api/v1") // /api

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello Heroku")
	})
	api.Get("/scanitems", createFakeScanItem)
	api.Get("/products/:code", getProductByCode)

	app.Server().ListenAndServe(":" + port)

}

func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func PushCommentToQueue(topic string, message []byte) error {

	brokersUrl := []string{"localhost:9092"}
	producer, err := ConnectProducer(brokersUrl)
	if err != nil {
		return err
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

	return nil
}

type Product []struct {
	CodigoDeBarra string `json:"codigo_de_barra"`
	Nombre        string `json:"nombre"`
	Departamento  string `json:"departamento"`
}

var BarCodes = []string{"ABC-500-000", "ABC-501-000", "ABC-502-000", "ABC-503-000", "ABC-504-000", "ABC-505-000", "ABC-506-000", "ABC-507-000", "ABC-508-000", "ABC-509-000", "ABC-510-000", "ABC-511-000", "ABC-512-000", "ABC-513-000", "ABC-514-000", "ABC-515-000", "ABC-516-000", "ABC-517-000", "ABC-518-000", "ABC-519-000", "ABC-520-000", "ABC-521-000", "ABC-522-000", "ABC-523-000", "ABC-524-000", "ABC-525-000", "ABC-526-000", "ABC-527-000", "ABC-528-000", "ABC-529-000", "ABC-530-000", "ABC-531-000", "ABC-532-000", "ABC-533-000", "ABC-534-000", "ABC-535-000", "ABC-536-000", "ABC-537-000", "ABC-538-000", "ABC-539-000", "ABC-540-000", "ABC-541-000", "ABC-542-000", "ABC-543-000", "ABC-544-000", "ABC-545-000", "ABC-546-000", "ABC-547-000", "ABC-548-000", "ABC-549-000", "ABC-550-000", "ABC-551-000", "ABC-552-000", "ABC-553-000", "ABC-554-000", "ABC-555-000", "ABC-556-000", "ABC-557-000", "ABC-558-000", "ABC-559-000", "ABC-560-000", "ABC-561-000", "ABC-562-000", "ABC-563-000", "ABC-564-000", "ABC-565-000", "ABC-566-000", "ABC-567-000", "ABC-568-000", "ABC-569-000", "ABC-570-000", "ABC-571-000", "ABC-572-000", "ABC-573-000", "ABC-574-000", "ABC-575-000", "ABC-576-000", "ABC-577-000", "ABC-578-000", "ABC-579-000", "ABC-580-000", "ABC-581-000", "ABC-582-000", "ABC-583-000", "ABC-584-000", "ABC-585-000", "ABC-586-000", "ABC-587-000", "ABC-588-000", "ABC-589-000", "ABC-590-000", "ABC-591-000", "ABC-592-000", "ABC-593-000", "ABC-594-000", "ABC-595-000", "ABC-596-000", "ABC-597-000", "ABC-598-000", "ABC-599-000", "ABC-600-000", "ABC-601-000", "ABC-602-000", "ABC-603-000", "ABC-604-000", "ABC-605-000", "ABC-606-000", "ABC-607-000", "ABC-608-000", "ABC-609-000", "ABC-610-000", "ABC-611-000", "ABC-612-000", "ABC-613-000", "ABC-614-000", "ABC-615-000", "ABC-616-000", "ABC-617-000", "ABC-618-000", "ABC-619-000", "ABC-620-000", "ABC-621-000", "ABC-622-000", "ABC-623-000", "ABC-624-000", "ABC-625-000", "ABC-626-000", "ABC-627-000", "ABC-628-000", "ABC-629-000", "ABC-630-000", "ABC-631-000", "ABC-632-000", "ABC-633-000", "ABC-634-000", "ABC-635-000", "ABC-636-000", "ABC-637-000", "ABC-638-000", "ABC-639-000", "ABC-640-000", "ABC-641-000", "ABC-642-000", "ABC-643-000", "ABC-644-000", "ABC-645-000", "ABC-646-000", "ABC-647-000", "ABC-648-000", "ABC-649-000", "ABC-650-000", "ABC-651-000", "ABC-652-000", "ABC-653-000", "ABC-654-000", "ABC-655-000", "ABC-656-000", "ABC-657-000", "ABC-658-000", "ABC-659-000", "ABC-660-000", "ABC-661-000", "ABC-662-000", "ABC-663-000", "ABC-664-000", "ABC-665-000", "ABC-666-000", "ABC-667-000", "ABC-668-000", "ABC-669-000", "ABC-670-000", "ABC-671-000", "ABC-672-000", "ABC-673-000", "ABC-674-000", "ABC-675-000", "ABC-676-000", "ABC-677-000", "ABC-678-000", "ABC-679-000", "ABC-680-000", "ABC-681-000", "ABC-682-000", "ABC-683-000", "ABC-684-000", "ABC-685-000", "ABC-686-000", "ABC-687-000", "ABC-688-000", "ABC-689-000", "ABC-690-000", "ABC-691-000", "ABC-692-000", "ABC-693-000", "ABC-694-000", "ABC-695-000", "ABC-696-000", "ABC-697-000", "ABC-698-000", "ABC-699-000", "ABC-700-000", "ABC-701-000", "ABC-702-000", "ABC-703-000", "ABC-704-000", "ABC-705-000", "ABC-706-000", "ABC-707-000", "ABC-708-000", "ABC-709-000", "ABC-710-000", "ABC-711-000", "ABC-712-000", "ABC-713-000", "ABC-714-000", "ABC-715-000", "ABC-716-000", "ABC-717-000", "ABC-718-000", "ABC-719-000", "ABC-720-000", "ABC-721-000", "ABC-722-000", "ABC-723-000", "ABC-724-000", "ABC-725-000", "ABC-726-000", "ABC-727-000", "ABC-728-000", "ABC-729-000", "ABC-730-000", "ABC-731-000", "ABC-732-000", "ABC-733-000", "ABC-734-000", "ABC-735-000", "ABC-736-000", "ABC-737-000", "ABC-738-000", "ABC-739-000", "ABC-740-000", "ABC-741-000", "ABC-742-000", "ABC-743-000", "ABC-744-000", "ABC-745-000", "ABC-746-000", "ABC-747-000", "ABC-748-000", "ABC-749-000", "ABC-750-000", "ABC-751-000", "ABC-752-000", "ABC-753-000", "ABC-754-000", "ABC-755-000", "ABC-756-000", "ABC-757-000", "ABC-758-000", "ABC-759-000", "ABC-760-000", "ABC-761-000", "ABC-762-000", "ABC-763-000", "ABC-764-000", "ABC-765-000", "ABC-766-000", "ABC-767-000", "ABC-768-000", "ABC-769-000", "ABC-770-000", "ABC-771-000", "ABC-772-000", "ABC-773-000", "ABC-774-000", "ABC-775-000", "ABC-776-000", "ABC-777-000", "ABC-778-000", "ABC-779-000", "ABC-780-000", "ABC-781-000", "ABC-782-000", "ABC-783-000", "ABC-784-000", "ABC-785-000", "ABC-786-000", "ABC-787-000", "ABC-788-000", "ABC-789-000", "ABC-790-000", "ABC-791-000", "ABC-792-000", "ABC-793-000", "ABC-794-000", "ABC-795-000", "ABC-796-000", "ABC-797-000", "ABC-798-000", "ABC-799-000", "ABC-800-000"}

func getProductByCode(c *fiber.Ctx) error {
	fmt.Println("PRODCUT : " + c.Params("code"))
	jsonFile, err := os.Open("products.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")

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
				"Producto": products[i],
			})
			return err
		}
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	return err
}

func createFakeScanItem(c *fiber.Ctx) error {
	scanItem := new(ScanItem)

	gofakeit.Seed(0)

	scanItem.Cantidad = gofakeit.Number(1, 10)
	scanItem.Departamento = gofakeit.Number(1, 10)
	scanItem.Precio = gofakeit.Number(500, 10000)
	scanItem.CodigoDeBarra = BarCodes[gofakeit.Number(0, 100)]
	scanItem.Tienda = gofakeit.Number(1, 3)

	err := c.JSON(&fiber.Map{
		"scanItem": scanItem,
	})
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating product",
		})
		return err
	}
	return err
}
