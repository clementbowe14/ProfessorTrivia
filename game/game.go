package game

import (
	"errors"
	"fmt"
	"strings"
	"github.com/google/uuid"
)

const (
	MAX_TRIVIA_PLAYERS = 8
	ErrorMaxPlayersExceeded = "Max players exceeded for this game"
	ErrorPlayerAlreadyExists = "This player has already been added to the game"
	ErrorRemovePlayerOperationFailed = "Not enough players added in the game to remove players"
	ErrorInsufficientPlayers = "At least one player must be added to start the game"
	ErrorNoPlayersToShow = "There are no players to show."
	ErrorPlayerNotInGame = "is not in the the game."
)

type Game struct {
	players map[string]*Player
	serverName string
	gameId string
	winner string
	questionList []Question
}

func New() (Game) {
	g := Game{
		gameId: uuid.New().String(),
	}
	g.players = make(map[string]*Player)

	return g
}

func (g *Game) SetServerName(serverName string) {
	g.serverName = serverName
}

func (g *Game) ComputeWinner() {
	max := 0
	playerName := ""
	for name, player := range(g.players) {
		if player.score > max {
			max = player.score
			playerName = name
		}
	}
	g.winner = playerName
}

func (g *Game) RemovePlayer(playerId string) (error) {

	if len(g.players) < 1 {
		return errors.New(ErrorRemovePlayerOperationFailed)
	}

	player, ok := g.players[playerId]
	if !ok {
		return errors.New(fmt.Sprintf("%s%s", playerId, ErrorPlayerNotInGame))
	}
	
	delete(g.players, player.playerId)
	
	return nil
}

func(g *Game) AddPlayer(playerId string) (error) {
	if len(g.players) > MAX_TRIVIA_PLAYERS {
		return errors.New(ErrorMaxPlayersExceeded)
	}
	
	player, ok := g.players[playerId]

	if ok == true {
		return errors.New(fmt.Sprintf("%s %s", player.playerId, ErrorPlayerAlreadyExists))
	}

	g.players[playerId] = &Player{
		playerId: playerId,
		score: 0,
	}

	return nil
}

func (g Game) ShowPlayers() (string, error) {
	if len(g.players) == 0 {
		return "", errors.New(ErrorNoPlayersToShow)
	}
	players := make([]string, len(g.players))
	idx := 0
	for k := range(g.players) {
		players[idx] = k
		idx++
		
	}
	return strings.Join(players, " "), nil
}


