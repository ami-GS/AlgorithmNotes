package main

import "fmt"

type Node struct {
	val int
	nxt *Node
}

// NOTE: need to be careful, node.nxt does not show beggining of list, out.nxt

func mergeKList(lists []*Node) *Node {
	node := &Node{0, nil}
	out := node
	for {
		merging := false
		nxtLoc := 0
		small := 1 << 21
		for k, np := range lists {
			if np != nil {
				if np.val < small {
					small = np.val
					nxtLoc = k
				}
				merging = true
			}
		}
		if !merging {
			return out.nxt
		}
		node.nxt = lists[nxtLoc]
		lists[nxtLoc] = lists[nxtLoc].nxt
		node = node.nxt
	}
}

func main() {
	K := 3
	N := 5
	lists := make([]*Node, K)

	for k, _ := range lists {
		lists[k] = &Node{k, nil}
		np := lists[k]
		for i := 1; i < N; i++ {
			//fmt.Println(fmt.Sprintf("%p", np))
			np.nxt = &Node{i*(k+1) - k/2, nil}
			np = np.nxt
		}
		np = lists[k]
		for i := 0; i < N; i++ {
			fmt.Print(np.val, "->")
			np = np.nxt
		}
		fmt.Println()
	}
	out := mergeKList(lists)
	for i := 0; i < N*K; i++ {
		fmt.Print(out.val, "->")
		out = out.nxt
	}
	fmt.Println()
}
