package main

import (
	"context"
	"fmt"
	"git.chirag.codes/chirag/todolist-go/db"
	"git.chirag.codes/chirag/todolist-go/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	//Loading Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could Not Load Environment Variables")
	}

	//Connecting to Database(Mongodb)
	client := db.ConncectDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	app := fiber.New()
	app.Use(cors.New())

	routes.Init(app, client)
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("You must set your 'PORT' environmental variable.")
	}
	log.Fatal(app.Listen(fmt.Sprintf(":%s", PORT)))

}
