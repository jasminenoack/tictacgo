// ticTacGo
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var board = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
var turn = "X"

func getMove() byte {
	fmt.Println(turn + " where would you like to move?")
	reader := bufio.NewReader(os.Stdin)
	number, err := reader.ReadByte()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	number = number - '0'
	if number < 1 || number > 9 {
		fmt.Println("Invalid input, try again")
		return getMove()
	}
	index := number - 1
	if board[index] == "X" || board[index] == "O" {
		fmt.Println("Invalid input, try again")
		return getMove()
	}
	return index
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

func switchTurn() {
	if turn == "X" {
		turn = "O"
	} else {
		turn = "X"
	}
}

func makeMove(index byte) {
	board[index] = turn
}

func checkRow(row int) string {
	if board[row*3] == board[row*3+1] && board[row*3+2] == board[row*3] {
		return board[row*3]
	}
	return ""
}

func checkColumn(column int) string {
	if board[column] == board[column+3] && board[column+6] == board[column] {
		return board[column]
	}
	return ""
}

func checkDiag(diag int) string {
	if diag == 0 && board[0] == board[4] && board[4] == board[8] {
		return board[0]
	}
	if diag == 1 && board[2] == board[4] && board[4] == board[6] {
		return board[2]
	}
	return ""
}

func getWinner() string {
	winner := ""
	for i := 0; i < 3; i++ {
		rowWinner := checkRow(i)
		columnWinner := checkColumn(i)
		diagWinner := checkDiag(i)
		if len(rowWinner) > 0 {
			winner = rowWinner
		} else if len(columnWinner) > 0 {
			winner = columnWinner
		} else if len(diagWinner) > 0 {
			winner = diagWinner
		}
		if len(winner) > 0 {
			return winner
		}
	}
	return ""
}

func main() {
	printBoard()
	winner := ""
	for i := 0; i < 9; i++ {
		makeMove(getMove())
		switchTurn()
		printBoard()
		winner = getWinner()
		if len(winner) > 0 {
			fmt.Println("The winner of the game is " + winner + "!")
			os.Exit(0)
		}
	}
	fmt.Println("There is no winner!")
}
