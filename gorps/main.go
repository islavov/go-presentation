package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	rps "./rps"
)

func serve(conn net.Conn, game *rps.Game) error {
	defer conn.Close()

	fmt.Printf("%+v\n", conn.RemoteAddr())
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	writer.Write([]byte("Enter player name: "))
	writer.Flush()

	name, _ := reader.ReadString('\n')

	player := rps.NewPlayer(strings.TrimSpace(name))
	game.AddPlayer(player)

	go func() {
		for {
			writer.Write([]byte(<-player.Messages))
			writer.Flush()
		}
	}()

	game.StartMatch(player)

	for {
		if player.State == rps.STATE_JOINED {
			game.StartMatch(player)
		}
		message, err := reader.ReadString('\n')
		if player.State == rps.STATE_PLAYING && message != "" {
			player.Act(message)
		}

		if err != nil {
			game.RemovePlayer(player)
			close(player.Messages)
			return err
		}
	}
}

func main() {
	game := rps.NewGame()
	listener, err := net.Listen("tcp", ":8080")

	fmt.Println("Listening on", listener.Addr())
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
