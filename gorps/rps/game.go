package rps

import "fmt"

const addPlayer, removePlayer, startMatch, endMatch = "addPlayer", "removePlayer", "startMatch", "removeMatch"

// Action struct holds a player action
type Action struct {
	player *Player
	action string
	param  string
}

// Game struct holds the action channel and players
type Game struct {
	actions chan Action
	players []*Player
}


// NewGame starts a new game
func NewGame() *Game {
	var game = Game{players: []*Player{}, actions: make(chan Action)}
	go game.cmdloop()
	return &game
}

// cmdloop starts a command processing loop, reading from the game command channel
func (game *Game) cmdloop() {
	var waiting *Player

	for {
		var action = <-game.actions
		fmt.Printf("Game message: %+v\n", action)

		switch action.action {
		case addPlayer:
			fmt.Printf("Joining: %s\n", action.player.Name)
			action.player.State = STATE_JOINED
			game.players = append(game.players, action.player)
		case removePlayer:
			if action.player.State == STATE_PLAYING {
				action.player.Act(fmt.Sprintf("%d", SURRENDER))
			}
			if action.player == waiting {
				waiting = nil
			}
			for idx, player := range game.players {
				if player == action.player {
					fmt.Printf("Leaving: %s\n", action.player.Name)
					game.players = append(game.players[:idx], game.players[idx + 1:]...)
					break
				}
			}
		case startMatch:

			if waiting == nil {
				waiting = action.player
				waiting.Messages <- "waiting for another player...\n"
				waiting.State = STATE_WAITING
			} else {

				var msg = fmt.Sprintf("Duel between %s and %s. FIGHT!\n",
					waiting.Name, action.player.Name)
				waiting.Messages <- msg
				action.player.Messages <- msg

				waiting.State = STATE_PLAYING
				action.player.State = STATE_PLAYING

				match := NewMatch(game, action.player, waiting)
				go match.start()
				waiting = nil
			}
		case endMatch:
			action.player.Messages <- fmt.Sprintf("You %s. Press any key to continue.\n", action.param)
			action.player.State = STATE_JOINED
		}
	}
}


// AddPlayer adds a player to the game
func (game Game) AddPlayer(player *Player) {
	game.actions <- Action{player: player, action: addPlayer, param: ""}
}

// RemovePlayer removes a player from the game
func (game Game) RemovePlayer(player *Player) {
	game.actions <- Action{player: player, action: removePlayer, param: ""}
}

// StartMatch starts a new match
func (game Game) StartMatch(player *Player) {
	game.actions <- Action{player: player, action: startMatch, param: ""}
}

// EndMatch ends a match
func (game Game) EndMatch(player *Player, param string) {
	game.actions <- Action{player: player, action: endMatch, param: param}
}
