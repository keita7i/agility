package usecase

import "github.com/keitam0/agility/domain/agile"

type JIRAServiceMock struct {
	LastSprintsFunc func(max int) ([]agile.Sprint, error)
	BoardOfTeamFunc func(team string, maxSprints int) (agile.Board, error)
}

func (s *JIRAServiceMock) LastSprints(max int) ([]agile.Sprint, error) {
	return s.LastSprintsFunc(max)
}

func (s *JIRAServiceMock) BoardOfTeam(team string, maxSprints int) (agile.Board, error) {
	return s.BoardOfTeamFunc(team, maxSprints)
}
