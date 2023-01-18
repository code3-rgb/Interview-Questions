package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

var pl = fmt.Println

var board = [20][20]int{}

func events() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if scanner.Text() == "1" {
			move()
		}
		drop()
	}

	if scanner.Err() != nil {
		pl("Error occured")
	}
}

func display() {
	for _, lane := range board {
		pl("\t\t", lane)
	}
}

func clear() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

var down int = 0
var col int = 10

func move() {
	col++
}

func drop() {
	for i := 1; i < 20; i++ {
		board[i][col] = 0
		board[i][col] = 1
		time.Sleep(1 * time.Second)
		clear()

		display()
	}
}

func main() {
	// drop()
	events()
}
