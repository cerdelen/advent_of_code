package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var keys [][]int
var locks [][]int

func key_fits(key, lock []int) bool {
    for i, _ := range lock {
        if lock[i] + key[i] > 5 {
            return false
        }
    }
    return true
}

func parse_key_or_lock(key_or_lock []string) {
    var int_key_lock []int = make([]int, 5)
	switch key_or_lock[0][0] {
        case '#': {
            for i, _ := range key_or_lock {
                for j := 0; j < 5; j++ {
                    if key_or_lock[i][j] == '#' {
                        int_key_lock[j] = i
                    }
                }
            }
            locks = append(locks, int_key_lock)
		}
        case '.': {
            for i, _ := range key_or_lock {
                for j := 0; j < 5; j++ {
                    if key_or_lock[i][j] == '.' {
                        int_key_lock[j] = 5 - i
                    }
                }
            }
            keys = append(keys, int_key_lock)
		}
        default : {
            log.Fatal("wrong first char of key/lok")
        }
	}
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
		key_or_lock := make([]string, 0)

		for len(line) != 0 {
			key_or_lock = append(key_or_lock, line)
			scanner.Scan()
			line = scanner.Text()
		}

		parse_key_or_lock(key_or_lock)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    for _, key := range keys {
        for _, lock := range locks {
            if key_fits(key, lock) {
                result++
            }
        }
    }

	fmt.Printf("result: %d\n", result)
}
