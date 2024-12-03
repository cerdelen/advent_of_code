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

func isSafe(intArray []int) bool {
    if intArray[0] < intArray[1] { // we are ascending
        for i := 1; i < len(intArray); i++ {
            if !(intArray[i] > intArray[i - 1]) {
                return false
            }
            if !(intArray[i] - intArray[i - 1] < 4) {
                return false
            }

        }
    } else { // we are descending
        for i := 1; i < len(intArray); i++ {
            if !(intArray[i] < intArray[i - 1]) {
                return false
            }
            if !(intArray[i - 1] - intArray[i] < 4) {
                return false
            }
        }
    }
    return true
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var safe_reports = 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        numbers := strings.Fields(line)
        intArray := make([]int, len(numbers))

        for i, str := range numbers {
            num, err := strconv.Atoi(str)
            if err != nil {
                log.Fatalf("Error converting string to int: %v", err)
            }
            intArray[i] = num
        }

        var row_safe bool = isSafe(intArray)

        if row_safe == false {
            fmt.Println("From: ", intArray)
            for i := 0; i < len(intArray); i++ {
                newArray := make([]int, 0, len(intArray)-1)
                newArray = append(newArray, intArray[:i]...)
                newArray = append(newArray, intArray[i+1:]...)
                fmt.Println("To: ", newArray)

                row_safe = isSafe(newArray)
                if row_safe {
                    break
                }
            }
            fmt.Println("")
        }

        if row_safe {
            // fmt.Println("Safe: ", intArray)
            safe_reports++
        } else {
            // fmt.Println("Unsafe: ", intArray)
        }
    }

    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    fmt.Printf("safeReports: %d\n", safe_reports)
}