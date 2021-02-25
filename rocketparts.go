package main

import (
	"fmt"

	"github.com/ordovician/rockets/parts"
	"github.com/ordovician/rockets/vehicles"
)

func main() {
	tank := parts.NewMediumTank()
	fmt.Printf("Tank size: %v\n", tank.Mass())

	engine := parts.NewRutherford()
	cluster := parts.NewEngineCluster(engine, 2)

	fmt.Printf("Single engine thrust: %v\nCluster thrust: %v\n", engine.Thrust(), cluster.Thrust())
	fmt.Printf("Single engine mass: %v\nCluster mass: %v\n", engine.Mass(), cluster.Mass())

	rocket := vehicles.NewSingleRocket(tank, cluster)
	rocket.Consume(100)
	fmt.Printf("Rocket Mass: %v\n", rocket.Mass())
	rocket.Refill()
	fmt.Printf("Rocket Mass after Refill: %v\n", rocket.Mass())

	fmt.Printf("Rocket thrust: %v\n", rocket.Thrust())
}
