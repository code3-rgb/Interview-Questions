package main

import "fmt"

var pl = fmt.Println

var N int = 4

var sol = [][]int{
	{0, 0, 0, 0},
	{0, 0, 0, 0},
	{0, 0, 0, 0},
	{0, 0, 0, 0},
}
var maze = [][]int{
	{1, 0, 0, 0},
	{1, 1, 0, 1},
	{0, 1, 0, 0},
	{1, 1, 1, 1},
}

func isSafe(maze [][]int, x int, y int) bool {
	if x >= 0 && x < N && y >= 0 && y < N && maze[x][y] == 1 {
		return true
	}
	return false
}

func solve(maze [][]int, sol [][]int, x int, y int) bool {
	if x == N-1 && y == N-1 && maze[x][y] == 1 {
		sol[x][y] = 1
		return true
	}

	if isSafe(maze, x, y) {
		if sol[x][y] == 1 {
			return false
		}
		sol[x][y] = 1

		if solve(maze, sol, x+1, y) {
			return true
		}
		if solve(maze, sol, x, y+1) {
			return true
		}
		sol[x][y] = 1
		return false
	}

	return false
}

func main() {
	if solve(maze, sol, 0, 0) {
		pl(sol)
	} else {
		pl("Solution does not exist!")
	}
}
