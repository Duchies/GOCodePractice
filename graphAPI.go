package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
Topic : Graph traversal algorithms (DFS with backtracking) witn API testing. 

Input 
{
  "edges": [[0, 1], [0, 2], [1, 2], [1, 3], [2, 3], [3, 4]],
  "start": 0,
  "end": 4
}

output
[
  [0,1,3,4],
  [0,2,3,4],
  [0,1,2,3,4]
]
*/

func main() {

	http.HandleFunc("/find-paths", findPaths)

	err := http.ListenAndServe(":3333", nil)

	if err != nil {
		fmt.Println("error occured", err)
		log.Printf(err.Error())
	}

}

var result = make([][]int, 0)

func findPaths(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /find-paths request\n")

	// edges := [][]int{
	// 	{0, 1},
	// 	{0, 2},
	// 	{1, 2},
	// 	{1, 3},
	// 	{2, 3},
	// 	{3, 4},
	// }

	// AdjacencyList represents a graph as a list of vertices and edges.
	// Where each vertex is connected to a set of vertices.
	adjacencyList := map[int][]int{
		0: {1, 2},
		1: {2, 3},
		2: {3},
		3: {4},
	}

	visited := make([]bool, 5)

	findPathsHelper(adjacencyList, 0, 4, []int{}, visited)
	fmt.Println("Ans : ", result)
}

func findPathsHelper(adjacencyList map[int][]int, startNode, endNode int, currList []int, visited []bool) {

	currList = append(currList, startNode)
	if startNode == endNode {
		result = append(result, currList)
		return
	}

	edges := adjacencyList[startNode]

	visited[startNode] = true
	for _, edge := range edges {
		if visited[edge] == false {
			findPathsHelper(adjacencyList, edge, endNode, currList, visited)
		}
	}
	visited[startNode] = false

}


// RUN :   go run for running code 
// open another terminal : $ curl http://localhost:3333/find-paths
