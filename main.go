package main

import (
	"fmt"
)

var positions [3][3]string

func initialize() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			positions[i][j] = " "
		}
	}
}

func needToContinue() (string, bool) {
	for _, v := range positions {
		if v == [3]string{"X", "X", "X"} || v == [3]string{"O", "O", "O"} {
			return v[0], false
		}
	}
	for i := 0; i < 3; i++ {
		if positions[0][i] == positions[1][i] && positions[0][i] == positions[2][i] && positions[0][i] != " " {
			return positions[0][i], false
		}
	}
	switch {
	case positions[0][0] == positions[1][1] && positions[0][0] == positions[2][2] && positions[0][0] != " ":
		return positions[0][0], false
	case positions[0][2] == positions[1][1] && positions[0][2] == positions[2][0] && positions[0][2] != " ":
		return positions[0][2], false
	default:
		return " ", true
	}
}

type player struct {
	name   string
	avatar string
}

func (p player) playerFunction() {
	var x, y int16
	fmt.Printf("You are %s\n", string(p.avatar))
	fmt.Print("Enter the x position ")
	fmt.Scanf("%v", &x)
	fmt.Print("Enter the y position: ")
	fmt.Scanf("%v", &y)
	positions[x][y] = p.avatar
}

func printPosition() {
	fmt.Printf(`\nBoard Position

   0     1    2

0  %v  |  %v  | %v
 _________________

1  %v  |  %v  | %v
 _________________

2  %v  |  %v  | %v

`, positions[0][0], positions[0][1], positions[0][2], positions[1][0], positions[1][1], positions[1][2], positions[2][0], positions[2][1], positions[2][2])
}

func main() {
	for {
		initialize()
		var player1, player2 player
		var temp string
		fmt.Println("Welcome to TicTacGo")
		fmt.Println(`Board Position 

   0     1      2

0     |     |
 _________________

1     |     |
 _________________

2     |     |

	`)
		fmt.Print("Enter the name of player 1: ")
		fmt.Scanln(&temp)
		player1 = player{temp, "X"}
		fmt.Print("Enter the name of player 2: ")
		fmt.Scanln(&temp)
		player2 = player{temp, "O"}
		fmt.Println("Game Starts")
		for {
			player1.playerFunction()
			printPosition()
			if k, v := needToContinue(); v == false {
				switch k {
				case "X":
					fmt.Printf("Congrats %s on winning\n", player1.name)
				case "O":
					fmt.Printf("Congrats %s on winning\n", player2.name)
				}
				break
			}
			player2.playerFunction()
			printPosition()
			if k, v := needToContinue(); v == false {
				switch k {
				case "X":
					fmt.Printf("Congrats %s on winning\n", player1)
				case "O":
					fmt.Printf("Congrats %s on winning\n", player2)
				}
				break
			}
		}
		fmt.Print("Wanna play again(y/n): ")
		fmt.Scanf("%s", &temp)
		if temp == "n" || temp == "N" {
			break
		}
	}
}
