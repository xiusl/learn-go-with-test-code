package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Tape struct {
	file *os.File
}

func (t *Tape) Write(p []byte) (n int, err error) {
	_  = t.file.Truncate(0)
	_, _ = t.file.Seek(0, 0)
	return t.file.Write(p)
}

type FileSystemStore struct {
	database *json.Encoder
	league League
}

func NewFileSystemStore(file *os.File) (*FileSystemStore, error) {
	err := initialisePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemStore{
		database: json.NewEncoder(&Tape{file}),
		league: league,
	}, nil
}

func (fs *FileSystemStore) GetLeague() League {
	sort.Slice(fs.league, func(i, j int) bool {
		return fs.league[i].Score > fs.league[j].Score
	})
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

	_ = fs.database.Encode(fs.league)
}

func initialisePlayerDBFile(file *os.File) error {
	_, _ = file.Seek(0, 0)

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	if  info.Size() == 0 {
		_, _ = file.Write([]byte("[]"))
		_, _ = file.Seek(0, 0)
	}
	return nil
}
/*NOTE
每当有人调用 GetLeague() 或 GetPlayerScore() 时，我们就从头读取该文件，并将其解析为 JSON。
	我们不应该这样做，因为 FileSystemStore 完全负责 league 的状态。
	我们只是希望在开始时使用该文件来获取当前状态，并在数据更改时更新它。
*/