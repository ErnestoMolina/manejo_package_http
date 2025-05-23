package rps

import (
	"fmt"
	"testing"
)

func TestPlayRound(t *testing.T) {
	for i := 0; i < 3; i++ {
		round := PlayRound(i)
		switch i {
		case 0:
			fmt.Println("El jugador eligió piedra")
		case 1:
			fmt.Println("El jugador eligió papel")
		case 2:
			fmt.Println("El jugador eligió tijera")
		}
		fmt.Printf("Message: %s\n", round.Message)
		fmt.Printf("ComputerChoise: %s\n", round.ComputerChoise)
		fmt.Printf("RoundResult: %s\n", round.RoundResult)
		fmt.Printf("ComputerChoiseInt: %d\n", round.ComputerChoiseInt)
		fmt.Printf("ComputerScore: %s\n", round.ComputerScore)
		fmt.Printf("PlayerScore: %s\n", round.PlayerScore)
		fmt.Println("\n=============================")
	}
}
