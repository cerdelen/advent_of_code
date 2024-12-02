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

        var isSafe bool = true
        if intArray[0] < intArray[1] { // we are ascending
            for i := 1; i < len(intArray); i++ {
                if !(intArray[i] > intArray[i - 1]) {
                    isSafe = false
                    fmt.Println("ascending Unsafe because of switch ", intArray)
                    break
                }
                if !(intArray[i] - intArray[i - 1] < 4) {
                    isSafe = false
                    fmt.Println("ascending Unsafe because of >4 ", intArray)
                    break
                }

            }
        } else { // we are descending
            for i := 1; i < len(intArray); i++ {
                if !(intArray[i] < intArray[i - 1]) {

                    fmt.Println("descending Unsafe because of switch ", intArray)
                    isSafe = false
                    break
                }
                if !(intArray[i - 1] - intArray[i] < 4) {
                    fmt.Println("descending Unsafe because of > 4", intArray)
                    isSafe = false
                    break
                }
            }
        }

        if isSafe {
            fmt.Println("Safe: ", intArray)
            safe_reports++
        } else {
            fmt.Println("Unsafe: ", intArray)
        }
    }

    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    fmt.Printf("safeReports: %d\n", safe_reports)
}
