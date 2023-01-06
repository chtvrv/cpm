package builder

import (
	"cpm/models"
)

type BuilderImpl struct{}

func CreateBuilder() Builder {
	return &BuilderImpl{}
}

func (impl *BuilderImpl) BuildAdjacencyMatrix(works *[]models.WorkInfo) (matrix *models.AdjMatrix, inputs []int, outputs []int) {
	dimension := len(*works)*2 + 2
	data := make([][]int, dimension)
	for i := range data {
		row := make([]int, dimension)
		for j := range row {
			row[j] = -1
		}
		data[i] = row
	}

	memo := map[string]int{}
	for id, work := range *works {
		memo[work.Name] = id
		data[id*2+1][id*2+2] = work.Duration
		if !work.HasSubworks() {
			inputs = append(inputs, id*2+1)
			data[0][id*2+1] = 0
		} else {
			for _, subwork := range work.Subworks {
				if idw, ok := memo[subwork]; ok {
					data[idw*2+2][id*2+1] = 0
				}
			}
		}
	}

	for i := 0; i < len(*works); i++ {
		if IsOutput(data[i*2+2]) {
			outputs = append(outputs, i*2+2)
			data[i*2+2][dimension-1] = 0
		}
	}

	matrix = &models.AdjMatrix{Data: data, RowsCount: dimension, ColumnsCount: dimension}
	return
}

func IsOutput(row []int) bool {
	for i := range row {
		if row[i] != -1 {
			return false
		}
	}
	return true
}
