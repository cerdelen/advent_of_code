package main

import (
	"bufio"
	// "container/heap"
	"fmt"
	"log"
	// "math"
	"os"
	"strings"
	// "regexp"
	// "strconv"
)

var grid [][]rune
var distance_grid [][]int
func print_grid(grid [][]rune) {
    for _, row := range grid {
        fmt.Println(string(row))
    }
}
func print_distance_grid() {
    for _, row := range distance_grid {
        fmt.Println(row)
    }
}

const y_max = 141
const x_max = 141
var start Pos
var end Pos

type Pos struct {
    y, x int
}

func init_distance_grid(y, x int) {
    distance_grid = make([][]int, y)
    for i := range grid {
        distance_grid[i] = make([]int, x)
        for j := range grid[i] {
            distance_grid[i][j] = -1
        }
    }
}

func fill_distance_grid(curr Pos, cost int) {
    if curr == end {
        distance_grid[curr.y][curr.x] = cost
        return
    }
    if curr.x < 0 || curr.x >= x_max || curr.y < 0 || curr.y >= y_max {
        return
    }
    if grid[curr.y][curr.x] == '#' {
        return
    }
    if distance_grid[curr.y][curr.x] != -1 {
        return
    }
    distance_grid[curr.y][curr.x] = cost
    fill_distance_grid(Pos{curr.y + 1, curr.x}, cost + 1)
    fill_distance_grid(Pos{curr.y - 1, curr.x}, cost + 1)
    fill_distance_grid(Pos{curr.y, curr.x + 1}, cost + 1)
    fill_distance_grid(Pos{curr.y, curr.x - 1}, cost + 1)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func check_cheat_validity(cheat_start, cheat_end Pos) int {
    if cheat_end.x < 0 || cheat_end.x >= x_max || cheat_end.y < 0 || cheat_end.y >= y_max {
        return 0
    }
    if distance_grid[cheat_end.y][cheat_end.x] == -1 ||  distance_grid[cheat_start.y][cheat_start.x] == -1{
        return 0
    }
    dist := abs(distance_grid[cheat_end.y][cheat_end.x] - distance_grid[cheat_start.y][cheat_start.x])

    if dist >= 102 {
        return 1
    }
    return 0
}

func find_cheats() int {
    res := 0
    for y, row := range distance_grid {
        for x := range row {
            res += check_cheat_validity(Pos{y, x}, Pos{y + 2, x})
            res += check_cheat_validity(Pos{y, x}, Pos{y + 1, x + 1})
            res += check_cheat_validity(Pos{y, x}, Pos{y - 1, x + 1})
            res += check_cheat_validity(Pos{y, x}, Pos{y, x + 2})
        }
    }
    return res
}


func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var result = 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        ind := strings.Index(line, "S")
        if ind != -1 {
            start = Pos{len(grid), ind}
        }
        ind = strings.Index(line, "E")
        if ind != -1 {
            end = Pos{len(grid), ind}
        }
        row := []rune(line)
        grid = append(grid, row)
    }

    if len(grid) != y_max && len(grid[0]) != x_max {
        log.Fatal("Wrong input Matrix")
    }

    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    init_distance_grid(len(grid), len(grid[0]))
    fill_distance_grid(start, 0)
    // print_distance_grid()

    result = find_cheats()

    fmt.Printf("result: %d\n", result)
}

