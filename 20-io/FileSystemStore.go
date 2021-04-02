package main

import (
	"encoding/json"
	"io"
)

type FileSystemStore struct {
	database io.Reader
}

func (fs *FileSystemStore) GetLeague() []Player {
	var league []Player
	_ = json.NewDecoder(fs.database).Decode(&league)
	return league
}