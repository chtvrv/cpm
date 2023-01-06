package main

import (
	"cpm/helpers"
	"flag"
)

var args struct {
	inputFilePath  string
	outputFilePath string
}

func main() {
	flag.StringVar(&args.inputFilePath, "i", "", "Path to the input file")
	flag.StringVar(&args.outputFilePath, "o", "", "Path to the output file")
	flag.Parse()
	runner := helpers.Runner{}
	runner.Run(args.inputFilePath, args.outputFilePath)
}
