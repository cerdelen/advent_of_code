package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const max_x = 101
const max_y = 103
const x_mid = 50
const y_mid = 51
type IntPair struct {
    Y, X int
}


func main() {
    var m map[IntPair]struct{} = make(map[IntPair]struct{})
    success := true
    for rounds := 0 ; rounds < 1000000; rounds++{
        success = true
        file, err := os.Open("input.txt")
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()
        scanner := bufio.NewScanner(file)
        grid := make([][]byte, 103)
        for i := range grid {
            grid[i] = make([]byte, 101)
            for j := range grid[i] {
                grid[i][j] = ' '
            }
        }
        for scanner.Scan() {
            line := scanner.Text()

            pattern := `p=(-?\d+),(-?\d+)\s+v=(-?\d+),(-?\d+)`

            re := regexp.MustCompile(pattern)

            matches := re.FindStringSubmatch(line)

            if len(matches) != 5 {
                log.Fatal("Missformed line", line)
            }

            px, _ := strconv.Atoi(matches[1])
            py, _ := strconv.Atoi(matches[2])
            vx, _ := strconv.Atoi(matches[3])
            vy, _ := strconv.Atoi(matches[4])

            px = ((rounds * vx) + px) % max_x
            py = ((rounds * vy) + py) % max_y

            if px < 0 {
                px = px + max_x
            }

            if py < 0 {
                py = py + max_y
            }

            if _, exists := m[IntPair{py, px}]; exists {
                success = false
                break
            } else {
                m[IntPair{py, px}] = struct{}{}
            }
            grid[py][px] = 'x'
        }

        if err := scanner.Err(); err != nil {
            log.Fatal(err)
        }
        if success {
            fmt.Println("round count: ", rounds)
            for _, row := range grid {
                fmt.Println(string(row))
            }
        }
        m = make(map[IntPair]struct{})
        file.Close()
    }

}
