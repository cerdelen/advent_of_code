package main

import (
    "container/heap"
    "math"
    "os"
    "bufio"
    "fmt"
    "log"
    // "sort"
    // "io/ioutil"
    "strings"
    // "strconv"
)

type State struct {
	position  Pos
	direction Dir
	cost      int
	priority  int
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
const MaxInt = 2147483647

func change_clockwise(direction Dir) Dir {
    return Dir{y_dir: direction.x_dir, x_dir: -direction.y_dir}
}

func change_counter_clockwise(direction Dir) Dir {
    return Dir{y_dir: -direction.x_dir, x_dir: direction.y_dir}
}

type Pos struct {
	y, x int
}

type Dir struct {
	y_dir, x_dir int
}

// func find_min_path_size(grid [][]rune, visited_spots map[Pos]struct{}, position Pos, direction Dir, current_value int, current_best *int) int {
//     // if position.x > len(grid[0]) || position.x < 0 || position.y > len(grid) || position.y < 0 {
//     //     return MaxInt
//     // }
//
//     if grid[position.y][position.x] == '#' || current_value >= *current_best {
//         return MaxInt
//     }
//
//     if _, exists := visited_spots[position]; exists {
//         return MaxInt
// 	}
//     //
//     //
//     // if grid[position.y][position.x] == 'S' {
//     //     return MaxInt
//     // }
//     //
//     if grid[position.y][position.x] == 'E' {
//         return current_value
//     }
//
//     // visited_spots[position] = struct{}{}
//     visited_spots[position] = struct{}{}
//     defer delete(visited_spots, position)
//
//     var temp_value int = MaxInt
//     // straight
//     straight := find_min_path_size(grid, visited_spots, Pos{position.y + direction.y_dir, position.x + direction.x_dir}, direction, current_value + 1, current_best)
//     temp_value = min(temp_value, straight)
//
//     // counter clockwise
//     var temp_dir Dir = change_counter_clockwise(direction)
//     counter_clockwise := find_min_path_size(grid, visited_spots, Pos{position.y + temp_dir.y_dir, position.x + temp_dir.x_dir}, temp_dir, current_value + 1001, current_best)
//     temp_value = min(temp_value, counter_clockwise)
//
//     // clockwise
//     temp_dir  = change_clockwise(direction)
//     clockwise := find_min_path_size(grid, visited_spots, Pos{position.y + temp_dir.y_dir, position.x + temp_dir.x_dir}, temp_dir, current_value + 1001, current_best)
//     temp_value = min(temp_value, clockwise)
//
//     delete(visited_spots, position)
//     fmt.Printf("curr best: %d\n", *current_best)
//
//     *current_best = min(*current_best, temp_value)
//     return temp_value
// }

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

func heuristic(curr, end Pos) int {
	return int(math.Abs(float64(end.y-curr.y)) + math.Abs(float64(end.x-curr.x)))
}

func a_star(grid [][]rune, start Pos, end Pos, startDir Dir) int {
    pq := &PriorityQueue{}
	heap.Init(pq)

	visited := make(map[Pos]map[Dir]bool)

	heap.Push(pq, &State{
		position:  start,
		direction: startDir,
		cost:      0,
		priority:  heuristic(start, end),
	})

    for pq.Len() > 0 {
		current := heap.Pop(pq).(*State)

		if current.position == end {
			return current.cost
		}

		if visited[current.position] == nil {
			visited[current.position] = make(map[Dir]bool)
		}
		if visited[current.position][current.direction] {
			continue
		}
		visited[current.position][current.direction] = true

		nextPos := Pos{
			y: current.position.y + current.direction.y_dir,
			x: current.position.x + current.direction.x_dir,
		}
		if grid[nextPos.y][nextPos.x] != '#' {
			heap.Push(pq, &State{
				position:  nextPos,
				direction: current.direction,
				cost:      current.cost + 1,
				priority:  current.cost + 1 + heuristic(nextPos, end),
			})
		}

		nextDir := change_clockwise(current.direction)
		heap.Push(pq, &State{
			position:  current.position,
			direction: nextDir,
			cost:      current.cost + 1000,
			priority:  current.cost + 1000 + heuristic(current.position, end),
		})

		nextDir = change_counter_clockwise(current.direction)
		heap.Push(pq, &State{
			position:  current.position,
			direction: nextDir,
			cost:      current.cost + 1000,
			priority:  current.cost + 1000 + heuristic(current.position, end),
		})
	}

	return -1
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    var result = 0

    var position, end Pos

    var grid [][] rune
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        row := []rune(line)

        ind := strings.Index(line, "S")
        if ind != -1 {
            position = Pos{len(grid), ind}
        }

        ind = strings.Index(line, "E")
        if ind != -1 {
            end = Pos{len(grid), ind}
        }

        grid = append(grid, row)
    }

    dir := Dir{0, 1}


    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    result = MaxInt

    result = a_star(grid, position, end, dir)

    fmt.Printf("result: %d\n", result)
}
