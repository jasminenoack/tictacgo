// ticTacGo
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var board = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func getMove() byte {
	fmt.Println("Where would you like to move?")
	reader := bufio.NewReader(os.Stdin)
	number, err := reader.ReadByte()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	number = number - '0'
	if number < 0 || number > 8 {
		fmt.Println("Invalid input")
		os.Exit(1)
	}
	return number
}

func printBoard() {
	fmt.Println("")
	fmt.Println(strings.Join(board[:3], " | "))
	fmt.Println("---------")
	fmt.Println(strings.Join(board[3:6], " | "))
	fmt.Println("---------")
	fmt.Println(strings.Join(board[6:], " | "))
	fmt.Println("")
}

func main() {
	printBoard()
	getMove()
	printBoard()
}
