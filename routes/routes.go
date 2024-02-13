package routes

import (
	"goLogin/security"
	"goLogin/service"

	"github.com/kataras/iris/v12"
)

func InitializeRoutes() {
	server := iris.Default()

	//? customer routes
	server.Post("/customer", service.CreateCustomer)
	server.Get("/customers", service.GetAllCustomer)
	server.Put("/customer/{id:int64}", service.UpdateCustomer)
	server.Delete("/customer/{id:int64}", service.DeleteCustomer)

	//? user routes
	server.Post("/user", service.CreateUser)

	//? login
	server.Post("/login", service.UserLogin)

	secureRouter := server.Party("/")
	secureRouter.Use(security.AuthMiddleware)
	secureRouter.Post("/landing",)

	server.Run(iris.Addr(":8080"))
}
