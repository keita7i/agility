package usecase

import "github.com/keitam0/agility/domain/agile"

type JIRAServiceMock struct {
	BoardOfTeamFunc func(team string, maxSprints int) (agile.Board, error)
}

func (s *JIRAServiceMock) BoardOfTeam(team string, maxSprints int) (agile.Board, error) {
	return s.BoardOfTeamFunc(team, maxSprints)
}
