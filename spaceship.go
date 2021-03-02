package rocket

import . "github.com/ordovician/rocket/physics"

// SpaceVehicle is a multi-staged rocket which can be launched from planetary surface and into orbit.
// It is potentially made up of many Rocket parts. A rocket can contain a rocket as a payload.
// The Rocket type alone only deals with mass and thrust. A SpaceVehicle can keep track of latitude.
// You need a SpaceVehicle to simulate a rocket launch, as a Rocket alone does not contain physics calculations.
type SpaceVehicle struct {
	Rocket
	RigidBody
}

// NewSpaceVehicle creates a space vehicle from a single or multi-staged rocket.
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

// Mass of the whole space vehicle measured in Kg.
func (ship *SpaceVehicle) Mass() Kg {
	return ship.Rocket.Mass()
}

// StageSeparate the space vehicle. This means dumping the bottom stage of
// a multi-stage rocket. What was the payload of the previous bottom stage becomes the
// new bottom stage. This action means the mass of the whole ship is reduced
// and it gets a new value for thrust derived from whatever engine is attached to the
// new active stage.
func (ship *SpaceVehicle) StageSeparate() Rocket {
	stage := ship.Rocket.StageSeparate()
	ship.Rocket = stage
	return stage
}

// Launch a space vehicle into space. The launch is simulated with time incremenets of Δt.
// Simulation lasts maximum max_duration. Smaller values for Δt gives more accurate simulation but
// requires more computations. In the simulation one stage will consume its fuel
// before being detached and the next stage take over
//
// Returns the time in seconds when the all the fuel ran out or we reached max duration along
// with the elevation achieved at that time
func (ship *SpaceVehicle) Launch(Δt, max_duration float64, monitor LaunchMonitor) (t, elevation float64) {
	t = 0.0
	for ; t <= max_duration; t += Δt {
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
	return t, ship.RigidBody.Elevation
}
