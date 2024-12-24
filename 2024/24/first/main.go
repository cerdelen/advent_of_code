package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var gates map[string]bool

func original_gates(line string) {
    regex := `^(\w+):\s*(\d+)$`
	re := regexp.MustCompile(regex)

	match := re.FindStringSubmatch(line)
	if len(match) > 0 {
		name := match[1]
		value := match[2]
        if value == "1" {
            gates[name] = true
        } else if value == "0" {
            gates[name] = false
        } else {
            log.Fatal("Wrong formation for original gate ", line)
        }
	} else {
        log.Fatal("Wrong formation for original gate ", line)
	}
}

func set_gate(val1, val2 bool, new_gate, operator string) {
    switch operator {
        case "AND": {
            gates[new_gate] = val1 && val2
        }
        case "OR": {
            gates[new_gate] = val1 || val2
        }
        case "XOR": {
            gates[new_gate] = (val1 || val2) && !(val1 && val2)
        }
        default: {
            log.Fatal("Bad operator ", operator)
        }
    }
}

type Combination struct {
    first, second, new, operator string
}

var queue []Combination

func set_comb_gate(comb Combination){
    val1, ok1 := gates[comb.first]
    val2, ok2 := gates[comb.second]
    if ok1 && ok2 {
        set_gate(val1, val2, comb.new, comb.operator)
    } else {
        queue = append(queue, comb)
    }
}

func combination_gates(line string) {
	regex := `^(\w+)\s+(AND|OR|XOR)\s+(\w+)\s+->\s+(\w+)$`
	re := regexp.MustCompile(regex)

	match := re.FindStringSubmatch(line)
	if len(match) > 0 {
		first_gate := match[1]
		operator := match[2]
		second_gate := match[3]
		new_gate := match[4]
        comb := Combination{first: first_gate, second: second_gate, new: new_gate, operator: operator}
        set_comb_gate(comb)
	} else {
		fmt.Println("No match found")
	}
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var result = int64(0)

    gates = make(map[string]bool)
    queue = make([]Combination, 0)

    var first_part = true
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if len(line) == 0 {
            first_part = false
            continue
        }
        if first_part {
            original_gates(line)
        } else {
            combination_gates(line)
        }

    }

    for len(queue) != 0 {
        comb := queue[0]
        queue = queue[1:]
        set_comb_gate(comb)
    }

    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    highest := 50


	for i := highest; i >= 0; i-- {
		key := fmt.Sprintf("z%02d", i)
		if value, exists := gates[key]; exists {
            result = result << 1
            if value {
                result += 1
            }
		}
	}

    fmt.Printf("result: %d (in binary %064b)\n", result, result)
}
