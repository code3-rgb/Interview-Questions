package main

import "fmt"

var pl = fmt.Println

var path = [][]int{
	{1, 0, 0, 0},
	{1, 1, 0, 1},
	{0, 1, 0, 0},
	{1, 1, 1, 1},
}

var sol = [][]int{
	{0, 0, 0, 0},
	{0, 0, 0, 0},
	{0, 0, 0, 0},
	{0, 0, 0, 0},
}

func isSafe(x int, y int) bool {
	if x >= 0 && y >= 0 && path[x][y] == 1 {
		return true
	}
	return false
}

func solve() {
	if !solve_maze(0, 0) {
		pl("Solution does not exist")
		return
	}
	pl(sol)
	return
}

func solve_maze(x int, y int) bool {
	if x == len(path)-1 && y == len(path)-1 && path[x][y] == 1 {
		sol[x][y] = 1
		return true
	}
	if isSafe(x, y) {
		if sol[x][y] == 1 {
			return false
		}
		sol[x][y] = 1
		if solve_maze(x+1, y) {
			return true
		}
		if solve_maze(x, y+1) {
			return true
		}
		sol[x][y] = 0
		return false
	}
	return false

}

func main() {
	solve()
}
