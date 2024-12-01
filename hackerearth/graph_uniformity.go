package hackerearth

import "fmt"

//https://www.hackerearth.com/practice/algorithms/graphs/graph-representation/practice-problems/algorithm/uniformity-794d7bdc/
func GraphUniformity() {
	// abcbabcbcbbacbacbacbbabcbacbbacbabcbbabc
	var n, k int
	fmt.Scan(&n, &k)
	var s string
	fmt.Scan(&s)

	var p1, p2 int
	ans := 0
	powerUsed := 0
	current := 0

	// For a
	for p2 = 0; p2 < n; {
		if s[p2] == 'a' {
			current++
			p2++
		} else {
			if k > powerUsed {
				powerUsed++
				current++
				p2++
			} else if k == 0 {
				p2++
				current = 0
			} else {
				for p1 < n {
					if s[p1] == 'a' {
						current--
						p1++
					} else {
						powerUsed--
						p1++
						current--
						break
					}
				}
			}
		}

		if current > ans {
			ans = current
		}
	}

	powerUsed = 0
	current = 0
	p1 = 0
	p2 = 0

	for p2 = 0; p2 < n; {
		if s[p2] == 'b' {
			current++
			p2++
		} else {
			if k > powerUsed {
				powerUsed++
				current++
				p2++
			} else if k == 0 {
				p2++
				current = 0
			} else {
				for p1 < n {
					if s[p1] == 'b' {
						current--
						p1++
					} else {
						powerUsed--
						p1++
						current--
						break
					}
				}
			}
		}

		if current > ans {
			ans = current
		}
	}

	powerUsed = 0
	current = 0
	p1 = 0
	p2 = 0

	for p2 = 0; p2 < n; {
		if s[p2] == 'c' {
			current++
			p2++
		} else {
			if k > powerUsed {
				powerUsed++
				current++
				p2++
			} else if k == 0 {
				p2++
				current = 0
			} else {
				for p1 < n {
					if s[p1] == 'c' {
						current--
						p1++
					} else {
						powerUsed--
						p1++
						current--
						break
					}
				}
			}
		}

		if current > ans {
			ans = current
		}
	}

	fmt.Println(ans)
}
