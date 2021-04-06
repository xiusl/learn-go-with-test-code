package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const PlayerPrompt = "Please enter the number of players: "

type CLI struct {
	in  *bufio.Scanner
	out io.Writer
	g *Game
}

func NewCLI(store PlayerStore, in io.Reader, out io.Writer, alerter BlindAlerter) *CLI {
	return &CLI{bufio.NewScanner(in), out, &Game{
		alerter,
		store,
	}}
}

func (cli *CLI) PlayPoker() {
	// 提示用户输入玩家数量
	_, _  = fmt.Fprint(cli.out, PlayerPrompt)

	// 读取用户输入，并转换为 int 类型，忽略错误
	numberOfPlayersInput := cli.readLine()
	numberOfPlayers, _ := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))

	//
	cli.g.Start(numberOfPlayers)

	winnerInput := cli.readLine()
	winner := extractWinner(winnerInput)

	cli.g.Finish(winner)
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}