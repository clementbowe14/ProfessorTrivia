package main

import "github.com/bwmarrin/discordgo"

type Game interface {
}
type Player struct {
	playerId  string
	score     int
	gameStart bool
}

type discordGame struct {
	players   map[string]Player
	session   *discordgo.Session
	channelId string
}

func NewDiscordGame(sess *discordgo.Session) Game {
	players := make(map[string]Player)
	return &discordGame{
		players: players,
		session: sess,
	}
}

func (g *discordGame) addPlayer(playerId string) error {
	g.players[playerId] = Player{
		playerId: playerId,
		score:    0,
	}
	return nil
}

func (g *discordGame) handleMessage(m *discordgo.MessageCreate) {
	
}

func (d *discordGame) run() {

}
