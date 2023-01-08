package helpers

import (
	"cpm/models"
	ps "github.com/inancgumus/prettyslice"
)

type Builder struct{}

func (builder *Builder) BuildAdjacencyMatrix(works []models.WorkInfo, shouldPrint bool) (matrix *models.AdjMatrix) {
	dimension := len(works)*2 + 2
	data := make([][]int, dimension)
	for i := range data {
		row := make([]int, dimension)
		for j := range row {
			row[j] = -1
		}
		data[i] = row
	}

	memo := map[string]int{}
	for id, work := range works {
		memo[work.Name] = id
		data[id*2+1][id*2+2] = work.Duration
		if !work.HasSubworks() {
			data[0][id*2+1] = 0
		} else {
			for _, subwork := range work.Subworks {
				if idw, ok := memo[subwork]; ok {
					data[idw*2+2][id*2+1] = 0
				}
			}
		}
	}

	for i := 0; i < len(works); i++ {
		if isOutput(data[i*2+2]) {
			data[i*2+2][dimension-1] = 0
		}
	}
	ps.MaxPerLine = 30
	ps.PrintHex = true
	for i := range data {
		ps.Show("row", data[i])
	}
	matrix = &models.AdjMatrix{Data: data, RowsCount: dimension, ColumnsCount: dimension}
	return
}

func isOutput(row []int) bool {
	for i := range row {
		if row[i] != -1 {
			return false
		}
	}
	return true
}
