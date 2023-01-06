package models

type WorkInfo struct {
	Name     string
	Duration int
	Subworks []string
}

func (info *WorkInfo) HasSubworks() bool {
	return len(info.Subworks) >= 1 && info.Subworks[0] != "-"
}
