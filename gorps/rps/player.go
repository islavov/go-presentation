package rps

import (
	"strings"
	"fmt"
)

const STATE_NEW = "new"
const STATE_JOINED = "joined"
const STATE_WAITING = "waiting"
const STATE_PLAYING = "playing"

type Player struct {
	Name     string
	State    string
	action   chan string
	messages chan string
}

func NewPlayer(name string) *Player {
	return &Player{Name: name, action: make(chan string), messages: make(chan string), State: STATE_NEW}
}

// Act dispatches a user action
func (p Player) Act(message string) {
	p.action <- strings.TrimSpace(message)
}

// WriteMsg sends a message to the user
func (p *Player) WriteMsg(message string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("failed to write message: ",message, r)
		}
	}()
	p.messages <- message
}

// ReadMsg reads a message, that was send to the user.
// Blocks until a message is received - should be used in a separate goroutine
func (p *Player) ReadMsg() string {
	return <-p.messages
}


func (p *Player) Leave() {
	close(p.messages)
	close(p.action)
}
