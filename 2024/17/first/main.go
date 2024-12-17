package main

import (
	"math"
	"fmt"
	"log"
)

const inst_adv = 0
const inst_bxl = 1
const inst_bst = 2
const inst_jnz = 3
const inst_bxc = 4
const inst_out = 5
const inst_bdv = 6
const inst_cdv = 7

func combo_value(reg_a *int, reg_b *int, reg_c *int, literal int) int {
    switch literal {
        case 4: {
            return *reg_a
        }
        case 5: {
            return *reg_b
        }
        case 6: {
            return *reg_c
        }
        case 7: {
            log.Fatal("Case 7 in combo value")
        }
        default: {
            return literal
        }
    }
    return literal
}

func adv(reg_a *int, reg_b *int, reg_c *int, literal int) {
    *reg_a = *reg_a / int(math.Pow(2, float64(combo_value(reg_a, reg_b, reg_c, literal))))
}

func bxl(reg_b *int, literal int) {
    *reg_b = *reg_b ^ literal
}

func bst(reg_a *int, reg_b *int, reg_c *int, literal int) {
    *reg_b = combo_value(reg_a, reg_b, reg_c, literal) % 8
}

func jnz(reg_a *int, literal int, ind *int) {
    if *reg_a == 0 {
        return
    }
    *ind = literal
}

func bxc(reg_b *int, reg_c *int) {
    *reg_b = *reg_b ^ *reg_c
}

func out(reg_a *int, reg_b *int, reg_c *int, literal int) {
    fmt.Printf("%d,", int(float64(combo_value(reg_a, reg_b, reg_c, literal))) % 8)
}

func bdv(reg_a *int, reg_b *int, reg_c *int, literal int) {
    *reg_b = *reg_a / int(math.Pow(2, float64(combo_value(reg_a, reg_b, reg_c, literal))))
}

func cdv(reg_a *int, reg_b *int, reg_c *int, literal int) {
    *reg_c = *reg_a / int(math.Pow(2, float64(combo_value(reg_a, reg_b, reg_c, literal))))
}

func run_program(reg_a int, reg_b int, reg_c int, program [16]int) {
    ind := 0
    for ind < len(program) && ind >= 0 {
        switch program[ind] {
            case inst_adv: {
                adv(&reg_a, &reg_b, &reg_c, program[ind + 1])
            }
            case inst_bxl: {
                bxl(&reg_b, program[ind + 1])
            }
            case inst_bst: {
                bst(&reg_a, &reg_b, &reg_c, program[ind + 1])
            }
            case inst_jnz: {
                jnz(&reg_a, program[ind + 1], &ind)
                if reg_a != 0 {
                    ind -= 2
                }
            }
            case inst_bxc: {
                bxc(&reg_b, &reg_c)
            }
            case inst_out: {
                out(&reg_a, &reg_b, &reg_c, program[ind + 1])
            }
            case inst_bdv: {
                bdv(&reg_a, &reg_b, &reg_c, program[ind + 1])
            }
            case inst_cdv: {
                cdv(&reg_a, &reg_b, &reg_c, program[ind + 1])
            }
            default: {
                log.Fatal("not correct command: \n", program[ind])
            }
        }
        ind += 2
    }
}

func main() {
    var reg_a = 66171486
    var reg_b = 0
    var reg_c = 0
    var program = [16]int{2, 4, 1, 6, 7, 5, 4, 6, 1, 4, 5, 5, 0, 3, 3, 0}
    run_program(reg_a, reg_b, reg_c, program)
}
