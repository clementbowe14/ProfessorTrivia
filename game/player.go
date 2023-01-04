package game


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
func(p *Player) AddScore(v int) {
	p.score += v
}

func (p *Player) CorrectAnswer() {
	p.correctAnswers++
}

func (p *Player) IncorrectAnswer() {
	p.incorrectAnswers++
}

