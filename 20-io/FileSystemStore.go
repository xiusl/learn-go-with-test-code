package main

import (
	"encoding/json"
	"io"
)

type FileSystemStore struct {
	database io.ReadWriteSeeker
}

func (fs *FileSystemStore) GetLeague() []Player {
	_, _ = fs.database.Seek(0, 0)
	league, _ := NewLeague(fs.database)
	return league
}

func (fs *FileSystemStore) GetPlayerScore(name string) int {
	var score int
	for _, player := range fs.GetLeague() {
		if name == player.Name {
			score = player.Score
		}
	}
	return score
}

func (fs *FileSystemStore) RecordWin(name string) {
	league := fs.GetLeague()
	for i, player := range league {
		if name == player.Name {
			league[i].Score++     // not player.Score++，当使用 range 遍历时，使用的是切片的副本
		}
	}
	_, _ = fs.database.Seek(0,0)
	_ = json.NewEncoder(fs.database).Encode(league)
}