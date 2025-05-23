package rps

import (
	"math/rand/v2"
	"strconv"
)

const (
	ROCK     = 0 // Piedra. Vence a tijeras. (tijeras + 1) % 3 = 0
	PAPER    = 1 // Papel. Vence a piedra. (piedra + 1) % 3 = 1
	SCISSORS = 2 // Tijeras. Vence a papel. (papel + 1) % 3 = 2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoise    string `json:"computer_choise"`
	RoundResult       string `json:"round_result"`
	ComputerChoiseInt int    `json:"computer_choise_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var winMessages = []string{
	"¡Bien echo!",
	"¡Buen trabajo!",
	"Deberías comprar un boleto de lotería.",
}

var loseMessages = []string{
	"¡Qué lástima!",
	"¡Inténtalo de nuevo!",
	"Hoy simplemente no es tu dia.",
}

var drawMessages = []string{
	"Las grandes mentes piensan igual.",
	"Oh oh. Intenta de nuevo.",
	"Nadie gana, pero puedes intentarlo de nuevo.",
}

var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.IntN(3)
	var computerChoise, roundResult string
	var computerChoiseInt int
	switch computerValue {
	case ROCK:
		computerChoiseInt = ROCK
		computerChoise = "La computadora eligió piedra"
	case PAPER:
		computerChoiseInt = PAPER
		computerChoise = "La computadora eligió papel"
	case SCISSORS:
		computerChoiseInt = SCISSORS
		computerChoise = "La computadora eligió tijera"
	}

	// Generar un numero aleatorio entre 0-2, qe usamos para elegir el mensaje
	messageInt := rand.IntN(3)
	// generar una variable para contener el mensaje
	var message string
	if playerValue == computerValue {
		roundResult = "Es un empate"
		// selecciona un mensaje de drawMessage
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "!El jugador gana¡"
		// seleccionar un mensage de winMessage
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "!La computadora gana¡"
		// seleccionar un mensaje de LoseMessage
		message = loseMessages[messageInt]
	}

	return Round{
		Message:           message,
		ComputerChoise:    computerChoise,
		RoundResult:       roundResult,
		ComputerChoiseInt: computerChoiseInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}
