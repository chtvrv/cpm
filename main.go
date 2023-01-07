package main

import (
	"cpm/helpers"
	"flag"
	"log"
)

var args struct {
	inputFilePath         string
	outputResultsFilePath string
	criticalPathFilePath  string
}

func main() {
	flag.StringVar(&args.inputFilePath, "i", "", "Path to the input file")
	flag.StringVar(&args.outputResultsFilePath, "o", "", "Path to the output file")
	flag.StringVar(&args.criticalPathFilePath, "c", "", "Path to the critical output file")
	flag.Parse()
	runner := helpers.Runner{}
	err := runner.Run(args.inputFilePath, args.outputResultsFilePath, args.criticalPathFilePath)
	if err != nil {
		log.Fatal(err)
	}
}
