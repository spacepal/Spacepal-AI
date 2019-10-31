package ihelpers

import "github.com/spacepal/Spacepal-AI/internal/model/imodel"

// PlanetsGetter has methods for planets listing by some conditions
type PlanetsGetter interface {
	Self() []imodel.PlanetGetter
	Foreign() []imodel.PlanetGetter
	SelfGroups(int) map[int][]imodel.PlanetGetter
}
