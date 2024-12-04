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

func recursive_find_word(x int, y int, grid [][]rune, substr string, x_dif int, y_dif int) bool {
    if len(substr) == 0 {
        return true
    }
    if x < 0 || y < 0 {
        return false
    }
    if y >= len(grid) {
        return false
    }
    if x >= len(grid[y]) {
        return false
    }

    if (grid[y][x] == rune(substr[0])) {
        return recursive_find_word(x + x_dif, y + y_dif, grid, substr[1:], x_dif, y_dif)
    }

    return false
}

func recurs_start(x int, y int, grid [][]rune) int {
    res := 0
    str := "XMAS"

    if (recursive_find_word(x, y, grid, str, 1, 0) == true) {
        res +=1
    }
    if (recursive_find_word(x, y, grid, str, 0, 1) == true) {
        res +=1
    }
    if (recursive_find_word(x, y, grid, str, 1, 1) == true) {
        res +=1
    }
    if (recursive_find_word(x, y, grid, str, -1, 0) == true) {
        res +=1
    }
    if (recursive_find_word(x, y, grid, str, 0, -1) == true) {
        res +=1
    }
    if (recursive_find_word(x, y, grid, str, -1, -1) == true) {
        res +=1
    }
    if (recursive_find_word(x, y, grid, str, 1, -1) == true) {
        res +=1
    }
    if (recursive_find_word(x, y, grid, str, -1, 1) == true) {
        res +=1
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
             if r == 'X' {
                 result += recurs_start(x, y, grid)
             }
         }
        fmt.Println(row)
	}

    fmt.Printf("safeReports: %d\n", result)
}
