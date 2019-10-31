package ihelpers

import (
	"github.com/spacepal/Spacepal-AI/internal/ai/list/ilist"
	"github.com/spacepal/Spacepal-AI/internal/model/imodel"
)

// ChoiceMaker is an interface for PlanetChoiceMaker
type ChoiceMaker interface {
	MakeChoice(planets []imodel.PlanetGetter, main imodel.PlanetGetter,
		factor ilist.FactorGetter) (planet imodel.PlanetGetter, distance int)
}
