package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setUpEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Failed to load .env file, make sure you already set the env vars.")
	}
}

func main() {
	app := gin.Default()

	setUpEnv()

	SetUpDiscordBot()

	app.POST("/send", func(c *gin.Context) {
		msg := c.Query("message")

		if c.Query("tag") != "" {
			msg = "[" + c.Query("tag") + "] " + msg
		}

		if c.Query("channel_id") != "" {
			Send(c.Query("channel_id"), msg)
			c.AbortWithStatus(200)
			return
		} else {
			SendMessageToUser(os.Getenv("YUANLIN_DISCORD_ID"), msg)
			c.AbortWithStatus(200)
		}
	})

	log.Fatal(app.Run(":9087"))
}
