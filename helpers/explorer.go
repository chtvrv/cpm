package helpers

import "cpm/models"

type Explorer struct{}

func (explorer *Explorer) TraverseGraph(works []models.WorkInfo, graph *models.AdjMatrix) []models.Result {
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

	var results []models.Result
	for i := 0; i < len(works); i++ {
		results = append(results, models.Result{
			WorkName:    works[i].Name,
			Duration:    works[i].Duration,
			EarlyStart:  forward[2*i],
			LateStart:   reversed[2*i],
			EarlyFinish: forward[2*i+1],
			LateFinish:  reversed[2*i+1],
			TimeMargin:  reversed[2*i] - forward[2*i],
		})
	}

	return results
}
