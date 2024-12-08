package main

import (
    "os"
    "bufio"
    "fmt"
    "log"
    // "math"
    // "sort"
    // "io/ioutil"
    // "strings"
    // "strconv"
)

type IntPair struct {
    Y, X int
}

func parseMap(grid [][]rune) int {
    var m map[rune][]IntPair = make(map[rune][]IntPair)
    var antinodes map[IntPair]bool = make(map[IntPair]bool)
    y_max := len(grid)
    x_max := len(grid[0])
    counter := 0

    for y, line := range grid {
        for x, c := range line {
            if c == '.' {
                continue
            }
            if val, exists := m[c]; exists {
                for _, prev := range val {
                    x_dist := x - prev.X
                    y_dist := y - prev.Y
                    k := 0
                    for {
                        if prev.Y - k*y_dist >= 0 && prev.X - k*x_dist >= 0 && prev.X - k*x_dist < x_max {
                            if _, exists := antinodes[IntPair{prev.Y - k*y_dist, prev.X - k*x_dist}]; exists {
                            } else {
                                antinodes[IntPair{prev.Y - k*y_dist, prev.X - k*x_dist}] = true
                                counter++
                            }
                            k++
                        } else {
                            break
                        }
                    }
                    k = 0
                    for {
                        if y + k*y_dist < y_max && x + k*x_dist >= 0 && x + k*x_dist < x_max {
                            if _, exists := antinodes[IntPair{y + k*y_dist, x+ k*x_dist}]; exists {
                            } else {
                                antinodes[IntPair{y + k*y_dist, x+ k*x_dist}] = true
                                counter++
                            }
                            k++
                        } else {
                            break
                        }
                    }
                }
                m[c] = append(val, IntPair{y, x})
            } else {
                m[c] = []IntPair{{y, x}}
            }
        }
    }

    return len(antinodes)
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
    result += parseMap(grid)

    for _, line := range grid {
        for _, c := range line {
            fmt.Print(string(c))
        }
        fmt.Println()
    }

    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    fmt.Printf("result: %d\n", result)
}
