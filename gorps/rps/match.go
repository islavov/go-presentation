package rps

import (
	"fmt"
	"math"
)

var SURRENDER = -1
var ROCK = 0
var PAPER = 1
var SCISSORS = 2

const MSG_WEAPON = "Choose your weapon (1: rock, 2:paper, 3:scissors): "

var MSG_MAP = map[int]string{
	-1: "surrender",
	0: "rock",
	1: "paper",
	2: "scissors",
}


// msgPlayers messages all players with a message
func msgPlayers(msg string, players ...*Player) {
	for _, player := range players {
		player.Msg(msg)
	}
}

// Match struct holds the match players and decisions
type Match struct {
	game      *Game
	player1   *Player
	player2   *Player
	decision1 *int
	decision2 *int
}

// NewMatch creates a new match
func NewMatch(game *Game, player1 *Player, player2 *Player) *Match {
	return &Match{game: game, player1: player1, player2: player2}
}

// handleUserAction handles a player
func (m Match) handlePlayerAction(player *Player, action string) *int {
	switch action {
	case "1":
		return &ROCK
	case "2":
		return &PAPER
	case "3":
		return &SCISSORS
	case "-1":
		return &SURRENDER
	default:
		return nil
	}
}

func gameLogic(decision1 int, decision2 int) bool {
	if decision1 == decision2 {
		return false
	}
	if math.Mod(float64(decision1 + 3) - float64(decision2), 3) == 1 {
		return true
	}
	return false
}

func (m *Match) checkWinner() (*Player, *Player) {
	switch {
	case *m.decision1 == SURRENDER:
		return m.player2, nil
	case *m.decision2 == SURRENDER:
		return m.player1, nil
	case gameLogic(*m.decision1, *m.decision2):
		return m.player1, m.player2
	case gameLogic(*m.decision2, *m.decision1):
		return m.player2, m.player1
	default:
		return nil, nil
	}
}


// start starts a new match
func (m *Match) start() {
	msgPlayers(MSG_WEAPON, m.player1, m.player2)

	for {
		select {
		case action := <-m.player1.action:
			m.decision1 = m.handlePlayerAction(m.player1, action)
			if m.decision1 == nil {
				msgPlayers(MSG_WEAPON, m.player1)
			}

		case action := <-m.player2.action:
			m.decision2 = m.handlePlayerAction(m.player2, action);
			if m.decision2 == nil {
				msgPlayers(MSG_WEAPON, m.player2)
			}
		default:

			if m.decision1 != nil && m.decision2 != nil {
				choices := fmt.Sprintf(
					"%s chose %s and %s chose %s...\n",
					m.player1.Name, MSG_MAP[*m.decision1],
					m.player2.Name, MSG_MAP[*m.decision2],
				)
				msgPlayers(choices, m.player1, m.player2)

				winner, loser := m.checkWinner()
				if winner == nil && loser == nil {
					msgPlayers("Its a TIE...\n", m.player1, m.player2)
					msgPlayers(MSG_WEAPON, m.player1, m.player2)
				}

				if winner != nil {
					m.game.EndMatch(winner, "win")
				}
				if loser != nil {
					m.game.EndMatch(loser, "lose")
				}
				if winner != nil || loser != nil {
					return
				}

				m.decision1 = nil
				m.decision2 = nil
			}
		}
	}
}

