package helpers

import (
	"cpm/models"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Parser struct{}

func (parser *Parser) ParseInput(filepath string) []models.WorkInfo {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.FieldsPerRecord = 3
	reader.TrimLeadingSpace = true
	var works []models.WorkInfo
	var records [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		records = append(records, record)
	}
	for i := 0; i < len(records); i++ {
		if i == 0 {
			continue
		}
		name := records[i][0]
		duration, err := strconv.Atoi(records[i][1])
		if err != nil {
			log.Fatal(err)
		}
		subworks := strings.Fields(records[i][2])
		works = append(works, models.WorkInfo{Name: name, Duration: duration, Subworks: subworks})
	}
	return works
}
