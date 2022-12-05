package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type todo struct {
	Id   int    `json:"id"`
	Todo string `json:"todo"`
}

var lastId int = 0

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://todo.list",
		AllowMethods: "GET, POST, DELETE, HEAD, OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	var todolist []todo

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./index.html")
	})

	app.Get("/todo", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"status": "success",
			"result": fiber.Map{
				"todo": todolist,
			},
		})
	})

	app.Post("/todo", func(c *fiber.Ctx) error {
		newTodo := c.Query("todo", "soemting error")
		lastId += 1

		todolist = append(todolist, todo{Id: lastId, Todo: newTodo})

		return c.JSON(fiber.Map{
			"status": "create success",
			"result": fiber.Map{
				"item": newTodo,
				"id":   lastId,
			},
		})
	})

	app.Delete("/todo", func(c *fiber.Ctx) error {
		id := c.Query("id", "0")
		idInt, _ := strconv.Atoi(id)

		delSuccess := false

		var todo string

		for i, v := range todolist {
			todo = v.Todo
			if v.Id == idInt {
				todolist = append(todolist[:i], todolist[i+1:]...)
				delSuccess = true
			}
		}

		if !delSuccess {
			return c.JSON(fiber.Map{
				"status": "delete failed",
				"result": fiber.Map{
					"id": idInt,
				},
			})
		}

		return c.JSON(fiber.Map{
			"status": "delete success",
			"result": fiber.Map{
				"item": todo,
				"id":   idInt,
			},
		})
	})

	log.Fatal(app.Listen(":3000"))
}
