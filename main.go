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

func player1Function() {
	var x, y int16
	var v string
	fmt.Print("Enter X or anything else to give up: ")
	fmt.Scanln(&v)
	if strings.Trim(v, "\n ") != "X" || strings.Trim(v, "\n ") != "x" {
		//TODO: need to implement give up function and reference it here
	}
	fmt.Print("Enter the x position ")
	fmt.Scanf("%v", &x)
	fmt.Print("Enter the y position: ")
	fmt.Scanf("%v", &y)
	positions[x][y] = v
}

func player2Function() {
	var x, y int16
	var v string
	fmt.Print("Enter O or anything else to give up: ")
	fmt.Scanln(&v)
	if strings.Trim(v, "\n ") != "O" || strings.Trim(v, "\n ") != "o" {
		//TODO: need to implement give up function and reference it here
	}
	fmt.Print("Enter the x position ")
	fmt.Scanf("%v", &x)
	fmt.Print("Enter the y position: ")
	fmt.Scanf("%v", &y)
	positions[x][y] = v
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
		player1()
		player2()
	}
}
