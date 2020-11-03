package usecase

import "github.com/keitam0/agility/domain/agile"

type JIRAService interface {
	LastSprints(max int) ([]agile.Sprint, error)
	BoardOfTeam(team string, maxSprints int) (agile.Board, error)
}
