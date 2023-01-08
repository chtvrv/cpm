package helpers

import (
	"cpm/models"
	"cpm/utils"
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type Writer struct{}

func (writer *Writer) WriteResults(results []models.Result, outputDataFilepath string, criticalPathFilepath string) error {
	records := writer.makeRecords(results, func(result models.Result) []string {
		return []string{
			result.WorkName,
			strconv.Itoa(result.Duration),
			strconv.Itoa(result.EarlyStart),
			strconv.Itoa(result.LateStart),
			strconv.Itoa(result.EarlyFinish),
			strconv.Itoa(result.LateFinish),
			strconv.Itoa(result.TimeMargin),
		}
	})
	columns := []string{"stage", "duration", "early_start", "late_start", "early_finish", "late_finish", "time_margin"}
	err := writer.writeCSVRecords(outputDataFilepath, columns, records)
	if err != nil {
		log.Fatal(err)
	}

	var criticalPathResults []models.Result
	for i := range results {
		if results[i].EarlyStart == results[i].LateStart {
			criticalPathResults = append(criticalPathResults, results[i])
		}
	}
	records = writer.makeRecords(criticalPathResults, func(result models.Result) []string {
		return []string{
			result.WorkName,
			strconv.Itoa(result.Duration),
			strconv.Itoa(result.EarlyStart),
			strconv.Itoa(result.EarlyFinish),
		}
	})
	columns = []string{"stage", "duration", "early_start", "early_finish"}
	err = writer.writeCSVRecords(criticalPathFilepath, columns, records)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (writer *Writer) makeRecords(results []models.Result, transform func(models.Result) []string) [][]string {
	var records [][]string
	for i := range results {
		records = append(records, transform(results[i]))
	}
	return records
}

func (writer *Writer) writeCSVRecords(filepath string, columns []string, records [][]string) error {
	file, err := os.Create(filepath)
	defer utils.Close(file.Close)
	if err != nil {
		log.Fatal("failed to open file", err)
	}
	csvWriter := csv.NewWriter(file)
	csvWriter.Comma = '|'
	err = csvWriter.Write(columns)
	if err != nil {
		log.Fatal("failed to write records", err)
	}
	csvWriter.Flush()
	err = csvWriter.WriteAll(records)
	if err != nil {
		log.Fatal("failed to write records", err)
	}
	return nil
}
