package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/dty1er/discord/bot"
)

var (
	token    = "Bot " + os.Getenv("DISCORD_TOKEN")
	clientID = os.Getenv("DISCORD_CLIENT_ID")
	botName  = "vkgdog"
	stopBot  = make(chan bool)
)

func main() {
	session, err := discordgo.New()
	session.Token = token
	if err != nil {
		panic(err)
	}

	dog := bot.New(clientID)

	session.AddHandler(dog.OnMessageCreate)
	if err = session.Open(); err != nil {
		panic(err)
	}
	defer session.Close()

	fmt.Println("server started")
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
