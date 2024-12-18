package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

const y_max = 70
const x_max = 70

type Pos struct {
    y, x int
}

type State struct {
	position  Pos
	cost      int
	priority  int
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*State))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func heuristic(curr Pos) int {
	return int(math.Abs(float64(70-curr.y)) + math.Abs(float64(70-curr.x)))
}

func push_move(grid [][]byte, pq *PriorityQueue, nextPos Pos, current_cost int, visited map[Pos]int) {
    if nextPos.x > x_max || nextPos.x < 0 || nextPos.y > y_max || nextPos.y < 0 {
        return
    }
    if grid[nextPos.y][nextPos.x] == '#' {
        return
    }

    if prevCost, found := visited[nextPos]; !found || current_cost + 1 < prevCost {
        visited[nextPos] = current_cost + 1
        heap.Push(pq, &State{
            position:  nextPos,
            cost:      current_cost + 1,
            priority:  current_cost + 1 + heuristic(nextPos),
        })
    }


}

func find_min_path(grid [][]byte) int {
    pq := &PriorityQueue{}
	heap.Init(pq)
    end := Pos{70, 70}
    start := Pos{0, 0}

    visited := make(map[Pos]int)
    visited[start] = 0

	heap.Push(pq, &State{
		position:  Pos{0,0},
		cost:      0,
		priority:  heuristic(Pos{0,0}),
	})
    for pq.Len() > 0 {
		current := heap.Pop(pq).(*State)

		if current.position == end {
			return current.cost
		}
        push_move(grid, pq, Pos{current.position.y + 1, current.position.x + 0}, current.cost, visited)
        push_move(grid, pq, Pos{current.position.y - 1, current.position.x + 0}, current.cost, visited)
        push_move(grid, pq, Pos{current.position.y + 0, current.position.x + 1}, current.cost, visited)
        push_move(grid, pq, Pos{current.position.y + 0, current.position.x - 1}, current.cost, visited)
    }
    return 0
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    grid := make([][]byte, 71)
    for i := range grid {
        grid[i] = make([]byte, 71)
        for j := range grid[i] {
            grid[i][j] = '.'
        }
    }
    scanner := bufio.NewScanner(file)
    byte_count := 0
    var result string
    for scanner.Scan() {
        line := scanner.Text()

        pattern := `(\d+),(\d+)`

        re := regexp.MustCompile(pattern)

        matches := re.FindStringSubmatch(line)

        if len(matches) != 3 {
            log.Fatal("Missformed line", line)
        }

        x, _ := strconv.Atoi(matches[1])
        y, _ := strconv.Atoi(matches[2])

        if x > x_max || x < 0 || y > y_max || y < 0 {
            log.Fatal("coordinate out of bounds", line)
        }

        grid[y][x] = '#'

        if byte_count > 1024 {
            if find_min_path(grid) == 0 {
                result = line
                break
            }
        }
        byte_count++
    }

    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    fmt.Printf("result: %s\n", result)
}
