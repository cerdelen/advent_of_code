package main

import (
    // "sort"
    "os"
    "bufio"
    "fmt"
    // "io/ioutil"
    "log"
    "strings"
    "strconv"
)

func isMul(substr string) int {
    comma_ind := strings.Index(substr, ",")
    if comma_ind == -1 || comma_ind > 3 { return 0 }
    bracket_ind := strings.Index(substr[comma_ind + 1:], ")")
    if bracket_ind == -1 || bracket_ind > 3 { return 0 }
    if comma_ind < 1 || bracket_ind < 1 {
        log.Fatalf("numbers to short: %s\n", substr)
    }
    first, err := strconv.Atoi(substr[:comma_ind])
    if err != nil {
        log.Fatalf("Invalid first number in line: %s", substr)
    }
    second, err := strconv.Atoi(substr[comma_ind+1:comma_ind+1+bracket_ind])
    if err != nil {
        log.Fatalf("Invalid second number in line: %s", substr[comma_ind+1:comma_ind+1+bracket_ind])
    }
    return first * second
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var result = 0

    scanner := bufio.NewScanner(file)
    do := true
    for scanner.Scan() {
        line := scanner.Text()

        for {
            if do == false {
                do_ind := strings.Index(line, "do()")
                if do_ind == -1 {break}
                line = line[do_ind+1:]
                do = true
            }
            ind := strings.Index(line, "mul(")
            dont_ind := strings.Index(line, "don't()")
            if dont_ind < ind && dont_ind != -1 {
                line = line[dont_ind+1:]
                do_ind := strings.Index(line, "do()")
                if do_ind == -1 {
                    do = false
                    break
                }
                line = line[do_ind+1:]
                continue
            }
            if ind == -1 { break }
            if ind+12 >= len(line) {
                result += isMul(line[ind+4:])
            } else {
                result += isMul(line[ind+4:ind+12])
            }
            line = line[ind+4:]
        }
    }

    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    fmt.Printf("safeReports: %d\n", result)
}
