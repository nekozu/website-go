package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()
 
   app.Get("/",func(ctx *fiber.Ctx) error {
		    return ctx.SendFile("index.html")
   })
  
   log.Fatal(app.Listen(":3000"))
}
