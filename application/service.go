package application

import (
	"github.com/keitam913/agility/agile"
)

type Service struct {
	JIRAService JIRAService
}

func (s Service) LastSprints(max int) ([]agile.Sprint, error) {
	return s.JIRAService.LastSprints(max)
}
