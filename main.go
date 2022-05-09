package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

//Note BotToken will be set as an environment variable
var (
	GuildID  = flag.String("guild", "", "Test guild ID")
	BotToken = flag.String("token", "", "Bot token")
)

func init() {
	flag.Parse()
}

func main() {

	//initialize the discord bot
	s, err := discordgo.New("Bot " + *BotToken)

	if err != nil {
		log.Fatalf("Looks like you entered the wrong parameters for your bot")
	}

	fmt.Println(s)

}
