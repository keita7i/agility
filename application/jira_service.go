package application

import "github.com/keita7i/agility/agile"

type JIRAService interface {
	LastSprints(max int) ([]agile.Sprint, error)
}
