package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"errors"
)

// tockens = count_a * 3 + count_b
// count_a * mov_x_a = x - count_b * mov_x_b
// count_a * mov_y_a = y - count_b * mov_y_b
func extrac_Button(line string) (int, int) {
    re := regexp.MustCompile(`X\+(\d+),\s*Y\+(\d+)`)
    matches := re.FindStringSubmatch(line)
    if len(matches) < 3 {
        log.Fatal("Button Line Malformed", line)
	}

	x, errX := strconv.Atoi(matches[1])
	y, errY := strconv.Atoi(matches[2])

    if errX != nil || errY != nil {
        log.Fatal("Button Line Atoi error", line, errX, errY)
    }

    return y, x
}

func extrac_goal(line string) (int, int) {
    re := regexp.MustCompile(`X=(\d+),\s*Y=(\d+)`)
    matches := re.FindStringSubmatch(line)
    if len(matches) < 3 {
        log.Fatal("goal Line Malformed", line)
	}

	x, errX := strconv.Atoi(matches[1])
	y, errY := strconv.Atoi(matches[2])

    if errX != nil || errY != nil {
        log.Fatal("goal Line Atoi error", line, errX, errY)
    }

    return y, x
}

func isInt(f float64) bool {
	return f == math.Trunc(f)
}

func minimizeZ(ax, bx, ay, by, x, y float64) (float64, error) {
    det := ax * by - ay * bx
    if det == 0 {
        return 0, errors.New("the system of equations is singular (det = 0), no unique solution")
    }

    a := (x * by - y * bx) / det
    b := (ax * y - ay * x) / det

    if !isInt(a) || !isInt(b) {
        return 0, errors.New("one of the is not int")
    }
    z := 3 * a + b

    return z, nil
}

func main() {
    file, err := os.Open("input.txt")
    // file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var result = 0

    var text []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        text = append(text, line)
    }

    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    ind := 0

    for {
        if ind >= len(text) {
            break
        }
        ay, ax := extrac_Button(text[ind])
        ind++
        by, bx := extrac_Button(text[ind])
        ind++

        y, x := extrac_goal(text[ind])
        ind += 2

        z, err := minimizeZ(float64(ax), float64(bx), float64(ay), float64(by), float64(x), float64(y))
        if err != nil {
            fmt.Println("Error:", err)
        } else {
            result += int(z)
        }
    }


    fmt.Printf("result: %d\n", result)
}
