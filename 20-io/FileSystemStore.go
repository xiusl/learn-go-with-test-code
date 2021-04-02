package main

import (
	"encoding/json"
	"io"
)

type FileSystemStore struct {
	database io.ReadWriteSeeker
}

func (fs *FileSystemStore) GetLeague() League {
	_, _ = fs.database.Seek(0, 0)
	league, _ := NewLeague(fs.database)
	return league
}

func (fs *FileSystemStore) GetPlayerScore(name string) int {
	player := fs.GetLeague().Find(name)

	if player != nil {
		return player.Score
	}
	return 0
}

func (fs *FileSystemStore) RecordWin(name string) {
	league := fs.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Score++
	}

	_, _ = fs.database.Seek(0,0)
	_ = json.NewEncoder(fs.database).Encode(league)
}