package rps

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Score struct {
	Wins    int32
	Matches int32
}

type ScoreBoard struct {
	Scores    map[string]*Score
	scorefile string
}

func NewScoreBoard() *ScoreBoard {
	scoreboard := ScoreBoard{scorefile: "scoreboard.txt", Scores: map[string]*Score{}}
	scoreboard.readScores()
	go scoreboard.scoreWriter()
	return &scoreboard
}

func (s *ScoreBoard) readScores() {
	data, err := ioutil.ReadFile(s.scorefile)
	if err != nil {
		return
	}

	json.Unmarshal(data, s.Scores)
}

func (s *ScoreBoard) scoreWriter() {
	tmpfile := fmt.Sprintf("%s.tmp", s.scorefile)

	for {
		data, err := json.Marshal(s.Scores)
		if err != nil {
			fmt.Println("error saving score: ", err)
		}

		err = ioutil.WriteFile(tmpfile, data, 0666)
		if err != nil {
			fmt.Println("error saving score: ", err)
		}

		os.Rename(tmpfile, s.scorefile)
		time.Sleep(3 * time.Second)
	}

}

func (s *ScoreBoard) Add(user string, win bool) {
	_, found := s.Scores[user]
	if !found {
		s.Scores[user] = &Score{Wins: 0, Matches: 0}
	}
	s.Scores[user].Matches += 1
	if win {
		s.Scores[user].Wins += 1
	}
}
