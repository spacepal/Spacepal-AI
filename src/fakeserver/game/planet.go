package game

import (
	"fmt"
	"math"
	"math/rand"

	log "github.com/sirupsen/logrus"
)

const neutralPlanetID = -1
const killDev = 0.1     // standard deviation,
const killMean = 0.5    // expectation of the distribution
const prodDev = 4       // standard deviation,
const prodMean = 10     // expectation of the distribution
const prodMin = 1       // minimum value of production
const capitalProd = 10  // default production of capital
const capitalKill = 0.4 // default kill percentage of capital

// Planet is a model of game planet on map-grid
type planet struct {
	PlanetID          int     `json:"id"`
	OwnerID           int     `json:"ownerID"`
	CellID            int     `json:"cellID"`
	Production        int     `json:"production"`
	InitialProduction int     `json:"-"`
	KillPercentage    float64 `json:"killPerc"`
	ShipsCount        int     `json:"ships"`
}

func newNeutralPlanet(planetID, cellID int) planet {
	var p planet
	p.setRndKillPerc()
	p.setRndProduction()
	p.InitialProduction = p.Production
	p.OwnerID = neutralPlanetID
	p.CellID = cellID
	p.PlanetID = planetID
	p.ShipsCount = capitalProd
	return p
}

func (p planet) ID() int         { return p.PlanetID }
func (p planet) Owner() int      { return p.OwnerID }
func (p planet) Cell() int       { return p.CellID }
func (p planet) Prod() int       { return p.Production }
func (p planet) InitalProd() int { return p.InitialProduction }
func (p planet) Kill() float64   { return p.KillPercentage }
func (p planet) Ships() int      { return p.ShipsCount }
func (p planet) IsNeutral() bool { return p.OwnerID == neutralPlanetID }
func (p planet) String() string {
	player := "neutral"
	if !p.IsNeutral() {
		player = fmt.Sprint("player$", p.OwnerID)
	}
	return fmt.Sprintf("{Planet#%d %s cell: %d kill: %f prod: %d ships: %d}",
		p.PlanetID, player, p.CellID, p.KillPercentage, p.Production, p.ShipsCount)
}

// setCapital sets capital status for planet
func (p *planet) setCapital(owner int) {
	p.InitialProduction = capitalProd
	p.Production = p.InitialProduction
	p.KillPercentage = capitalKill
	p.OwnerID = owner
	p.ShipsCount = p.Production
}

// DecreaseShipsCount reduces a count of ships by value
func (p *planet) DecreaseShipsCount(byValue int) {
	p.ShipsCount -= byValue
	if p.ShipsCount < 0 {
		log.Errorf("planet.go DecreaseShipsCount(%d): Ships count is negative!",
			byValue)
		p.ShipsCount = 0
	}
}

// IncProduction increases production
// it used for cumulative production game mode
func (p *planet) IncProduction() {
	p.Production++
}

// ChangeOwner changes owner of planet
func (p *planet) ChangeOwner(ownerID int) {
	if p.OwnerID == neutralPlanetID {
		p.setCapital(ownerID)
	} else {
		p.Production = p.InitialProduction
		p.OwnerID = ownerID
	}
}

func (p *planet) setRndKillPerc() {
	var kill = rand.NormFloat64()*killDev + killMean
	kill = math.Min(math.Max(kill, 0), 1)
	p.KillPercentage = kill
}

func (p *planet) setRndProduction() {
	var prod = rand.NormFloat64()*prodDev + prodMean
	prod = math.Max(prod, prodMin)
	p.Production = int(prod)
}

// genPlanets generates a slice of random neutral planets
func (g *Game) genPlanets() {
	cellIDs := rand.Perm(g.params.MapWidth*g.params.MapHeight - 1)
	planets := make([]planet, g.params.PlanetsCount)
	for i := range planets {
		planets[i] = newNeutralPlanet(i, cellIDs[i]+1)
	}
	g.planets = planets
}
