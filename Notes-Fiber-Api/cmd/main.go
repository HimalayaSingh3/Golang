package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func main(){
	app := fiber.New()
	log.Println("Server is Running on Port 8080")
	log.Fatal((app.Listen(":8080")))
}