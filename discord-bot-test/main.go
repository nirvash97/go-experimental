package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// const greenColor = "\033[32m"
// const redColor = "\033[31m"
// const defaultColor = "\033[0m"
const bot_token = "BOT_TOKEN"
const tragetChannelID = "1352174956259311689"

func main() {
	disBot, err := discordgo.New("Bot " + bot_token)
	if err != nil {
		fmt.Println("Error creating bot session:", err)
		return
	}
	// Enable the Message Content intent
	disBot.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent
	disBot.AddHandler(messageHandler)
	err = disBot.Open()
	if err != nil {
		fmt.Println("Connection Error", err)
		return

	}
	defer disBot.Close()

	// message := "TEST #0"
	// _, err = disBot.ChannelMessageSend(tragetChannelID, message)
	// if err != nil {
	// 	fmt.Printf("%s [ERROR] Send Message fail%s", redColor, defaultColor)
	// }
	fmt.Println("Bot is running. Press CTRL+C to exit.")
	select {} // Keep the bot running
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore bot messages
	if m.Author.Bot {
		return
	}

	// Log full message data
	fmt.Printf("Received message: %#v\n", m)

	// Check if message content is empty
	if m.Content == "" {
		fmt.Println("Warning: Message content is empty!")
		return
	}
	fmt.Printf("Got messgae : %s", m.Content)
	// Check if the message is in the allowed channel
	if m.ChannelID != tragetChannelID {
		return // Ignore messages from other channels
	}
	if m.Content == "!echo" {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Hello <@%s>! Your ID is %s.", m.Author.ID, m.Author.ID))
	}
}
