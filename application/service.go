package application

import (
	"github.com/keitam0/agility/agile"
)

const MAX_SPRINTS = 8

type Service struct {
	JIRAService JIRAService
}

func (s Service) LastSprints(max int) ([]agile.Sprint, error) {
	return s.JIRAService.LastSprints(max)
}

func (s Service) BoardOfTeam(team string) (agile.Board, error) {
	return s.JIRAService.BoardOfTeam(team, MAX_SPRINTS)
}
