package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
)

type Graph struct {
    m map[string][]string
}
var graph Graph

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

func intersect(a, b []string) (res []string) {
	t := make(map[string]bool)
	for _, val := range a {
		t[val] = true
	}
	for _, val := range b {
		if t[val] {
			res = append(res, val)
		}
	}
	return
}

var biggest_lan []string

func BronKerbosch(R, P, X []string) {
    if len(P) == 0 && len(X) == 0 && len(biggest_lan) < len(R) {
        biggest_lan = slices.Clone(R)
        return
    }

    pCopy := slices.Clone(P)
	for _, v := range pCopy {
        newR := append(R, v)
		neighbours := graph.m[v]

		newP := intersect(P, neighbours)
		newX := intersect(X, neighbours)

		BronKerbosch(newR, newP, newX)

		vIdx := slices.Index(P, v)
		P = slices.Delete(P, vIdx, vIdx+1)

		X = append(X, v)
    }
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

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

    nodesList := make([]string, 0, len(graph.m))

	for node := range graph.m {
		nodesList = append(nodesList, node)
	}

    BronKerbosch([]string{}, nodesList, []string{})

    slices.Sort(biggest_lan)

    var result string

    for _, node := range biggest_lan {
        result += node + ","
    }

    fmt.Printf("result: %v\n", result[:len(result)-1])
}
