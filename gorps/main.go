package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"./rps"
)

func readPlayerName(reader bufio.Reader, writer bufio.Writer) string {
	for {
		writer.Write([]byte("Enter player name: "))
		writer.Flush()
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		if name != "" {
			return name
		}
	}

}

func serve(conn net.Conn, game *rps.Game) error {
	defer conn.Close()

	fmt.Printf("%+v\n", conn.RemoteAddr())
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	name := readPlayerName(reader, writer)
	player := rps.NewPlayer(name)
	game.AddPlayer(player)

	go func() {
		fmt.Println("Opening player message feed ", player.Name)
		defer fmt.Println("Closing player message feed ", player.Name)

		for {
			select {
			case msg := <- player.Messages:
				_, err := writer.Write([]byte(msg))
				if err != nil {
					fmt.Println(err)
					return
				}
				writer.Flush()
			case <- player.Finish:
				return
			}
		}
	}()

	game.StartMatch(player)

	for {
		message, err := reader.ReadString('\n')
		if player.State == rps.STATE_PLAYING && message != "" {
			player.Act(message)
		}

		if err != nil {
			game.RemovePlayer(player)
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
