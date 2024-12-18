package main

import (
    "os"
    "fmt"
    "log"
)

const rounds = 25

func get_digits(num int) []int {
	var out []int
	for num > 0 {
		digit := num % 10
		out = append([]int{digit}, out...)
		num /= 10
	}
	return out
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var stones = []int {1950139, 0, 3, 837, 6116, 18472, 228700, 45}

    for blinks := 0; blinks < rounds; blinks++{
        fmt.Println("Blink nr. ", blinks + 1)
        for i := 0; i < len(stones); i++{
            //  If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
            if stones[i] == 0 {
                stones[i] = 1
                continue
            }

            //  If the stone is engraved with a number that has an even number of digits, it is replaced by two stones.
            //      The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone.
            //      (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
            digits := get_digits(stones[i])
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

                stones[i] = left_value
                stones = append(stones[:i + 1], append([]int{right_value}, stones[i + 1:]...)...)
                i++
                continue
            }

            //  If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.
            stones[i] = stones[i] * 2024
        }
    }


    fmt.Println(stones)

    fmt.Printf("result: %d\n", len(stones))
}
