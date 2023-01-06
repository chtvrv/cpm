package helpers

type Runner struct{}

func (runner *Runner) Run(inputFile string, outputFile string) error {
	parser := Parser{}
	builder := Builder{}
	explorer := Explorer{}
	writer := Writer{}

	works := parser.ParseInput(inputFile)
	adjMatrix := builder.BuildAdjacencyMatrix(works)
	forward, reversed := explorer.TraverseGraph(adjMatrix)
	return writer.WriteResults(outputFile, works, forward, reversed)
}
