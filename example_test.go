package rocket

import (
	"fmt"

	. "github.com/ordovician/rocket/engine"
	. "github.com/ordovician/rocket/tank"
	// . "github.com/ordovician/rocket/physics"
)

// Shows how total mass and thrust will change when you do stage separation.
// This happens because you knock off the last part. What used to be the payload of
// the bottom part thus becomes the remaining part, and its engine provides the new thrust.
func ExampleMultiStaged_StageSeparate() {
	rutherford := Rutherford{}
	boosterEngines := Cluster{rutherford, 9}

	booster := NewMultiStaged(
		NewSpaceCraft(nil, NewMediumTank(), rutherford),
		NewLargeTank(),
		&boosterEngines)

	fmt.Println("Mass: ", booster.Mass())
	fmt.Println("Thrust: ", booster.Thrust())

	booster.StageSeparate()
	fmt.Println("Mass: ", booster.Mass())
	fmt.Println("Thrust: ", booster.Thrust())

	// Output:
	// Mass:  12850
	// Thrust:  225000
	// Mass:  10515
	// Thrust:  225000
}

// When launching a rocket you can provide a callback which can print out state of space vehicle
// each time a stage separation has happened. E.g. how many seconds have passed, how far up it has gotten
// and what the total mass is, given that lower stages have been removed.
//
// This show a two stage rocket which has a single separation. We record mass and elevation of rocket above
// ground at each stage.
func ExampleSpaceVehicle_Launch() {
	rutherford := Rutherford{}
	boosterEngines := Cluster{rutherford, 9}

	booster := NewMultiStaged(
		NewSpaceCraft(nil, NewMediumTank(), rutherford),
		NewLargeTank(),
		&boosterEngines)

	ship := NewSpaceVehicle(booster)
	fmt.Printf("Before launch, ship mass =  %.1f Kg\n", ship.Mass())
	fmt.Println("  No. stages = ", ship.NoStages())
	fmt.Println("  Booster propellant = ", ship.Propellant(), "Kg")

	ship.Launch(0.1, 500, func(t float64) {
		fmt.Printf("Stage separation at t = %.1f s\n", t)
		fmt.Printf("  elevation = %.1f km\n", ship.Elevation/1e3)
		fmt.Printf("  ship mass = %.1f Kg\n", ship.Mass())
		fmt.Printf("  propellant = %.1f Kg\n", ship.Propellant())
		fmt.Println("  stages left = ", ship.NoStages())
	})

	fmt.Printf("After launch, ship mass =  %.1f Kg\n", ship.Mass())
	fmt.Printf("  elevation = %.1f km\n", ship.Elevation/1e3)
	fmt.Println("  Stages left = ", ship.NoStages())
	fmt.Println("  Propellant left = ", ship.Propellant(), "Kg")

	// Output:
	// Before launch, ship mass =  12850.0 Kg
	//   No. stages =  2
	//   Booster propellant =  9250 Kg
	// Stage separation at t = 125.3 s
	//   elevation = 115.9 km
	//   ship mass = 2335.0 Kg
	//   propellant = 2050.0 Kg
	//   stages left =  1
	// After launch, ship mass =  285.0 Kg
	//   elevation = 1012.8 km
	//   Stages left =  1
	//   Propellant left =  0 Kg
}

func ExampleTankPropellant() {
	rutherford := Rutherford{}
	boosterEngines := Cluster{rutherford, 9}

	craft := NewSpaceCraft(nil, NewMediumTank(), rutherford)

	booster := NewMultiStaged(
		craft,
		NewLargeTank(),
		&boosterEngines)

	ship := NewSpaceVehicle(booster)

	mediumTank := craft.Tank
	largeTank := booster.Tank

	fmt.Println("Ship tank propellant: ", ship.Propellant())
	fmt.Println("Large tank propellant: ", largeTank.Propellant())
	fmt.Println("Medium tank propellant: ", mediumTank.Propellant())

	fmt.Println("Stage separation")
	ship.StageSeparate()

	fmt.Println("Ship tank propellant: ", ship.Propellant())

	// Output:
	// Ship tank propellant:  9250
	// Large tank propellant:  9250
	// Medium tank propellant:  2050
	// Stage separation
	// Ship tank propellant:  2050
}

func ExampleSpaceVehicle_Update() {

	for _, Isp := range []float64{1, 10, 100} {
		t, elevation, propellant := launchSpaceShip(200, 1000, Isp)
		fmt.Printf("Isp: %.2f, t: %.2f s, elevation: %.2f km, prop: %.2f Kg\n", Isp, t, elevation/1e3, propellant)
	}

	// Output:
	// Isp: 1.00, t: 1.60 s, elevation: 0.00 km, prop: 0.00 Kg
	// Isp: 10.00, t: 19.60 s, elevation: 1.30 km, prop: 0.00 Kg
	// Isp: 100.00, t: 196.00 s, elevation: 5.89 km, prop: 0.00 Kg
}
