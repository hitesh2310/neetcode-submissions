type Edge struct {
	Node int
	Weight int 
}

type MinHeap []Edge 

func (h MinHeap) Len() int{
	return len(h)
}

func (h MinHeap) Swap(i,j int) {
	h[i],h[j] = h[j],h[i]
}

func (h MinHeap) Less(i,j int) bool {
	return h[i].Weight < h[j].Weight
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h,x.(Edge))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}


func networkDelayTime(times [][]int, n int, k int) int {
  edges := make(map[int][]Edge)
	for _,time := range times {
		u,v,w := time[0],time[1], time[2]
		edges[u] = append(edges[u], Edge{Node:v, Weight: w})
	}


    visited := make(map[int]bool)
	t := 0
	pq := &MinHeap{}
	heap.Init(pq)
	heap.Push(pq, Edge{ Node: k, Weight:0} )
	
	for pq.Len() > 0 {
		current := heap.Pop(pq).(Edge)
		node, time := current.Node, current.Weight

		if visited[node] {
			continue
		}
		t=time
		visited[node] = true
		for _,neighbor := range edges[node] {

			if visited[neighbor.Node] {
				continue
			}
			heap.Push(pq,Edge{Node:neighbor.Node, Weight: time + neighbor.Weight})
		}
	}

	if len(visited) == n {
		return t
	}

	return -1
}
