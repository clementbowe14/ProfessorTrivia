package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type Config struct {
	BotToken string `json:"botToken"`
	ClientID string `json:"clientID"`
}

var token string
var conf Config

func init() {

	// flag.StringVar(&token, "token", "", "Bot token")
	// flag.StringVar(&message, "message", "", "This is the first message: ")
	// flag.Parse();
	file, err := os.Open("env-config.json")
	if err != nil {
		panic(err)
	}

	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&conf)

	if err != nil {
		panic("there was an error while trying to decode the config file")
	}
	token = "Bot token" + conf.BotToken
}

func main() {

	//check if token provided
	if token == "" {
		panic("Missing token")
	}

	dg, err := discordgo.New("Bot " + conf.BotToken)
	if err != nil {
		panic("Discord go was not able to create your session")
	}

	//add handler
	dg.AddHandler(discordMessageCreate)

	//open a websocket connection to discord
	err = dg.Open()
	defer dg.Close()

	fmt.Println("You now have a socket connected to the server!")

	//runs process until interrupted by ctrl + c or contains an invalid command
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT)
	<-sc

	if err != nil {
		panic("Could not connect to discord")
	}

}

//listens to when a message is created
func discordMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("New Message!")
}
