package main

import (
	"net"
	"fmt"
	"bufio"
)

type Game struct {

}




func serve(conn net.Conn) error {
	defer conn.Close()

	fmt.Printf("%+v\n", conn.RemoteAddr())
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			return err
		}

		writer.Write([]byte(message))
		writer.Flush()
	}


	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err == nil {
			go serve(conn)
		} else {
			fmt.Printf("%+v", err)
		}
	}

}
