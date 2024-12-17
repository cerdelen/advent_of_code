package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	// "sort"
	// "io/ioutil"
	// "strings"
	// "strconv"
)

type Pos struct {
    y, x int
}

func print_grid(grid [][]rune) {
    for _, row := range grid {
        fmt.Println(string(row))
    }
}

func calc_sum(grid [][]rune) int {
    out := 0
    for y := 0; y < len(grid); y++ {
        for x := 0; x < len(grid[0]); x++ {
            if grid[y][x] == 'O' {
                out += 100 * y + x
            }
        }
    }
    return out
}

func make_move(grid [][]rune, y int, x int, pos *Pos) {
    check_pos := *pos
    boxes := false
    loop:
    for {
        check_pos.y += y
        check_pos.x += x
        switch grid[check_pos.y][check_pos.x] {
            case '.': {
                break loop
            }
            case '#': {
                return
            }
            case 'O': {
                boxes = true
            }
            default : {
                // print_grid(grid)
                fmt.Printf("y: %d, x: %d, char: %c\n", check_pos.y, check_pos.x, grid[check_pos.y][check_pos.x])
                fmt.Printf("line: %s\n", string(grid[check_pos.y]))
                log.Fatal("unknown character found")
            }
        }
    }

    if boxes {
        grid[check_pos.y][check_pos.x] = 'O'
        grid[pos.y + y][pos.x + x] = '.'
    }
    pos.y += y
    pos.x += x
}

func make_moves(grid [][]rune, moves string, pos *Pos) {
    for _, move := range moves {
        switch move {
            case '<': {
                make_move(grid, 0, -1, pos)
            }
            case '>': {
                make_move(grid, 0, 1, pos)
            }
            case '^': {
                make_move(grid, -1, 0, pos)
            }
            case 'v': {
                make_move(grid, 1, 0, pos)
            }
            default: {
                log.Fatal("Not cool move")
            }
        }
    }
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var result = 0
    var map_being_read = true

    var grid [][]rune
    scanner := bufio.NewScanner(file)
    var pos Pos = Pos{-1, -1}
    for scanner.Scan() {
        line := scanner.Text()

        if map_being_read {
            if len(line) > 0 {
                ind := strings.Index(line, "@")
                row := []rune(line)
                grid = append(grid, row)
                if ind != -1 {
                    pos.y = len(grid) - 1
                    pos.x = ind
                    grid[pos.y][pos.x] = '.'
                }
            } else {
                if pos.x == -1 || pos.y == -1 {
                    log.Fatal("no start found")
                }
                print_grid(grid)
                map_being_read = false
            }
        } else {
            make_moves(grid, line, &pos)
        }
    }

    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    result = calc_sum(grid)

    print_grid(grid)
    fmt.Printf("result: %d\n", result)
}
