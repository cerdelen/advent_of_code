package main

import (
    "sort"
    "os"
    "bufio"
    "fmt"
    // "io/ioutil"
    "log"
    "strings"
    "strconv"
)

func absInt(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    file, err := os.Open("input.txt")
    // content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var firstRow []int
    var secondRow []int

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        numbers := strings.Fields(line)
        if len(numbers) != 2 {
            log.Fatalf("Invalid line format: %s", line)
        }


        first, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatalf("Invalid number in line: %s", line)
		}
		second, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatalf("Invalid number in line: %s", line)
		}

        firstRow = append(firstRow, first)
        secondRow = append(secondRow, second)
    }

    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    sort.Ints(firstRow)
	sort.Ints(secondRow)

    var distance = 0

    for i := 0; i < len(firstRow); i++ {
        distance += absInt(firstRow[i] - secondRow[i])

    }

    fmt.Println(distance)




    // fmt.Println(string(content))
}
