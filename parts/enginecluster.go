package parts

import . "github.com/ordovician/rockets/physics"

type EngineCluster struct {
	Engine
	count uint8
}

func (cluster *EngineCluster) Mass() Kg {
	return cluster.Engine.Mass() * Kg(cluster.count)
}

func (cluster *EngineCluster) Thrust() Newton {
	return cluster.Engine.Thrust() * Newton(cluster.count)
}

func NewEngineCluster(engine Engine, count uint8) Engine {
	cluster := EngineCluster{engine, count}
	return &cluster
}
