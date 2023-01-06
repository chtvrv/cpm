package parser

import "cpm/models"

type Parser interface {
	Parse(filepath string) *[]models.WorkInfo
}
