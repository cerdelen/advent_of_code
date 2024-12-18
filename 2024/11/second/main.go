package main

import (
    "fmt"
)

func get_digits(num int) []int {
	var out []int
	for num > 0 {
		digit := num % 10
		out = append([]int{digit}, out...)
		num /= 10
	}
	return out
}

const max_depth = 75
type KEY struct {
    val, depth int
}
var memo map[KEY]int

func count_resulting_stones(val int, depth int) int {
    if depth == max_depth {
        return 1
    }
    key := KEY{val, depth}
    if memo[key] != 0 {
		return memo[key]
	}

    //  If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
    if val == 0 {
        stones := count_resulting_stones(1, depth + 1)
        memo[key] = stones
        return stones
    }

    //  If the stone is engraved with a number that has an even number of digits, it is replaced by two stones.
    //      The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone.
    //      (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
    digits := get_digits(val)
    if len(digits) % 2 == 0 {
        left_value := 0
        for j := 0; j < len(digits) / 2; j++ {
            left_value *= 10
            left_value += digits[j]
        }
        right_value := 0
        for j := len(digits) / 2; j < len(digits); j++ {
            right_value *= 10
            right_value += digits[j]
        }

        stones := count_resulting_stones(left_value, depth + 1) + count_resulting_stones(right_value, depth + 1)
        memo[key] = stones
        return stones
    }

    //  If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.
    stones := count_resulting_stones(val * 2024, depth + 1)
    memo[key] = stones
    return stones
}

func main() {
    memo = make(map[KEY]int)

    var stones = []int {1950139, 0, 3, 837, 6116, 18472, 228700, 45}

    result := 0

    for i := 0; i < len(stones); i++{
        result += count_resulting_stones(stones[i], 0)
    }

    fmt.Printf("result: %d\n", result)
}
