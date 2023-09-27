package main

import (
	"DevOpsShop/handlers"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/register", handlers.ShowRegisterPage)
	r.POST("/register", handlers.RegisterHandler)

	r.GET("/", handlers.ShowMainPage)

	r.GET("/login", handlers.ShowLoginPage)
	r.POST("/login", handlers.LoginHandler)
	r.POST("/logout", handlers.Logout)

	r.POST("/products/add", handlers.CreateProductHandler)
	r.GET("/products", handlers.GetAllProducts)

	port := ":8080"
	if len(os.Args) > 1 {
		port = ":" + os.Args[1]
	}
	r.Run(port)
}
