package rps

import "strings"

const STATE_NEW = "new"
const STATE_JOINED = "joined"
const STATE_WAITING = "waiting"
const STATE_PLAYING = "playing"


type Player struct {
	Name     string
	State    string
	action   chan string
	Messages chan string
}

func NewPlayer(name string) *Player {
	return &Player{Name: name, action: make(chan string), Messages: make(chan string), State: STATE_NEW}
}

func (p Player) Act(message string) {
	p.action <- strings.TrimSpace(message)
}

func (p Player) Msg(message string) {
	defer func(){
		recover()
	}()
	p.Messages <- message
}
