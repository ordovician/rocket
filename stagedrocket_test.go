package rocket

import (
	"testing"

	"math"

	// . "github.com/ordovician/rocket/engine"
	. "github.com/ordovician/rocket/physics"
	. "github.com/ordovician/rocket/tank"
)

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

// If thrust is too low we cannot fight gravity and we don't managed to get off ground
func TestThrustTooLowForGravity(tst *testing.T) {
	var mass Kg = 200
	var Isp float64 = 50

	// The whole weight of this space ship is in the propellant. That means downward force is -mass*Gravity
	// If we set thurst of Rocet to +0.5*mass*Gravity then we will get elevation around 0. Why?
	// Because initially gravity will be stronger and pull ship down, but eventually half the weight is burned up, and then the
	// thrust is stronger than gravity and thus half the tank is spent getting back to the starting point
	_, elevation, _ := launchSpaceShip(mass, Newton(float64(mass)*Gravity/2), Isp)

	tolerance := 5.0
	if diff := math.Abs(elevation); diff > tolerance {
		tst.Errorf("Got elevation = %.3f, want elevation ± %.3f", elevation, tolerance)
	}
}

// Show that get e.g. at least 900 meter above ground as long as Force is slightly above
// half the force of gravity.
func TestPlentyPowerful(tst *testing.T) {
	var mass Kg = 200
	var Isp float64 = 1000
	_, elevation, _ := launchSpaceShip(mass, Newton(float64(mass)*Gravity/2)+0.01, Isp)

	threshold := 900.0
	if elevation < threshold {
		tst.Errorf("Got elevation = %.3f, want elevation > %.3f", elevation, threshold)
	}
}

// We can calculate distance with Newton equations. In these mass will not change which means you should expect shorter elevation
// achieved with the same starting mass.
func TestCompareWithNewtonMotionEquations(tst *testing.T) {
	var mass Kg = 1331
	var Isp float64 = 42
	thrust := Newton(float64(mass) * Gravity)

	t, elevation, _ := launchSpaceShip(mass, thrust, Isp)

	gravityForce := Newton(float64(mass) * Gravity)
	force := thrust - gravityForce
	accel := float64(force) / float64(mass)

	distance := Distance(0, 0, accel)

	if distance(t) > elevation {
		tst.Errorf("Got distance(%.2f) <= %.2f  Want distance(%.2f) > %.2f", t, elevation, t, elevation)
	}
}
