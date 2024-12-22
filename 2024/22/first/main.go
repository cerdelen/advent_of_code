package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	// "sort"
	// "io/ioutil"
	// "strings"
	// "strconv"
)

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

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var result = 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        secret, _ = strconv.Atoi(line)

        for i := 0; i < 2000; i++ {
            next_secret_number()
        }

        result += secret
    }

    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


    fmt.Printf("result: %d\n", result)
}
