package datastruct

type GraphEdge struct {
	To     int
	Weight int
}

type WeightedAdjacencyList [][]GraphEdge

type EdgeTuple struct {
	Current int
	Edge    GraphEdge
}
