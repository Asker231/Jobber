package home

import (
	"github.com/Asker231/Jobber.git/pkg/tadapter"
	view "github.com/Asker231/Jobber.git/view/components"
	"github.com/gofiber/fiber/v2"
)


type HomeHandler struct{
}

func NewHomeHandler(router *fiber.App){
	h := HomeHandler{}
	router.Get("/",h.Hello)
}

func(h *HomeHandler)Hello(ctx *fiber.Ctx)error{
	component := view.Main()
	return tadapter.Render(ctx,component)
}