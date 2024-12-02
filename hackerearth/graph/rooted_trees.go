package graph

import (
	"fmt"
	"main/templates"
)

func traverseAndAdd(adj [][]int, v int, x int, y int, val []int) {
	for i := 0; i < len(adj[v]); i++ {
		if val[adj[v][i]] < y {
			val[adj[v][i]] += x
		}
		traverseAndAdd(adj, adj[v][i], x, y, val)
	}
}

func RootedTrees() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n, q, y int
		templates.TripleInput(&n, &q, &y)

		adj := make([][]int, n+1)
		val := make([]int, n+1)

		for i := 1; i < n; i++ {
			var x, y int
			fmt.Scan(&x, &y)

			adj[x] = append(adj[x], y)
		}

		for ; q > 0; q-- {
			var operation int
			fmt.Scan(&operation)

			var v, x int
			if operation == 0 {
				fmt.Scan(&v, &x)
				traverseAndAdd(adj, v, x, y, val)
			} else {
				fmt.Scan(&v)
				fmt.Println(val[v])
			}
		}
	}
}
