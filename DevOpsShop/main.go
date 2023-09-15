package main

import (
	"DevOpsShop/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/register", handlers.ShowRegisterPage)
	r.POST("/register", handlers.RegisterHandler)

	r.GET("/main", handlers.ShowMainPage)

	r.GET("/login", handlers.ShowLoginPage)
	r.POST("/login", handlers.LoginHandler)
	r.POST("/logout", handlers.Logout)

	r.POST("/products/add", handlers.CreateProductHandler)
	r.GET("/products", handlers.GetAllProducts)

	http.ListenAndServe(":8080", r)
}
