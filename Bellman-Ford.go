package main

import "fmt"

func main(){
	var edges []Edge
	edges = append(edges, Edge{weight:6, u:1,v:2})
	edges = append(edges, Edge{weight:6, u:1,v:2})
	edges = append(edges, Edge{weight:7, u:1,v:4})
	edges = append(edges, Edge{weight:5, u:2,v:3})
	edges = append(edges, Edge{weight:8, u:2,v:4})
	edges = append(edges, Edge{weight:-4, u:2,v:5})
	edges = append(edges, Edge{weight:-2, u:3,v:2})
	edges = append(edges, Edge{weight:-3, u:4,v:3})
	edges = append(edges, Edge{weight:9, u:4,v:5})
	edges = append(edges, Edge{weight:2, u:5,v:1})
	edges = append(edges, Edge{weight:7, u:5,v:3})
	dis, ok := BellmanFord(edges, 5, 1)
	fmt.Println(ok, dis)
}

// type Vertex struct {
// 	no int
// 	before []*Vertex
// 	next []*Vertex
// }

type Edge struct {
	weight int
	u, v int	//u->v 下标
}
//顶点下标从1开始 vertexnum是实际顶点个数
func BellmanFord(edges []Edge, vertexnum, src int ) ([]int, bool) {
	// 初始化
	var dis = make([]int, vertexnum + 1)
	var preNode = make([]int, vertexnum + 1)
	for i := range dis {
		dis[i] = 1 << 10
	}
	dis[src] = 0

	// 松弛操作
	for range dis {
		for j := 1; j < len(edges); j++ {
			// 如果s到v的距离大于s到u + u到v的距离
			if dis[edges[j].v] > (dis[edges[j].u] + edges[j].weight) {
				dis[edges[j].v] = dis[edges[j].u] + edges[j].weight
				// fmt.Print(dis[edges[j].v], " ")
				preNode[edges[j].v] = edges[j].u
			}
		}
	}

	for i := range edges {
		if dis[edges[i].v] > (dis[edges[i].u] + edges[i].weight) {
			return nil, false
		}
	}
	return dis, true
}
