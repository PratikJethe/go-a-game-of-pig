package core

import "fmt"

const WIN_SCORE int = 100
const NO_OF_GAMES int = 10

type game struct {
	p1 *player
	p2 *player
}

// start function starts game for a particular strategy
// it returns an map representing number of wins by both player
func (g *game) start() map[string]int {

	gameResult := map[string]int{}
	for i := 0; i < NO_OF_GAMES; i++ {
		playingPlayer := "p1"
		currentPlayer := g.p1
		for {

			roundScore := currentPlayer.playRound(WIN_SCORE)
			currentPlayer.totalScore += roundScore

			if currentPlayer.totalScore > WIN_SCORE {
				gameResult[playingPlayer]++
				g.resetPlayers()
				break
			}

			//switch players
			if playingPlayer == "p1" {
				playingPlayer = "p2"
				currentPlayer = g.p2
			} else {
				playingPlayer = "p1"
				currentPlayer = g.p1
			}

		}

	}

	return gameResult

}

// resetPlayers resets both players score to 0 at end of each round
func (g *game) resetPlayers() {
	g.p1.resteScore()
	g.p2.resteScore()
}

func newGame(p1 *player, p2 *player) game {
	return game{
		p1: p1,
		p2: p2,
	}
}

//getWinLossPercentage prints formatted message for the results of the game
func getWinLossPercentage(p1HoldLimt int, p2HoldLimit int, gameResult map[string]int) {
	p1WinPercentage := 100 * float64(gameResult["p1"]) / float64(NO_OF_GAMES)
	fmt.Printf(
		"Holding at %v vs Holding at  %v: wins: %v/%v (%.2f%%), losses: %v/%v (%.2f%%)\n",
		p1HoldLimt,
		p2HoldLimit,
		gameResult["p1"],
		NO_OF_GAMES,
		p1WinPercentage,
		NO_OF_GAMES-gameResult["p1"],
		NO_OF_GAMES,
		100-p1WinPercentage,
	)

}

//getWinLossPercentageForMultiMultiStrat prints formatted message for the results of the game when Multi Multi strategy is used
func getWinLossPercentageForMultiMultiStrat(p1HoldLimt int, p1WinNo int, p1GamesPlayed int) {

	p1WinPercentage := 100 * float64(p1WinNo) / float64(p1GamesPlayed)
	fmt.Printf(
		"Result: Wins, losses staying at k = %v: %v/%v (%.2f%%), %v/%v (%.2f%%)\n",
		p1HoldLimt,
		p1WinNo,
		p1GamesPlayed,
		p1WinPercentage,
		p1GamesPlayed-p1WinNo,
		p1GamesPlayed,
		100-p1WinPercentage,
	)

}

// StartSimulation is resposible to simulate game between 2 players using provided stategy
func StartSimulation(p1UpperHoldLimit int, p1LowerHoldLimit int, p2UpperHoldLimit int, p2LowerHoldLimit int) {
	isMultiMultiStrategy := false

	// to check if both player are changing srategy
	// in case of multi multi srategy output will be generated for each P1's srategy against all P2's strategy
	if p1LowerHoldLimit != p1UpperHoldLimit && p2LowerHoldLimit != p2UpperHoldLimit {
		isMultiMultiStrategy = true
	}

	for p1HoldLimit := p1LowerHoldLimit; p1HoldLimit <= p1UpperHoldLimit; p1HoldLimit++ {
		p1OverAllWins := 0
		p1GamesPlayedWithOneSrategy := 0
		for p2HoldLimit := p2LowerHoldLimit; p2HoldLimit <= p2UpperHoldLimit; p2HoldLimit++ {

			if p1HoldLimit == p2HoldLimit {
				continue
			}
			p1 := newPlayer(p1HoldLimit)
			p2 := newPlayer(p2HoldLimit)
			game := newGame(&p1, &p2)

			gameResult := game.start()
			if !isMultiMultiStrategy {
				getWinLossPercentage(p1HoldLimit, p2HoldLimit, gameResult)
			} else {
				p1OverAllWins += gameResult["p1"]
				p1GamesPlayedWithOneSrategy += NO_OF_GAMES
			}

		}

		if isMultiMultiStrategy {
			getWinLossPercentageForMultiMultiStrat(p1HoldLimit, p1OverAllWins, p1GamesPlayedWithOneSrategy)
		}
	}

}
