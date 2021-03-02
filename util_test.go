package rocket

import (
	. "github.com/ordovician/rocket/part"
	. "github.com/ordovician/rocket/physics"
)

func newElectronSpaceVehicle() *SpaceVehicle {
	rutherford := Rutherford{}
	boosterEngines := EngineCluster{Engine: rutherford, Count: 9}

	craft := NewSpaceCraft(nil, NewMediumTank(), rutherford)

	booster := NewMultiStaged(
		craft,
		NewLargeTank(),
		&boosterEngines)

	ship := NewSpaceVehicle(booster)
	return ship
}

func launchSpaceShip(mass Kg, thrust Newton, Isp float64) (t, elevation float64, propellant Kg) {
	engine := NewFlexiEngine(0.0, thrust, Isp)
	tank := NewFlexiTank(0.0, mass)
	probe := NewProbe(0.001)

	craft := NewSpaceCraft(probe, tank, engine)
	ship := NewSpaceVehicle(craft)

	Δt := 0.4
	t = 0.0
	for ; t <= 2000; t += Δt {
		ship.Update(Δt)

		// fmt.Printf("t: %.2f, elevation: %.2f, prop: %.2f\n", t, ship.Elevation, ship.Propellant())
		if ship.IsEmpty() {
			break
		}
	}
	return t, ship.Elevation, ship.Propellant()
}
