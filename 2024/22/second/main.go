package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type fourChanges struct {
    data [4]int
    size int
}

func (r *fourChanges) Push(value int) {
	if r.size < 4 {
		r.data[r.size] = value
        r.size++
	} else {
		r.data[0] = r.data[1]
		r.data[1] = r.data[2]
		r.data[2] = r.data[3]
		r.data[3] = value
	}
}

var secret int

func mix(number int) {
    secret =  number ^ secret
}

func prune() {
    secret %= 16777216
}

func next_secret_number() {
// Calculate the result of multiplying the secret number by 64.
//     Then, mix this result into the secret number. Finally, prune the secret number.
    mix(secret * 64)
    prune()
// Calculate the result of dividing the secret number by 32.
//     Round the result down to the nearest integer. Then, mix this result into the secret number.
//      Finally, prune the secret number.
    mix(int(secret / 32))
    prune()
// Calculate the result of multiplying the secret number by 2048.
//     Then, mix this result into the secret number. Finally, prune the secret number.
    mix(secret * 2048)
    prune()
}

var m map[[4]int]int
var localm map[[4]int]int

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var result = 0

    scanner := bufio.NewScanner(file)
    m = make(map[[4]int]int)
    for scanner.Scan() {
        var changes fourChanges
        changes.size = 0
        prevVal := -10000


        line := scanner.Text()

        secret, _ = strconv.Atoi(line)
        localm = make(map[[4]int]int)

        for i := 0; i < 2000; i++ {
            next_secret_number()
            value := secret % 10

            if prevVal != -10000 {
                change := prevVal - value
                changes.Push(change)
                if changes.size == 4 {
                    if _, exists := localm[changes.data]; exists {
                    } else {
                        localm[changes.data] = value
                    }
                }
            }

            prevVal = value
        }

        for key, value := range localm {
            m[key] += value
        }
    }

    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    for _, value := range m {
        result = max(result, value)
	}

    fmt.Printf("result: %d\n", result)
}
