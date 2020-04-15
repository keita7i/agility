package application

import "github.com/keitam913/agility/agile"

type JIRAService interface {
	LastSprints(max int) ([]agile.Sprint, error)
}
