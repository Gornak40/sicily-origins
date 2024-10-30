package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/Gornak40/sicily-origins/internal/gamecore"
)

//nolint:forbidigo,gosec // Temp solution.
func main() {
	eng := gamecore.NewEngine(1)
	go eng.Run()

	for {
		select {
		case s := <-eng.State():
			listInfo(&s)
			rh := s.Red.Deck.GetHand(s.Turn)
			ridx := rand.Intn(len(rh))
			bh := s.Black.Deck.GetHand(s.Turn)
			bidx := rand.Intn(len(bh))
			fmt.Printf("RED PLAYS: %s\n", rh[ridx].Name())
			fmt.Printf("BLACK PLAYS: %s\n", bh[bidx].Name())
			fmt.Println()
			eng.Round() <- gamecore.GameRound{
				RedCardIdx:   ridx,
				BlackCardIdx: bidx,
			}
		case s := <-eng.Status():
			endGame(s)

			return
		}
	}
}

func getHandString(turn int, deck gamecore.DeckHand) string {
	hand := deck.GetHand(turn)
	shand := make([]string, 0, len(hand))
	for _, h := range hand {
		shand = append(shand, fmt.Sprintf("%s (%d)", h.Name(), h))
	}

	return strings.Join(shand, ", ")
}

//nolint:forbidigo // Temp solution.
func listInfo(s *gamecore.GameState) {
	fmt.Printf("TURN: %d\n", s.Turn)
	fmt.Printf("RED SCORE: %d\n", s.Red.Score)
	fmt.Printf("RED HAND: %s\n", getHandString(s.Turn, s.Red.Deck))
	fmt.Printf("BLACK SCORE: %d\n", s.Black.Score)
	fmt.Printf("BLACK HAND: %s\n", getHandString(s.Turn, s.Black.Deck))
	fmt.Println("===")
}

//nolint:forbidigo // Temp solution.
func endGame(s gamecore.GameStatus) {
	switch s {
	case gamecore.StatusRedWinner:
		fmt.Println("RED PLAYER IS THE WINNER")
	case gamecore.StatusBlackWinner:
		fmt.Println("BLACK PLAYER IS THE WINNER")
	case gamecore.StatusDraw:
		fmt.Println("BATTLE ENDED UP IN A DRAW")
	}
}
