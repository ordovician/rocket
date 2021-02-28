package rocket

import . "github.com/ordovician/rocket/physics"

type SpaceVehicle struct {
	Rocket
	RigidBody
}

func NewSpaceVehicle(rocket Rocket) *SpaceVehicle {
	rigidBody := RigidBody{Mass: rocket.Mass(), Force: 0}
	ship := SpaceVehicle{rocket, rigidBody}
	return &ship
}

// Update propellant and elevantion of space vehicle
func (ship *SpaceVehicle) Update(Δt float64) {
	stage := ship.Rocket
	stage.Update(Δt) // consume fuel. Affects mass
	body := &ship.RigidBody
	body.Mass = stage.Mass()
	body.Force = stage.Thrust()
	body.Force -= body.GravityForce()
	body.Integrate(Δt)
}

func (ship *SpaceVehicle) Mass() Kg {
	return ship.Rocket.Mass()
}

// Dump the bottom stage of the multi-stage rocket
// What was the payload of the previous bottom stage becomes the
// new bottom stage. This action means the mass of the whole ship is reduced
// and it gets a new value for thrust derived from whatever engine is attached to the
// new active stage.
func (ship *SpaceVehicle) StageSeparate() Rocket {
	stage := ship.Rocket.StageSeparate()
	ship.Rocket = stage
	return stage
}

// Simulate a launch of a space vehicle (a multi-stage rocket). The simulation
// is performed in time steps which are Δt long each. Simulation lasts maximum max_duration.
// In the simulation one stage will consume its fuel before being detached and the next stage take over
// Returns the elevation reached by rocket.
func (ship *SpaceVehicle) Launch(Δt, max_duration float64, monitor LaunchMonitor) float64 {

	for t := 0.0; t <= max_duration; t += Δt {
		ship.Update(Δt)

		if ship.Rocket.IsEmpty() {
			if ship.NoStages() == 1 {
				break
			}

			stage := ship.StageSeparate()

			// tell the world that a stage got separated and when it happened
			if monitor != nil {
				monitor(t)
			}

			// no more stages left, so we cannot keep flying
			if stage == nil {
				break
			}
		}
	}
	return ship.RigidBody.Elevation
}
