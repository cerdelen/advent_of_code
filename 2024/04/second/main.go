package main

import (
    // "sort"
    "os"
    "bufio"
    "fmt"
    // "io/ioutil"
    "log"
    // "strings"
    // "strconv"
)

func diagonal_1(x int, y int, grid [][]rune) bool {
    if grid[y-1][x-1] == 'M' {
        if grid[y+1][x+1] == 'S' {
            return true
        }
    }
    if grid[y-1][x-1] == 'S' {
        if grid[y+1][x+1] == 'M' {
            return true
        }
    }
    return false
}

func diagonal_2(x int, y int, grid [][]rune) bool {
    if grid[y-1][x+1] == 'M' {
        if grid[y+1][x-1] == 'S' {
            return true
        }
    }
    if grid[y-1][x+1] == 'S' {
        if grid[y+1][x-1] == 'M' {
            return true
        }
    }
    return false
}

func recurs_start(x int, y int, grid [][]rune) bool {
    if x < 1 || y < 1 {
        return false
    }

    if y >= len(grid) -1 {
        return false
    }
    if x >= len(grid[y]) - 1 {
        return false
    }

    if diagonal_1(x, y, grid) && diagonal_2(x, y, grid) {
        return true
    }
    return false
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var result = 0

    var grid [][] rune
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        row := []rune(line)

        grid = append(grid, row)
    }

    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


    for y, row := range grid {
        for x, r := range row {
             if r == 'A' {
                 if recurs_start(x, y, grid) {
                     result++
                 }
                 // result = result + 1
             }
         }
        fmt.Println(row)
	}

    fmt.Printf("safeReports: %d\n", result)
}
