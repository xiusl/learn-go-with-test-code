package main

import (
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

}