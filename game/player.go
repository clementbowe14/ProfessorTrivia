package game

import (
	"errors"
)
const (
	MAX_POINTS_ADDED = 1000000

)
var (
	ErrorInvalidPointsAdded = "An invalid amount of points was added to the total."
)
type Player struct {
	playerId string
	score int
	correctAnswers int
	incorrectAnswers int
	answeredQuestion bool
}


func NewPlayer(playerId string) (Player) {
	return Player {
		playerId: playerId,	
		score: 0,
	}
}
func(p *Player) AddScore(v int) error {
	
	if v < 0  || v > MAX_POINTS_ADDED {
		return errors.New(ErrorInvalidPointsAdded)
	}
	
	p.score += v
	return nil
}

func (p *Player) CorrectAnswer() {
	p.correctAnswers++
}

func (p *Player) IncorrectAnswer() {
	p.incorrectAnswers++
}

