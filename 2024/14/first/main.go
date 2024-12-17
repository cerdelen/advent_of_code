package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	// "sort"
	// "io/ioutil"
	// "strings"
	// "strconv"
)

const max_x = 101
const max_y = 103
const x_mid = 50
const y_mid = 51

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var result = 0

    scanner := bufio.NewScanner(file)
    q1 := 0
    q2 := 0
    q3 := 0
    q4 := 0
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

        px = ((100 * vx) + px) % max_x
        py = ((100 * vy) + py) % max_y

        if px < 0 {
            px = px + max_x
        }

        if py < 0 {
            py = py + max_y
        }

        if px < x_mid && py < y_mid {
            q1++
        }
        if px < x_mid && py > y_mid {
            q2++
        }
        if px > x_mid && py > y_mid {
            q3++
        }
        if px > x_mid && py < y_mid {
            q4++
        }
    }

    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


    result = q1 * q2 * q3 * q4
    fmt.Printf("result: %d\n", result)
}
