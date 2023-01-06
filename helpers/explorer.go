package helpers

import "cpm/models"

type Explorer struct{}

func (explorer *Explorer) TraverseGraph(graph *models.AdjMatrix) ([]int, []int) {
	forward := make([]int, graph.RowsCount)
	forward[0] = 0

	for i := 0; i < graph.RowsCount; i++ {
		for j := 0; j < i; j++ {
			if graph.Data[j][i] != -1 && forward[j]+graph.Data[j][i] > forward[i] {
				forward[i] = forward[j] + graph.Data[j][i]
			}
		}
	}

	reversed := make([]int, graph.RowsCount)
	for i := 0; i < len(forward); i++ {
		reversed[i] = forward[len(forward)-1]
	}

	for i := graph.RowsCount - 1; i >= 0; i-- {
		for j := i; j < graph.RowsCount-1; j++ {
			if graph.Data[i][j] != -1 && forward[j]-graph.Data[i][j] < reversed[i] {
				reversed[i] = reversed[j] - graph.Data[i][j]
			}
		}
	}

	forward = forward[1 : len(forward)-1]
	reversed = reversed[1 : len(reversed)-1]
	return forward, reversed
}
