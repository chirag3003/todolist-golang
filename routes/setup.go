package routes

import (
	"encoding/json"
	"git.chirag.codes/chirag/todolist-go/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var app *fiber.App
var client *mongo.Client

type HomeResponseStruct struct {
	Name   string `json:"Name"`
	Author string `json:"Author"`
}

func Init(a *fiber.App, c *mongo.Client) {
	app = a
	client = c
	Todo := client.Database("todo").Collection("List")
	app.Get("/", func(ctx *fiber.Ctx) error {
		res, err := json.Marshal(HomeResponseStruct{
			Name:   "Chirag Bhalotia",
			Author: "chirag3003",
		})
		if err != nil {
			log.Fatal("Error ")
		}
		return ctx.Send(res)
	})
	app.Get("/list", func(ctx *fiber.Ctx) error {
		var todos []models.Todo
		list, _ := Todo.Find(ctx.UserContext(), bson.D{})
		err := list.All(ctx.UserContext(), &todos)
		if err != nil {
			return err
		}
		return ctx.JSON(todos)
	})
	app.Post("/", func(ctx *fiber.Ctx) error {
		body := models.NewTodo{}
		err := ctx.BodyParser(&body)
		if err != nil {
			return err
		}
		body.SetUpdatedAt()
		body.SetCreatedAt()
		one, err := Todo.InsertOne(ctx.UserContext(), body)
		if err != nil {
			return err
		}

		return ctx.JSON(one)
	})
	app.Put("/:id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		log.Println(id)
		body := models.NewTodo{}
		err := ctx.BodyParser(&body)
		if err != nil {
			return err
		}
		body.SetUpdatedAt()
		body.SetCreatedAt()
		one, err := Todo.UpdateOne(ctx.UserContext(), struct {
			id string `json:"id"`
		}{
			id,
		}, bson.D{{"$set", body}})
		if err != nil {
			return err
		}

		return ctx.JSON(one)
	})
	app.Delete("/", func(ctx *fiber.Ctx) error {
		var id struct {
			Id string `json:"id"`
		}
		err := ctx.BodyParser(&id)
		if err != nil {
			return err
		}

		one, err := Todo.DeleteOne(ctx.UserContext(), struct {
			id string `json:"id"`
		}{
			id.Id,
		})
		if err != nil {
			return err
		}
		return ctx.JSON(one)
	})
	app.Delete("/all", func(ctx *fiber.Ctx) error {
		res, er := Todo.DeleteMany(ctx.UserContext(), bson.D{{}})
		if er != nil {
			return er
		}
		return ctx.JSON(res)
	})
}
