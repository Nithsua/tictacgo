package main

import (
	"fmt"
	"strings"
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
		// fmt.Println("Loop 1")
		if v == [3]string{"X", "X", "X"} || v == [3]string{"O", "O", "O"} {
			return v[0], false
		}
	}
	for i := 0; i < 3; i++ {
		// fmt.Println("Loop 2")
		if positions[0][i] == positions[1][i] && positions[0][i] == positions[2][i] && positions[0][i] != " " {
			return positions[0][i], false
		}
	}
	switch {
	case positions[0][0] == positions[1][1] && positions[0][0] == positions[2][2] && positions[0][0] != " ":
		// fmt.Println("Case 1")
		return positions[0][0], false
	case positions[0][2] == positions[1][1] && positions[0][2] == positions[2][0] && positions[0][2] != " ":
		// fmt.Println("Case 2")
		return positions[0][2], false
	default:
		return "", true
	}
}

func player1Function() {
	var x, y int16
	var v string
	fmt.Print("Enter X or anything else to give up: ")
	fmt.Scanln(&v)
	fmt.Println(v)
	if strings.Trim(v, "\n ") != "X" || strings.Trim(v, "\n ") != "x" {
		fmt.Println("Exiting the Program") //temporary
	}
	fmt.Print("Enter the x position ")
	fmt.Scanf("%v", &x)
	fmt.Print("Enter the y position: ")
	fmt.Scanf("%v", &y)
	fmt.Println(x, y)
	positions[x][y] = "X"
}

func player2Function() {
	var x, y int16
	var v string
	fmt.Print("Enter O or anything else to give up: ")
	fmt.Scanln(&v)
	if strings.Trim(v, "\n ") != "O" || strings.Trim(v, "\n ") != "o" {
		fmt.Println("Exiting the Program") //temporary
	}
	fmt.Print("Enter the x position ")
	fmt.Scanf("%v", &x)
	fmt.Print("Enter the y position: ")
	fmt.Scanf("%v", &y)
	positions[x][y] = "O"
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
	initialize()
	var player1, player2 string
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
	fmt.Scanln(&player1)
	fmt.Print("Enter the name of player 2: ")
	fmt.Scanln(&player2)
	fmt.Println("Game Starts")
	for {
		player1Function()
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
		player2Function()
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
}
