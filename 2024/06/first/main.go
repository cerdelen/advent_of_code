package main

import (
    "os"
    "bufio"
    "fmt"
    "log"
    // "sort"
    // "io/ioutil"
    "strings"
    // "strconv"
)

func change_dir(y_dir *int, x_dir *int) {
    temp := -(*y_dir)
    *y_dir = *x_dir
    *x_dir = temp
}

type IntPair struct {
	First, Second int
}

func move(grid [][]rune, y int, x int, y_dir int, x_dir int, m map[IntPair]struct{}) {
    m[IntPair{y, x}] = struct{}{}

    grid[y][x] = 'x'
    fmt.Printf("move, y: %d, x: %d, y_dir: %d, x_dir: %d\n", y, x, y_dir, x_dir)
    if y+y_dir < 0 || x+x_dir < 0 || y+y_dir > len(grid) - 1|| x+x_dir > len(grid[0]) - 1 {
        return
    }
    for grid[y + y_dir][x + x_dir] == '#' {
        // fmt.Printf("before change dir, y_dir: %d, x_dir: %d\n", y_dir, x_dir)
        change_dir(&y_dir, &x_dir)
        // fmt.Printf("after change dir, y_dir: %d, x_dir: %d\n", y_dir, x_dir)
        if y+y_dir < 0 || x+x_dir < 0 || y+y_dir > len(grid) - 1|| x+x_dir > len(grid[0]) - 1 {
            return
        }
    }
    move(grid, y + y_dir, x + x_dir, y_dir, x_dir, m)
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var result = 0

    var x,y int

    var grid [][] rune
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        row := []rune(line)

        ind := strings.Index(line, "^")
        if ind != -1 {
            y = len(grid) - 1
            x = ind
        }

        grid = append(grid, row)
    }

    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    pairCount := make(map[IntPair]struct{})
    move(grid, y, x, -1, 0, pairCount)

    other_result := 0

    for _, row := range grid {
        for _, r := range row {
            if r == 'x' {
                other_result++
            }
            fmt.Printf("%c", r)
        }
        fmt.Println()
    }

    fmt.Printf("result: %d\nother_result: %d\nmap_length: %d\n", result, other_result, len(pairCount))
}
