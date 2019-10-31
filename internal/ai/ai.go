package ai

import (
	"github.com/spacepal/Spacepal-AI/internal/model"
	"github.com/spacepal/Spacepal-AI/internal/model/imodel"
)

func safeCreateTask(
	main imodel.PlanetGetter,
	target imodel.PlanetGetter,
	quantity float64,
	tasks *[]imodel.TaskGetter,
	distance int,
) {
	if target == nil {
		return
	}
	if target.ID() == main.ID() {
		return
	}
	shipsCount := float64(main.Ships()) * quantity
	if shipsCount == 0 {
		return
	}
	task := model.NewTask(main.Owner(),
		main.ID(), target.ID(),
		int(shipsCount), distance)
	*tasks = append(*tasks, task)
}
