package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/clementbowe14/ProfessorTrivia/game"
)

const (
	helpMessage = `Here is the list of commands that you can use.
	All commands here start with the prefix !prof before the word.
	addplayer 	<player_name> <player_2name>  add a player to the game.
	removeplayer 	<player_name> <player_2name> remove a player from the game.
	start 	to start the game.
	stop 	to stop the game.
	current_stats 	to see your player's current stats in the game
	If you have any more questions please check out our page using the command /page`
	addSuccess    = "All players have been successfully added"
	removeSuccess = "All players have been removed from the game"
)

func HelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, helpMessage)

}

func AddPlayers(s *discordgo.Session, m *discordgo.MessageCreate, players []*discordgo.User, g *game.Game) {

	if len(players) == 0 {
		err := g.AddPlayer(m.Author.Username)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}
		s.ChannelMessageSend(m.ChannelID, addSuccess)
		return
	}

	for _, player := range players {
		err := g.AddPlayer(player.Username)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
		}
	}

	s.ChannelMessageSend(m.ChannelID, addSuccess)

}

func RemovePlayers(s *discordgo.Session, m *discordgo.MessageCreate, players []*discordgo.User, g *game.Game) {
	if len(players) == 0 {
		err := g.RemovePlayer(m.Author.Username)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}
	}

	for _, player := range players {
		if err := g.RemovePlayer(player.Username); err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
		}
	}

	s.ChannelMessageSend(m.ChannelID, removeSuccess)

}

func ShowPlayers(s *discordgo.Session, m *discordgo.MessageCreate, g *game.Game) {
	msg, err := g.ShowPlayers()
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}
	s.ChannelMessageSend(m.ChannelID, msg)
}

func StopGame(s *discordgo.Session, m *discordgo.MessageCreate, g *game.Game) {

}

func StartGame(s *discordgo.Session, m *discordgo.MessageCreate, g *game.Game) {

}

func InvalidCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Invalid command please try again."+"You can use !prof help to see the list of commands you can use")
}
