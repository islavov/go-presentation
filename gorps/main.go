package main

import (
	"net"
	"fmt"
	"bufio"
	"strings"
)


type Player struct {
	name string
	action chan string
	message chan string
}


func NewPlayer(name string) *Player {
	return &Player{name: name, action: make(chan string), message: make(chan string)}
}

type Action struct {
	player *Player
	action string
	param string
}

type Game struct {
	actions chan Action
	players []*Player
}


func NewGame() *Game{
	var game = Game{players: []*Player{}, actions: make(chan Action)}
	go game.cmdloop()
	return &game
}


func startMatch(game *Game, player1 *Player, player2 *Player) {
	var decision1 string
	var decision2 string

	var msg = fmt.Sprintf("Duel between %s and %s. FIGHT!", player1.name, player2.name)
	player1.message <- msg
	player2.message <- msg

	for {
		select {
		case decision1 = <-player1.action:
			fmt.Printf("Hit player 1: %s", decision1)
		case decision2 = <-player2.action:
			fmt.Printf("Hit player 2: %s", decision2)
		default:
			if (decision1 != "" && decision2 != "") {
				game.endMatch(player1, "win")
				game.endMatch(player2, "lose")
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
			game.players = append(game.players, action.player)
		case "removePlayer":
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
				waiting.message <- "waiting for another player..."
			} else {
				go startMatch(game, action.player, waiting)
				waiting = nil
			}
		case "endMatch":
			action.player.message <- fmt.Sprintf("You %s", action.param)
			//game.startMatch(action.player)
		}

	}
}

func (game *Game) addPlayer(player *Player) {
	game.actions <- Action{player:player, action: "addPlayer", param:""}
}


func (game *Game) removePlayer(player *Player) {
	game.actions <- Action{player:player, action: "removePlayer", param:""}
}


func (game *Game) startMatch(player *Player) {
	game.actions <- Action{player:player, action: "startMatch", param:""}
}


func (game *Game) endMatch(player *Player, param string) {
	game.actions <- Action{player:player, action: "endMatch", param:param}
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

	go func () {
		for {
			writer.Write([]byte(<-player.message))
			writer.Flush()
		}
	}()

	for {
		game.startMatch(player)

		for {
			message, err := reader.ReadString('\n')
			fmt.Println(strings.TrimSpace(message))
			player.action <- message
			if err != nil {
				game.removePlayer(player)
				return err
			}

		}
	}

	return nil
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
