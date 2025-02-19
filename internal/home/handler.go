package home

import "github.com/gofiber/fiber/v3"


type HomeHandler struct{
}

func NewHomeHandler(router *fiber.App){
	h := HomeHandler{}
	router.Get("/",h.Hello)
}

func(h *HomeHandler)Hello(ctx fiber.Ctx)error{
	return ctx.SendString("Hello go")
}