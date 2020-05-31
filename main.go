package main

import (
	"fmt"
)

var positions [3][3]string

func initialize(){
	for i := 0; i < 3; i++{
		for j := 0; j < 3; j++{
    		positions[i][j] = " "
    	}
  	}
}

func main(){
	fmt.Println("Welcome to TicTacGo")
	fmt.Println(`Board Position 

   0     1      2
0     |     |
 _________________

1     |     |
 _________________

2     |     |

	`)
	fmt.Println("Game Starts")
}