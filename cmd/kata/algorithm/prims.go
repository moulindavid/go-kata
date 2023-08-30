package main

import (
	"fmt"
	"math"
	"moulindavid/go-kata/cmd/kata/datastruct"
)

func prims(list datastruct.WeightedAdjacencyList) datastruct.WeightedAdjacencyList {
	// Prims algorithm : minimum spanning tree
	// What is a minimum spanning tree?
	//  - Requires no cycles
	//  - For it to technically be a minimum spanning tree, the graph requires
	//  to be strongly connected
	// 1. Select starting node,
	// 2. put edges of current selected node into a list
	// 3. select edge that is the lowest value and go to a node we haven't seen yet
	// 4. we need to insert the edge from current to new into our mst
	// 5. the newly selected node becomes the current node,
	// 6. repeat to step 2 until unvisited is empty or unreachable
	// ...
	// 8. Profit

	visited := make([]bool, len(list))
	mst := make([][]datastruct.GraphEdge, len(list))

	// Step 1.

	visited[0] = true
	current := 0

	var edges []datastruct.EdgeTuple

	for containsFalse(visited) && current >= 0 {
		fmt.Println(visited)
		// Step 2. all edges into the list
		for _, edge := range list[current] {
			edges = append(edges, datastruct.EdgeTuple{Current: current, Edge: edge})
		}
		// Step 3. select lowest value edge going to node we havent seen yet
		lowest := math.MaxInt64
		lowestEdge := datastruct.EdgeTuple{Current: -1, Edge: datastruct.GraphEdge{To: -1, Weight: -1}}

		for _, edge := range edges {
			fmt.Println(edge)
			if !visited[edge.Edge.To] && edge.Edge.Weight < lowest {
				lowest = edge.Edge.Weight
				lowestEdge = edge
			}
		}

		// 4. we need to insert the edge from current into our mst, set visited, and remove the potential edge
		if lowestEdge.Edge.To != -1 {
			mst[lowestEdge.Current] = append(mst[lowestEdge.Current], lowestEdge.Edge)
			graphEdge := datastruct.GraphEdge{
				To:     lowestEdge.Current,
				Weight: lowestEdge.Edge.Weight,
			}
			mst[lowestEdge.Edge.To] = append(mst[lowestEdge.Edge.To], graphEdge)

			visited[lowestEdge.Edge.To] = true

			for i, edge := range edges {
				if edge.Current == lowestEdge.Current && edge.Edge == lowestEdge.Edge {
					edges = append(edges[:i], edges[i+1:]...)
					break
				}
			}
			current = lowestEdge.Edge.To
		}

		for vertex, edges := range mst {
			fmt.Printf("Vertex %d:\n", vertex)
			for _, edge := range edges {
				fmt.Printf("----> {to: %d, weight: %d}\n", edge.To, edge.Weight)
			}
			fmt.Println()
		}
	}
	return mst
}

func containsFalse(arr []bool) bool {
	for _, v := range arr {
		if v == false {
			return true
		}
	}
	return false
}

func main() {
	weightedAdjacencyList := [][]datastruct.GraphEdge{
		{{To: 2, Weight: 1}, {To: 1, Weight: 3}},
		{{To: 0, Weight: 3}, {To: 4, Weight: 1}, {To: 3, Weight: 3}},
		{{To: 0, Weight: 1}, {To: 3, Weight: 7}},
		{{To: 6, Weight: 1}, {To: 1, Weight: 3}, {To: 2, Weight: 7}},
		{{To: 1, Weight: 1}, {To: 5, Weight: 2}},
		{{To: 4, Weight: 2}, {To: 6, Weight: 1}},
		{{To: 5, Weight: 1}, {To: 3, Weight: 1}},
	}

	fmt.Println(prims(weightedAdjacencyList))
}
