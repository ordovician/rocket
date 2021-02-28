package rocket

import (
	. "github.com/ordovician/rocket/part"
	"testing"
)

func newElectronSpaceVehicle() *SpaceVehicle {
	rutherford := Rutherford{}
	boosterEngines := EngineCluster{rutherford, 9}

	craft := NewSpaceCraft(nil, NewMediumTank(), rutherford)

	booster := NewMultiStaged(
		craft,
		NewLargeTank(),
		&boosterEngines)

	ship := NewSpaceVehicle(booster)
	return ship
}

func TestStageSeparation(tst *testing.T) {
	ship := newElectronSpaceVehicle()

	largeTank := NewLargeTank()
	mediumTank := NewMediumTank()

	if ship.Propellant() != largeTank.Propellant() {
		tst.Errorf("Got %f but wanted %f", ship.Propellant(), largeTank.Propellant())
	}

	ship.StageSeparate()

	if ship.Propellant() != mediumTank.Propellant() {
		tst.Errorf("Got %f but wanted %f", ship.Propellant(), mediumTank.Propellant())
	}
}
