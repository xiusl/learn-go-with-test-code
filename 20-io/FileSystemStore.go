package main

import (
	"io"
)

type FileSystemStore struct {
	database io.Reader
}

func (fs *FileSystemStore) GetLeague() []Player {
	league, _ := NewLeague(fs.database)
	return league
}