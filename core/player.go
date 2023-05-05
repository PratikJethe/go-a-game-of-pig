package core

import (
	"math/rand"
	"time"
)

type player struct {
	holdLimit  int
	totalScore int
}

// playround function is responsible to keep rolling dice and decide outcome based on dice value and players strategy
// it returns player's score for that particular round
func (p *player) playRound(winingScore int) int {
	currentScore := 0
	for {
		diceValue := p.rollDice()
		if diceValue == 1 {
			return 0
		}

		currentScore += diceValue
		if currentScore >= p.holdLimit || currentScore >= winingScore {
			return currentScore
		}
	}

}

// resetScore sets player's totalScore to 0
func (p *player) resteScore() {
	p.totalScore = 0
}

// rollDice is responsible to generate random value between 1 to 6
func (p *player) rollDice() int {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(6)

	return randomInt + 1
}

func newPlayer(holdLimit int) player {

	return player{
		holdLimit: holdLimit,
	}
}
