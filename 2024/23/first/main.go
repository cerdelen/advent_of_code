package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
)

type Graph struct {
    m map[string][]string
}

func (g *Graph) AddEdge(base_node, node_two string) {
    if g.m == nil {
        g.m = map[string][]string{}
    }
    g.m[base_node] = append(g.m[base_node], node_two)
    g.m[node_two] = append(g.m[node_two], base_node)
}

func (g *Graph) Print() {
    for node, neighbors := range g.m {
        fmt.Printf("%s -> %v\n", node, neighbors)
    }
}

func Contains[T comparable](slice []T, item T) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var graph Graph
    m := make(map[[3]string]struct{})

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        re := regexp.MustCompile(`^([a-z]+)-([a-z]+)$`)

        matches := re.FindStringSubmatch(line)
        if matches == nil {
            log.Fatal("mismatched node Pair! ", line)
        }
        base_node := matches[1]
        node_two := matches[2]
        graph.AddEdge(base_node, node_two)
    }

    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    for base_node, edges := range graph.m {
        if base_node[0] == 't' {
            for i := 0; i < len(edges); i++ {
                for j := i + 1; j < len(edges); j++ {
                    if Contains(graph.m[edges[i]], edges[j]) {
                        trio := []string{base_node, edges[i], edges[j]}
                        sort.Strings(trio)
                        var sorted_trio [3]string
                        copy(sorted_trio[:], trio)
                        if _, exists := m[sorted_trio]; !exists {
                            m[sorted_trio] = struct{}{}
                        }
                    }
                }
            }

        }
    }

    fmt.Printf("result: %d\n", len(m))
}
