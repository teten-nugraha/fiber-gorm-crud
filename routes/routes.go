package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/golang-crud-fiber/controller"
)

func Setup(server *fiber.App) {

	api := server.Group("/api/v1/")

	api.Get("/status", func(c *fiber.Ctx) error {
		return c.SendString("Service already UP 👋!")
	})


	// PRODUCT ROUTEs
	products := api.Group("products")
	products.Get("/", controller.GetProducts)
	products.Post("/", controller.SaveProducts)
	products.Get("/:id", controller.FindProductById)
	products.Put("/updateProduct/:id", controller.UpdateProduct)
	products.Get("/findByNama/:kode_barang", controller.FindProductsByKodeBarang)
	products.Delete("/:id", controller.DeleteProductById)
}
