package usecase

import "github.com/keitam0/agility/domain/agile"

type JIRAService interface {
	BoardOfTeam(team string, maxSprints int) (agile.Board, error)
}
