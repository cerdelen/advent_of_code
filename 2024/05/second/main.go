package main

import (
    "os"
    "bufio"
    "fmt"
    "log"
    // "sort"
    // "io/ioutil"
    "strings"
    "strconv"
)

// ich mache eine map der rules
// key == second part of rules
// value, all the first part of rules, that have the same second part
// this way i can linearly loop over an update, search for current key and look which ones are not allowed to have been seen before

// edit ... first and second swapped

func rules(line string, m map[int][]int) {
    parts := strings.Split(line, "|")

    if len(parts) != 2 {
        log.Fatal("Not 2 numbers in rule")
    }

    first, err1 := strconv.Atoi(parts[0])
    second, err2 := strconv.Atoi(parts[1])

    if err1 != nil || err2 != nil {
        log.Fatal("Error parsing Atoi")
    }

    if _, exists := m[first]; !exists {
        m[first] = []int{}
    }
    m[first] = append(m[first], second)

    // fmt.Printf("rule: %d, %d\n", first, second)
}

func move_indeces(slice []int, from int, to int) {
    element := slice[from]

    slice = append(slice[:from], slice[from+1:]...)

    slice= append(slice[:to], append([]int{element}, slice[to:]...)...)
}

func rec_updates(update []int, rules map[int][]int) int {
    exists := make(map[int]int)
    for i, num := range update {
        if forbidden_arr, found := rules[num]; found {
            for _, forbidden := range forbidden_arr {
                if index_of_wrong, found := exists[forbidden]; found {
                    move_indeces(update, index_of_wrong, i)
                    return rec_updates(update, rules)
                }
            }
        }
        exists[num] = i
    }
    return update[len(update)/2]
}

func updates(line string, rules map[int][]int) int {
    exists := make(map[int]struct{})
    var update []int = []int{}
    parts := strings.Split(line, ",")
    var error = false

    for _, num := range parts {
        number, err := strconv.Atoi(num)
        if err != nil {
            log.Fatal("Error parsing Atoi in update")
        }
        if !error {
            if forbidden_arr, found := rules[number]; found {
                for _, forbidden := range forbidden_arr {
                    if _, found := exists[forbidden]; found {
                        error = true
                    }
                }
            }
            exists[number] = struct{}{}
        }
        update = append(update, number)
    }
    if error {
        return rec_updates(update, rules)
    }
    return 0
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var result = 0
    parse_rules := false

    rule_map := make(map[int][]int)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        if len(line) == 0 {
            parse_rules = true
            continue
        }

        if parse_rules == true {
            result += updates(line, rule_map)
        } else {
            rules(line, rule_map)
        }

        // fmt.Printf("line: %s\n", line)
    }

    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    fmt.Printf("result: %d\n", result)
}