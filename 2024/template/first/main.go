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

        fmt.Printf("line: %s\n", line)
    }

    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


    fmt.Printf("result: %d\n", result)
}
