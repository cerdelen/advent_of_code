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

func check_for_loop(grid [][]rune, y int, x int, y_dir int, x_dir int, m map[IntPair][]IntPair) int {
    dir_par := IntPair{y_dir, x_dir}
    if val, exists := m[IntPair{y, x}]; exists {
        for _, dir := range val {
            if dir == dir_par {
                // fmt.Printf("found a possible loop")
                return 1
            }
        }
        m[IntPair{y, x}] = append(val, IntPair{y_dir, x_dir})
    } else {
        m[IntPair{y, x}] = []IntPair{{y_dir, x_dir}}
    }

    // fmt.Printf("move, y: %d, x: %d, y_dir: %d, x_dir: %d\n", y, x, y_dir, x_dir)
    if y+y_dir < 0 || x+x_dir < 0 || y+y_dir > len(grid) - 1|| x+x_dir > len(grid[0]) - 1 {
        return 0
    }
    for grid[y + y_dir][x + x_dir] == '#' {
        change_dir(&y_dir, &x_dir)
        if y+y_dir < 0 || x+x_dir < 0 || y+y_dir > len(grid) - 1|| x+x_dir > len(grid[0]) - 1 {
            return 0
        }
    }
    return check_for_loop(grid, y + y_dir, x + x_dir, y_dir, x_dir, m)
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

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            pairCount := make(map[IntPair][]IntPair)
            if grid[i][j] == '.' {
                grid[i][j] = '#'
                result += check_for_loop(grid, y, x, -1, 0, pairCount)
                grid[i][j] = '.'
            }
            // fmt.Printf("loop")
        }
    }

    // fmt.Printf("result: %d\nother_result: %d\nmap_length: %d\n", result, other_result, len(pairCount))
    fmt.Printf("result: %d\n", result)
}
