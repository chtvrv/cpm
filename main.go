package main

import (
	cpm_builder "cpm/builder"
	cpm_parser "cpm/parser"
	"fmt"
	s "github.com/inancgumus/prettyslice"
)

func main() {
	parser := cpm_parser.CreateParser()
	works := parser.Parse("example_path")
	builder := cpm_builder.CreateBuilder()
	adj_matrix, inputs, outputs := builder.BuildAdjacencyMatrix(works)
	s.MaxPerLine = 30
	s.PrintHex = true
	for i := range adj_matrix.Data {
		s.Show("row", adj_matrix.Data[i])
	}
	fmt.Println(inputs, outputs)
}
