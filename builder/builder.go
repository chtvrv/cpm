package builder

import "cpm/models"

type Builder interface {
	BuildAdjacencyMatrix(works *[]models.WorkInfo) (matrix *models.AdjMatrix, inputs []int, outputs []int)
}
