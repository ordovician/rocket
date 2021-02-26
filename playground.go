package main

import (
	"fmt"
	"rockets/physics"
)

func main() {
	body := physics.RigidBody{Mass: 4, Force: 2}
	body.Force = 4
	body.Integrate(4)
	body.Integrate(2)
	body.Integrate(2)
	fmt.Println("Velocity: ", body.Velocity)
	fmt.Println("Elevation: ", body.Elevation)
}
