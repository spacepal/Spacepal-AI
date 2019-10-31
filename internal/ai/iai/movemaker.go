package iai

import (
	"github.com/spacepal/Spacepal-AI/internal/helpers/ihelpers"
	"github.com/spacepal/Spacepal-AI/internal/model/imodel"
)

// MoveMaker is an interface for artificial intelligence end turn
type MoveMaker interface {
	MakeMove(
		planets ihelpers.PlanetsGetter,
		globStat ihelpers.GlobStatGetter,
		mapSize imodel.MapSizeGetter,
	) []imodel.TaskGetter
}
