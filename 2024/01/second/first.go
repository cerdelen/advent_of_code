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

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var firstRow []int
    count_second_row := make(map[int]int)

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
        // secondRow = append(secondRow, second)
        count_second_row[second]++
    }

    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    sort.Ints(firstRow)

    var distance = 0

    // fmt.Println(count_second_row)

    for i := 0; i < len(firstRow); i++ {
        distance += firstRow[i] * count_second_row[firstRow[i]]
        // distance += absInt(firstRow[i] - secondRow[i])
    }

    fmt.Println(distance)
}
