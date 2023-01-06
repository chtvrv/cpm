package helpers

import (
	"cpm/models"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Writer struct{}

func (writer *Writer) WriteResults(filepath string, works []models.WorkInfo, forward []int, reversed []int) error {
	var results []models.Result
	for i := 0; i < len(works); i++ {
		results = append(results, models.Result{
			WorkName:    works[i].Name,
			Duration:    works[i].Duration,
			EarlyStart:  forward[2*i],
			LateStart:   reversed[2*i],
			EarlyFinish: forward[2*i+1],
			LateFinish:  reversed[2*i+1],
			TimeMargin:  reversed[2*i] - forward[2*i],
		})
	}

	file, err := os.Create(filepath)
	defer file.Close()
	if err != nil {
		log.Fatal("failed to open file", err)
		return err
	}

	csvWriter := csv.NewWriter(file)
	for i := range results {
		record := []string{
			results[i].WorkName,
			strconv.Itoa(results[i].Duration),
			strconv.Itoa(results[i].EarlyStart),
			strconv.Itoa(results[i].LateStart),
			strconv.Itoa(results[i].EarlyFinish),
			strconv.Itoa(results[i].LateFinish),
			strconv.Itoa(results[i].TimeMargin),
		}
		fmt.Println(record)
		err := csvWriter.Write(record)
		if err != nil {
			log.Fatal(err)
		}
	}
	csvWriter.Flush()
	return nil
}
