package application

import (
	"github.com/keitam913/agility/agile"
)

type Service struct {
	JIRAService JIRAService
}

func (s Service) LastSprints(n int) ([]agile.Sprint, error) {
	var sps []agile.Sprint
	lastSprint := 13
	for sp := lastSprint; lastSprint-n < sp; sp-- {
		sp, err := s.JIRAService.Sprint(sp)
		if err != nil {
			return nil, err
		}
		sps = append(sps, sp)
	}
	return sps, nil
}
