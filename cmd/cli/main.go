package main

import (
	"fmt"

	"github.com/Gornak40/sicily-origins/internal/cliengine"
	"github.com/Gornak40/sicily-origins/internal/gamecore"
)

func main() {
	eng := cliengine.New()
	eng.Reset(1)

	for {
		switch eng.GetStatus() {
		case gamecore.StatusProgress:
		case gamecore.StatusRedWinner:
			fmt.Println("RED PLAYER IS THE WINNER") //nolint:forbidigo // Temp solution.

			return
		case gamecore.StatusBlackWinner:
			fmt.Println("BLACK PLAYER IS THE WINNER") //nolint:forbidigo // Temp solution.

			return
		case gamecore.StatusDraw:
			fmt.Println("BATTLE ENDED UP IN A DRAW") //nolint:forbidigo // Temp solution.

			return
		}
	}
}
