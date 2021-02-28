package part

import (
	. "github.com/ordovician/rocket/physics"
	"math"
)

// Rockets such as Falcon 9 and Electron have clusters of multiple engines
// at the bottom booster stage. This allows us to create an "engine" made up of multiple engines.
type EngineCluster struct {
	Engine
	Count uint8
}

func (cluster *EngineCluster) Mass() Kg {
	multiplier := math.Max(float64(cluster.Count), 1.0)
	return cluster.Engine.Mass() * Kg(multiplier)
}

func (cluster *EngineCluster) Thrust() Newton {
	count := float64(cluster.Count)
	multiplier := math.Max(count, 1.0)
	return cluster.Engine.Thrust() * Newton(multiplier)
}
