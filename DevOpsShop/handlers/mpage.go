package handlers

import (
	"DevOpsShop/database"
	"DevOpsShop/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ShowMainPage(c *gin.Context) {

	var products []models.Product

	tokenString, err := c.Cookie("token")
	if err != nil {
		c.HTML(http.StatusOK, "mainpage.html", gin.H{
			"Products":        products,
			"IsAuthenticated": false,
		})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte("secret"), nil
	})
	if err != nil {
		c.HTML(http.StatusOK, "mainpage.html", gin.H{
			"Products":        products,
			"IsAuthenticated": false,
		})
		return
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		c.HTML(http.StatusOK, "mainpage.html", gin.H{
			"Products":        products,
			"IsAuthenticated": false,
		})
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	db.Find(&products)

	c.HTML(http.StatusOK, "mainpage.html", gin.H{
		"Products":        products,
		"IsAuthenticated": true,
	})
}

func GetAllProducts(c *gin.Context) {
	db, err := database.ConnectDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var products []models.Product
	err = db.Find(&products).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}
func CreateProductHandler(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = db.Create(&product).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": product})
}

func GetUserIDFromToken(c *gin.Context) (uint, error) {
	tokenString, err := c.Cookie("token")
	if err != nil {
		return 0, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("your-secret-key"), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("user ID not found in token claims")
	}

	fmt.Println("User ID:", uint(userID))

	return uint(userID), nil
}

func GetUserCartItems(c *gin.Context) {
	userID, err := GetUserIDFromToken(c)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var cartItems []models.CartItem
	db, err := database.ConnectDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err := db.Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, "cart.html", gin.H{
		"CartItems": cartItems,
	})
}

func AddToCart(c *gin.Context) {

	db, err := database.ConnectDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	userID, err := GetUserIDFromToken(c)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var requestData struct {
		ProductID uint `json:"product_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cartItem := models.CartItem{
		UserID:    userID,
		ProductID: requestData.ProductID,
		Quantity:  1,
	}

	if err := db.Create(&cartItem).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added to cart successfully"})
}
