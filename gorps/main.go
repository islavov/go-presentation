package main

import (
	"bufio"
	"fmt"
	"math"
	"net"
	"strings"
)

var SURRENDER = -1

var ROCK = 0
var PAPER = 1
var SCISSORS = 2

const MSG_WEAPON = "Choose your weapon (1: rock, 2:paper, 3:scissors): "

const STATE_NEW = "new"
const STATE_JOINED = "joined"
const STATE_WAITING = "waiting"
const STATE_PLAYING = "playing"

type Player struct {
	name    string
	state   string
	action  chan string
	message chan string
}

func NewPlayer(name string) *Player {
	return &Player{name: name, action: make(chan string), message: make(chan string), state: STATE_NEW}
}

type Action struct {
	player *Player
	action string
	param  string
}

type Game struct {
	actions chan Action
	players []*Player
}

func NewGame() *Game {
	var game = Game{players: []*Player{}, actions: make(chan Action)}
	go game.cmdloop()
	return &game
}

func msgPlayers(msg string, players ...*Player) {
	for _, player := range players {
		player.message <- msg
	}
}

func handleAction(player *Player, action string) *int {
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
		player.message <- MSG_WEAPON
		return nil
	}
}

func gameLogic(decision1 int, decision2 int) bool {
	if decision1 == decision2 {
		return false
	}
	if math.Mod(float64(decision1+3)-float64(decision2), 3) == 1 {
		return true
	}
	return false
}

func startMatch(game *Game, player1 *Player, player2 *Player) {
	var decision1 *int
	var decision2 *int

	msgPlayers(MSG_WEAPON, player1, player2)

	for {
		select {
		case action := <-player1.action:
			decision1 = handleAction(player1, action)
		case action := <-player2.action:
			decision2 = handleAction(player2, action)
		default:
			var winner *Player
			var loser *Player

			if decision1 != nil && decision2 != nil {
				if *decision1 == SURRENDER {
					winner = player2
				} else if *decision2 == SURRENDER {
					winner = player1
				} else if gameLogic(*decision1, *decision2) {
					winner = player1
					loser = player2
				} else if gameLogic(*decision2, *decision1) {
					winner = player2
					loser = player1
				} else {
					msgPlayers("DRAW", player1, player2)
					decision1 = nil
					decision2 = nil
				}

			}

			if winner != nil {
				game.endMatch(winner, "win")
			}
			if loser != nil {
				game.endMatch(loser, "lose")
			}
			if winner != nil || loser != nil {
				return
			}
		}
	}
}

func (game *Game) cmdloop() {
	var waiting *Player

	for {
		var action = <-game.actions
		fmt.Printf("Game message: %+v\n", action)

		switch action.action {
		case "addPlayer":
			fmt.Printf("Joining: %s\n", action.player.name)
			action.player.state = STATE_JOINED
			game.players = append(game.players, action.player)
		case "removePlayer":
			if action.player.state == STATE_PLAYING {
				action.player.action <- fmt.Sprintf("%d", SURRENDER)
			}
			if action.player == waiting {
				waiting = nil
			}
			for idx, player := range game.players {
				if player == action.player {
					fmt.Printf("Leaving: %s\n", action.player.name)
					game.players = append(game.players[:idx], game.players[idx+1:]...)
					break
				}
			}
		case "startMatch":

			if waiting == nil {
				waiting = action.player
				waiting.message <- "waiting for another player...\n"
				waiting.state = STATE_WAITING
			} else {

				var msg = fmt.Sprintf("Duel between %s and %s. FIGHT!\n",
					waiting.name, action.player.name)
				waiting.message <- msg
				action.player.message <- msg

				waiting.state = STATE_PLAYING
				action.player.state = STATE_PLAYING

				go startMatch(game, action.player, waiting)
				waiting = nil
			}
		case "endMatch":
			action.player.message <- fmt.Sprintf("You %s. Press any key to continue.\n", action.param)
			action.player.state = STATE_JOINED
		}

	}
}

func (game *Game) addPlayer(player *Player) {
	game.actions <- Action{player: player, action: "addPlayer", param: ""}
}

func (game *Game) removePlayer(player *Player) {
	game.actions <- Action{player: player, action: "removePlayer", param: ""}
}

func (game *Game) startMatch(player *Player) {
	game.actions <- Action{player: player, action: "startMatch", param: ""}
}

func (game *Game) endMatch(player *Player, param string) {
	game.actions <- Action{player: player, action: "endMatch", param: param}
}

func serve(conn net.Conn, game *Game) error {
	defer conn.Close()

	fmt.Printf("%+v\n", conn.RemoteAddr())
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	writer.Write([]byte("Enter player name: "))
	writer.Flush()

	name, _ := reader.ReadString('\n')

	player := NewPlayer(strings.TrimSpace(name))
	game.addPlayer(player)

	go func() {
		for {
			writer.Write([]byte(<-player.message))
			writer.Flush()
		}
	}()

	game.startMatch(player)

	for {
		if player.state == STATE_JOINED {
			game.startMatch(player)
		}
		message, err := reader.ReadString('\n')
		if player.state == STATE_PLAYING && message != "" {
			player.action <- strings.TrimSpace(message)
		}

		if err != nil {
			game.removePlayer(player)
			close(player.message)
			return err
		}

	}
}

func main() {
	game := NewGame()
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err == nil {
			go serve(conn, game)
		} else {
			fmt.Printf("%+v", err)
		}
	}

}
