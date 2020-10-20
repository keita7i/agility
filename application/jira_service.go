package application

import "github.com/keitam0/agility/agile"

type JIRAService interface {
	LastSprints(max int) ([]agile.Sprint, error)
	BoardOfTeam(team string, maxSprints int) (agile.Board, error)
}
