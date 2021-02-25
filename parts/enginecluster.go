package parts

type EngineCluster struct {
	engine Engine
	count  uint8
}

func (cluster *EngineCluster) Mass() Kg {
	return cluster.engine.Mass() * Kg(cluster.count)
}

func (cluster *EngineCluster) Thrust() Newton {
	return cluster.engine.Thrust() * Newton(cluster.count)
}

func NewEngineCluster(engine Engine, count uint8) Engine {
	cluster := EngineCluster{engine, count}
	return &cluster
}
