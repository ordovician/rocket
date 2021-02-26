package physics

import (
	. "fmt"
)

func ExampleDistance() {
	distance := Distance(0, 0, 3)
	d := distance(2)

	Println("Distances for distance = Distance(0, 0, 3)")
	Println("  distance(2) =", d)
	Println("  distance(0) =", distance(0))
	Println("  distance(4) =", distance(4))

	distance = Distance(10, 0, 3)
	Println("Distances for distance = Distance(10, 0, 3)")
	Println("  distance(0) =", distance(2))
	Println("  distance(0) =", distance(0))
	Println("  distance(4) =", distance(4))

	// Output:
	// Distances for distance = Distance(0, 0, 3)
	//   distance(2) = 6
	//   distance(0) = 0
	//   distance(4) = 24
	// Distances for distance = Distance(10, 0, 3)
	//   distance(0) = 16
	//   distance(0) = 10
	//   distance(4) = 34
}

func ExampleRigidBody_Integrate() {
	b := RigidBody{Mass: 1, Force: 0, Velocity: 1}

	Println("Zero Force")
	b.Integrate(0)
	Println(b.Elevation)

	b.Integrate(2)
	Println(b.Elevation)

	b.Integrate(3)
	Println(b.Elevation)

	b.Integrate(5)
	Println(b.Elevation)

	Println("Force is 2 Newton on 1 Kg object")
	b = RigidBody{Mass: 1, Force: 2}

	const Δt = 0.1
	for t := 0.0; t <= 12; t += Δt {
		b.Integrate(Δt)
	}
	accel := b.Acceleration()
	distance := Distance(0, 0, accel)

	Printf("Integration: %.2f\n", b.Elevation)
	Printf("Newton: %.2f\n", distance(12))

	// Output:
	// Zero Force
	// 0
	// 2
	// 5
	// 10
	// Force is 2 Newton on 1 Kg object
	// Integration: 145.20
	// Newton: 144.00
}
