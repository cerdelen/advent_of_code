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
    parents   []*State
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

    visited := make(map[Pos]map[Dir]*State)

    initialState := &State{
		position:  start,
		direction: startDir,
		cost:      0,
		priority:  heuristic(start, end),
		parents:   nil,
	}
	heap.Push(pq, initialState)

    var goalStates []*State
	minCostToGoal := math.MaxInt
    for pq.Len() > 0 {
		current := heap.Pop(pq).(*State)

        if current.cost > minCostToGoal {
			break
		}
        if current.position == end {
			if current.cost < minCostToGoal {
				goalStates = []*State{current}
				minCostToGoal = current.cost
			} else if current.cost == minCostToGoal {
				goalStates = append(goalStates, current)
			}
			continue
		}

		if visited[current.position] == nil {
			visited[current.position] = make(map[Dir]*State)
		}
		if _, alreadyVisited := visited[current.position][current.direction]; alreadyVisited {
			if visited[current.position][current.direction].cost == current.cost {
				visited[current.position][current.direction].parents = append(
					visited[current.position][current.direction].parents, current.parents...,
				)
			}
			continue
		}
		visited[current.position][current.direction] = current

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
                parents:   []*State{current},
			})
		}

		nextDir := change_clockwise(current.direction)
		heap.Push(pq, &State{
			position:  current.position,
			direction: nextDir,
			cost:      current.cost + 1000,
			priority:  current.cost + 1000 + heuristic(current.position, end),
            parents:   []*State{current},
		})

		nextDir = change_counter_clockwise(current.direction)
		heap.Push(pq, &State{
			position:  current.position,
			direction: nextDir,
			cost:      current.cost + 1000,
			priority:  current.cost + 1000 + heuristic(current.position, end),
            parents:   []*State{current},
		})
	}

    uniquePositions := make(map[Pos]struct{})
    fmt.Printf("goalStates: %d\n", len(goalStates))
    for _, path := range goalStates {
        count_unique(*path, uniquePositions)
    }

    fmt.Printf("uniquePositions: %d\n", len(uniquePositions))
	return -1
}

func count_unique(state State, unique_pos map[Pos]struct{}) {
    for _, states := range state.parents {
        count_unique(*states, unique_pos)
    }
    unique_pos[state.position] = struct{}{}
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
