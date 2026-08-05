package main

import (
	"chat/pkg"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	pkg.Init()

	port := os.Getenv("PORT")

	r := gin.Default()

	{
		api := r.Group("/api")
		api.GET("/message", GetMessageHandler)
		api.POST("/message", CreateMessageHandler)
	}

	log.Fatalln(r.Run(":" + port))
}

func GetMessageHandler(c *gin.Context) {
	messages := []*pkg.Message{}
	db := pkg.OpenDB()
	err := db.Scopes(pkg.Paginate(c.Request)).Find(&messages).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}

func CreateMessageHandler(c *gin.Context) {
	message := &pkg.Message{}
	err := c.ShouldBindJSON(message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := pkg.OpenDB()
	err = db.Create(&message).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}
