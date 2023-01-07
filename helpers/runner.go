package helpers

type Runner struct{}

func (runner *Runner) Run(inputFile string, outputFile string, criticalPath string) error {
	parser := Parser{}
	builder := Builder{}
	explorer := Explorer{}
	writer := Writer{}

	works := parser.ParseInput(inputFile)
	adjMatrix := builder.BuildAdjacencyMatrix(works)
	results := explorer.TraverseGraph(works, adjMatrix)
	return writer.WriteResults(results, outputFile, criticalPath)
}
