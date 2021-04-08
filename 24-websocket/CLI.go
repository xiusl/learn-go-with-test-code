package poker

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in 			*bufio.Scanner
	out 		io.Writer
	game 		Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in: bufio.NewScanner(in),
		out: out,
		game: game,
	}
}

const (
	PlayerPrompt = "Please enter the number of players: "
	BadPlayerInputErrMsg = "Bad value received"
	BadWinnerInputErrMsg = "invalid winner input, expect format of 'PlayerName wins'"
)

func (cli *CLI) PlayPoker() {
	_, _ = fmt.Fprint(cli.out, PlayerPrompt)

	numberOfPlayers, err := strconv.Atoi(cli.readLine())

	if err != nil {
		_, _ = fmt.Fprintf(cli.out, BadPlayerInputErrMsg)
		return
	}

	cli.game.Start(numberOfPlayers)

	winnerInput := cli.readLine()
	winner, err := cli.extractWinner(winnerInput)
	if err != nil {
		_, _ = fmt.Fprintf(cli.out, BadWinnerInputErrMsg)
		return
	}

	cli.game.Finish(winner)
}

func (cli *CLI) extractWinner(userInput string) (string, error) {
	if !strings.Contains(userInput, " wins") {
		return "", errors.New(BadWinnerInputErrMsg)
	}
	return strings.Replace(userInput, " wins", "", 1), nil
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}