package parts

import (
	. "github.com/ordovician/rocket/physics"
	"math"
)

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
