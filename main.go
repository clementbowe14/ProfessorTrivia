package main

import (
	"github.com/clementbowe14/ProfessorTrivia/game"
	"github.com/clementbowe14/ProfessorTrivia/handler"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"github.com/bwmarrin/discordgo"
)

type Config struct {
	BotToken string `json:"botToken"`
	ClientID string `json:"clientID"`
}
var(
	token string
	conf Config
	g game.Game
)

func init() {

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
	dg.Open()

	if err != nil {
		panic("failed to open Websocket connection")
	}

	defer dg.Close()

	fmt.Println("You now have a socket connected to the server!")
	g = game.New()
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
	msg := m.Content
	comm := "!prof"
	mentions := m.Message.Mentions

	if strings.HasPrefix(msg, comm) {
		words := strings.Split(msg, " ")

		switch words[1] {
		case "help":
			handler.HelpCommand(s, m)
		case "addplayers":
			handler.AddPlayers(s, m, mentions, &g) 
		case "removeplayers":
			handler.RemovePlayers(s, m, mentions, &g)
		case "start":
			handler.StartGame(s, m, &g)
		case "stop":
			handler.StopGame(s, m, &g)
		case "showplayers":
			handler.ShowPlayers(s, m, &g)
		default:
			handler.InvalidCommand(s, m)
		} 

	}
}
