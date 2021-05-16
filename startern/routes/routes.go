package routes

import (
"github.com/woodencoder/startern-go/core/controllers"

"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/api/vacancies", controllers.Vacancies)
}
