package main

import (
	"fmt"

	"github.com/ordovician/rockets/parts"
)

func main() {
	tank := parts.NewMediumTank()
	fmt.Printf("Tank size: %v\n", tank.Mass())

	engine := parts.NewRutherford()
	cluster := parts.NewEngineCluster(engine, 2)

	fmt.Printf("Single engine thrust: %v\nCluster thrust: %v\n", engine.Thrust(), cluster.Thrust())
	fmt.Printf("Single engine mass: %v\nCluster mass: %v\n", engine.Mass(), cluster.Mass())

}
