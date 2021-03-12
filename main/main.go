package main
import(
	"fmt"
	"strings"
	"strconv"
)

type player struct {
	symbol rune
	name   string
}
type board [3][3]*player
type game struct {
	gameboard board
	player1   player
	player2   player
}

func main() {
	games := new(game)
	fmt.Println()
	printboard(games.gameboard)
	games.player1.name = "first"
	games.player2.name = "second"
	games.player1.symbol = 'X'
	games.player2.symbol = 'O'
	p := make([]*player, 2)
	p[0] = &games.player1
	p[1] = &games.player2
	for i := 0; ; i = (i + 1) % 2 {
		fmt.Println(p[i].name, " turn")
		fmt.Println("enter position")
		b := ""
		fmt.Scanln(&b)
		fmt.Println(b)
		pos := strings.Split(b, ",")
		num1, _ := strconv.Atoi(pos[0])
		num2, _ := strconv.Atoi(pos[1])
		ret := games.gameboard.play(num1, num2, p[i])
		if ret == true {
			fmt.Println(p[i].name, " wins")
			printboard(games.gameboard)
			break
		}
		printboard(games.gameboard)
	}

}

func printboard(a board) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] == nil {
				fmt.Print("__ ")
			} else {
				fmt.Print(string(a[i][j].symbol), " ")
			}
		}
		fmt.Println()
		fmt.Println()
		fmt.Println()
	}

}

func (c *board) play(i, j int, turn *player) bool {
	i = i - 1
	j = j - 1
	c[i][j] = turn
	return c.won(i, j)

}

func (c *board) won(i, j int) bool {

	switch {
	case i == 0 && j == 0:
		if c[0][0] == c[0][1] {
			if c[0][0] == c[0][2] {
				return true
			}
		}
		if c[0][0] == c[1][0] {
			if c[0][0] == c[2][0] {
				return true
			}
		}
		if c[0][0] == c[1][1] {
			if c[0][0] == c[2][2] {
				return true
			}
		}

	case i == 0 && j == 1:
		if c[0][1] == c[1][1] {
			if c[0][1] == c[2][1] {
				return true
			}
		}
		if c[0][1] == c[0][0] {
			if c[0][1] == c[0][2] {
				return true
			}
		}
	case i == 0 && j == 2:
		if c[0][2] == c[0][1] {
			if c[0][0] == c[0][2] {
				return true
			}
		}
		if c[0][2] == c[1][2] {
			if c[0][2] == c[2][2] {
				return true
			}
		}
		if c[0][2] == c[1][1] {
			if c[0][2] == c[2][0] {
				return true
			}
		}
	case i == 1 && j == 0:
		if c[0][0] == c[1][0] {
			if c[0][0] == c[2][0] {
				return true
			}
		}
		if c[1][0] == c[1][2] {
			if c[1][0] == c[1][1] {
				return true
			}
		}

	case i == 1 && j == 1:
		if c[1][1] == c[0][1] {
			if c[1][1] == c[2][1] {
				return true
			}
		}
		if c[1][1] == c[1][0] {
			if c[1][1] == c[1][2] {
				return true
			}
		}

	case i == 1 && j == 2:
		if c[1][2] == c[0][2] {
			if c[0][2] == c[2][2] {
				return true
			}
		}
		if c[1][0] == c[1][1] {
			if c[1][0] == c[1][2] {
				return true
			}
		}

	case i == 2 && j == 0:
		if c[2][0] == c[1][0] {
			if c[0][0] == c[2][0] {
				return true
			}
		}
		if c[2][0] == c[2][1] {
			if c[2][0] == c[2][2] {
				return true
			}
		}
		if c[2][0] == c[1][1] {
			if c[2][0] == c[0][2] {
				return true
			}
		}
	case i == 2 && j == 1:
		if c[2][1] == c[2][0] {
			if c[2][0] == c[2][2] {
				return true
			}
		}
		if c[2][1] == c[1][1] {
			if c[0][1] == c[2][1] {
				return true
			}
		}

	case i == 2 && j == 2:
		if c[2][2] == c[2][0] {
			if c[2][0] == c[2][1] {
				return true
			}
		}
		if c[2][2] == c[1][2] {
			if c[0][2] == c[2][2] {
				return true
			}
		}
		if c[1][1] == c[2][2] {
			if c[2][2] == c[0][0] {
				return true
			}
		}
	}
	return false
}
