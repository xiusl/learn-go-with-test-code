package main

import (
	"io"
)

type FileSystemStore struct {
	database io.ReadSeeker
}

func (fs *FileSystemStore) GetLeague() []Player {
	_, _ = fs.database.Seek(0, 0)
	league, _ := NewLeague(fs.database)
	return league
}

func (fs *FileSystemStore) GetPlayerScore(name string) int {
	return 0
}