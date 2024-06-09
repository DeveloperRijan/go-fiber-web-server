package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Home page
func homePage(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

// user
type User struct {
	Id   string
	Name string
	Age  int8
}

func UserHandler(c *fiber.Ctx) error {
	user := User{
		Id:   "1234",
		Name: "John Doe",
		Age:  30,
	}

	u, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(u))
	return c.SendString(string(u))
}

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Test App",
	})

	//set static assets path
	app.Static("/assets", "./public")

	//routes
	app.Get("/", homePage)
	app.Get("/api/user", UserHandler)

	app.Listen(":8000")
}
