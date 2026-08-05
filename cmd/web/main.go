package main

import (
	"chat/pkg"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var hub *Hub
var db *gorm.DB

func main() {
	pkg.Init()

	port := os.Getenv("PORT")

	router := gin.Default()
	router.Use(cors.Default()) //TODO url server

	hub = newHub()
	go hub.run()

	db = pkg.OpenDB()

	{
		api := router.Group("/api")
		api.GET("/message", GetMessageHandler)
		api.POST("/message", CreateMessageHandler)
	}

	router.GET("/ws", serveWs)

	router.Run(":" + port)
}

func GetMessageHandler(c *gin.Context) {
	messages := []*pkg.Message{}
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

	err = db.Create(&message).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jsonData, err := json.Marshal(message)
	hub.broadcast <- jsonData

	c.JSON(http.StatusOK, message)
}
