package main

import (
	"bufio"
	"fmt"
	"github.com/emou/go-presentation/gorps/rps"
	"net"
	"strings"
)

func writeMsg(writer *bufio.Writer, msg string) error {
	_, err := writer.Write([]byte(msg + "\n"))
	return err
}

func readPlayerName(reader *bufio.Reader, writer *bufio.Writer) (string, error) {
	for {
		err := writeMsg(writer, "LOGIN")
		if err != nil {
			return "", err
		}
		name, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		name = strings.TrimSpace(name)
		if name != "" {
			return name, nil
		}
	}

}

func serve(conn net.Conn, game *rps.Game) error {
	defer conn.Close()

	fmt.Printf("Incoming connection: %+v\n", conn.RemoteAddr())
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	name, err := readPlayerName(reader, writer)
	if err != nil {
		fmt.Printf("Error on login: %s", name)
	}
	player := rps.NewPlayer(name)
	game.AddPlayer(player)

	// TODO: Pull function
	go func() {
		fmt.Println("Opening player message feed ", player.Name)
		defer fmt.Println("Closing player message feed ", player.Name)

		for {
			select {
			case msg := <-player.Messages:
				_, err := writer.Write([]byte(msg))
				if err != nil {
					fmt.Println(err)
					return
				}
				writer.Flush()
			case <-player.Finish:
				return
			}
		}
	}()

	game.StartMatch(player)

	// TODO: Pull function
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
	listener, err := net.Listen("tcp", ":9000")

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
