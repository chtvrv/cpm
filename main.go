package main

import (
	cpmbuilder "cpm/builder"
	cpmparser "cpm/parser"
	"flag"
	"fmt"
	s "github.com/inancgumus/prettyslice"
)

var args struct {
	inputFilePath string
}

func main() {
	flag.StringVar(&args.inputFilePath, "c", "", "Path to the input file")
	flag.Parse()

	parser := cpmparser.CreateParser()
	works := parser.Parse(args.inputFilePath)
	builder := cpmbuilder.CreateBuilder()
	adjMatrix, inputs, outputs := builder.BuildAdjacencyMatrix(works)
	s.MaxPerLine = 30
	s.PrintHex = true
	for i := range adjMatrix.Data {
		s.Show("row", adjMatrix.Data[i])
	}
	fmt.Println(inputs, outputs)
}
