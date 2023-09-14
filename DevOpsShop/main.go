package main

import (
	"DevOpsShop/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/register", handlers.ShowRegisterPage)
	r.POST("/register", handlers.RegisterHandler)

	r.GET("/login", handlers.ShowLoginPage)
	r.POST("/login", handlers.LoginHandler)
	r.POST("/logout", handlers.Logout)

	http.ListenAndServe(":8080", r)
}
