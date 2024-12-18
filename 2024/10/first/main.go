package main

import (
    "os"
    "bufio"
    "fmt"
    "log"
    // "sort"
    // "io/ioutil"
    // "strings"
    // "strconv"
)

func print_grid(grid [][]rune) {
    for _, row := range grid {
        fmt.Println(string(row))
    }
}

type Pos struct {
    y, x int
}

var visited map[Pos]struct{}

func calc_tailhead(grid [][]rune, y, x, val int) int {
    if y >= len(grid) || x >= len(grid[0]) || y < 0 || x < 0 {
        return 0
    }
    if grid[y][x] != rune(val+'0') {
        return 0
    }
    key := Pos{y, x}
    if _, exists := visited[key]; exists {
        return 0
    } else {
        visited[key] = struct{}{}
    }
    if val == 9 {
        return 1
    }
    out := calc_tailhead(grid, y + 1, x + 0, val + 1) +
            calc_tailhead(grid, y - 1, x + 0, val + 1) +
            calc_tailhead(grid, y + 0, x + 1, val + 1) +
            calc_tailhead(grid, y + 0, x - 1, val + 1)
    return out
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var result = 0

    var grid [][]rune
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        row := []rune(line)
        grid = append(grid, row)
    }

    for y, row := range grid {
        for x, c := range row {
            if c == '0' {
                visited = make(map[Pos]struct{})
                result += calc_tailhead(grid, y, x, 0)
            }
        }
    }

    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    print_grid(grid)

    fmt.Printf("result: %d\n", result)
}
