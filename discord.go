package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func SetUpDiscordBot() *discordgo.Session {

	log.Print("Opening discord session ...")

	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))

	if err != nil {
		log.Fatal(err)
	}

	err = discord.Open()

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Discord session opened successfully.")

	return discord
}

func Send(channelID string, text string) {
	discord := SetUpDiscordBot()
	_, err := discord.ChannelMessageSend(channelID, text)
	if err != nil {
		fmt.Println("出大事啦:", err)
	}
}

func SendMessageToUser(userID string, text string) {
	discord := SetUpDiscordBot()
	st, err := discord.UserChannelCreate(userID)
	if err != nil {
		fmt.Print(err)
	}
	_, err = discord.ChannelMessageSend(st.ID, text)
	if err != nil {
		fmt.Print(err)
	}
}

func DiscordMessageEmbed(channelID string, a discordgo.MessageEmbed) {
	discord := SetUpDiscordBot()
	_, err := discord.ChannelMessageSendEmbed(channelID, &a)
	if err != nil {
		fmt.Println("出大事啦:", err)
		return
	}
}
