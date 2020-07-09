package application

import (
	"github.com/keita7i/agility/agile"
)

type Service struct {
	JIRAService JIRAService
}

func (s Service) LastSprints(max int) ([]agile.Sprint, error) {
	return s.JIRAService.LastSprints(max)
}
