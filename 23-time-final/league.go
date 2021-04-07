package poker

import (
	"encoding/json"
	"fmt"
	"io"
)

// 存储 player
// League stores a collection of players.
type League []Player

// 根据名字查找用户
// Find tries to return a player from a League.
func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

// NewLeague creates a League from JSON.
func NewLeague(rdr io.Reader) (League, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)

	if err != nil {
		err = fmt.Errorf("problem parsing League %v", err)
	}
	return league, err
}