package physics

import (
	"sort"

	"testing"
	// "part"
)

func compare(test *testing.T, b *RigidBody, tmax, wanted float64) {
	Δerror := make([]float64, 0, 400)
	for _, Δt := range []float64{0.01, 0.1, 0.5, 1.0} {
		for t := 0.0; t <= tmax; t += Δt {
			b.Integrate(Δt)
		}
		Δerror = append(Δerror, b.Elevation-wanted)
	}

	// When Δt → ∞ then Δerror → ∞. Hence errors should be ascending with increasing Δt
	if !sort.Float64sAreSorted(Δerror) {
		test.Errorf("b.Integrate(Δt) larger Δt is not having larger error  %v", Δerror)
		test.Errorf("b.Mass = %v, b.Velocity = %v, b.Force = %v, b.Elevation = %v",
			b.Mass, b.Velocity, b.Force, b.Elevation)
	}
}

var bodies = []RigidBody{
	RigidBody{Mass: 1, Force: 4},
	RigidBody{Mass: 12, Force: 2},
	RigidBody{Mass: 0.01, Force: 2, Velocity: 200},
	RigidBody{Mass: 5, Force: 0.4, Elevation: -250},
}

func TestIntegrationImprovesWithSmallerDeltaT(test *testing.T) {

	for _, b := range bodies {
		distance := Distance(b.Elevation, b.Velocity, b.Acceleration())
		for _, tmax := range []float64{3, 10, 50, 200} {
			wanted := distance(tmax)
			compare(test, &b, tmax, wanted)
		}
	}
}

func TestIntegration(test *testing.T) {
	body := RigidBody{Mass: 4, Force: 2}

	body.Integrate(1)

	if body.Velocity != 0.5 {
		test.Errorf("Got %f but wanted %f", body.Velocity, 0.5)
	}

	body.Integrate(1)

	if body.Velocity != 1.0 {
		test.Errorf("Got %f but wanted %f", body.Velocity, 1.0)
	}

	body.Integrate(5)

	if body.Velocity != 3.5 {
		test.Errorf("Got %f but wanted %f", body.Velocity, 3.5)
	}

}
