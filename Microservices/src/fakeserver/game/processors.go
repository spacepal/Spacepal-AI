package game

import (
	"fakeserver/game/model"
	"math"

	log "github.com/sirupsen/logrus"
)

// processShips processes tasks
func (g *Game) processShips() {
	for id, t := range g.tasks {
		t.Steps--
		if t.Steps == 0 {
			if g.planets[t.To].OwnerID == t.Player {
				g.planets[t.To].ShipsCount += t.Count
				log.Info("Ships(", t.Count, ") of the player$",
					t.Player, " arrived to the destination planet$", t.To)
			} else {
				g.processAttack(t)
			}
			delete(g.tasks, id)
		}
	}
}

func (g *Game) processAttack(t model.Task) {
	target, source := g.planets[t.To], g.planets[t.From]
	attackPower := source.KillPercentage * float64(t.Count)
	defensePower := target.KillPercentage * float64(target.ShipsCount)
	result := math.Abs(attackPower - defensePower)
	killPerc := g.planets[t.To].KillPercentage
	if attackPower < defensePower {
		log.Info("Player$%d ships (%d) lost on the planet$%d",
			t.Player, t.Count, t.To)
	} else {
		log.Info("The player$", t.Player, " captured planet by %d ship(s)$",
			t.To, t.Count)
		killPerc = g.planets[t.From].KillPercentage
		g.planets[t.To].OwnerID = t.Player
		g.planets[t.To].Production = g.planets[t.To].InitialProduction
	}
	g.planets[t.To].ShipsCount = int(math.Ceil(result / killPerc))
}

func (g *Game) processPlanets() {
	for id, p := range g.planets {
		if !p.IsNeutral() {
			g.planets[id].IncProduction()
		}
		g.planets[id].ShipsCount += p.Production
	}
}
