package main

import (
	"encoding/json"
	"io"
)

type Tape struct {
	file io.ReadWriteSeeker
}

func (t *Tape) Write(p []byte) (n int, err error) {
	_, _ = t.file.Seek(0, 0)
	return t.file.Write(p)
}

type FileSystemStore struct {
	database io.Writer
	league League
}

func NewFileSystemStore(database io.ReadWriteSeeker) *FileSystemStore {
	_, _ = database.Seek(0, 0)
	league, _ := NewLeague(database)
	return &FileSystemStore{database: &Tape{database}, league: league}
}

func (fs *FileSystemStore) GetLeague() League {
	return fs.league
}

func (fs *FileSystemStore) GetPlayerScore(name string) int {
	player := fs.league.Find(name)

	if player != nil {
		return player.Score
	}
	return 0
}

func (fs *FileSystemStore) RecordWin(name string) {
	player := fs.league.Find(name)

	if player != nil {
		player.Score++
	} else {
		fs.league = append(fs.league, Player{name, 1})
	}

	_ = json.NewEncoder(fs.database).Encode(fs.league)
}

/*NOTE
每当有人调用 GetLeague() 或 GetPlayerScore() 时，我们就从头读取该文件，并将其解析为 JSON。
	我们不应该这样做，因为 FileSystemStore 完全负责 league 的状态。
	我们只是希望在开始时使用该文件来获取当前状态，并在数据更改时更新它。
*/