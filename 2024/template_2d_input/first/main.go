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

    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    print_grid(grid)

    fmt.Printf("result: %d\n", result)
}
