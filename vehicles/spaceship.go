package vehicles

import . "github.com/ordovician/rockets/physics"
import . "github.com/ordovician/rockets/parts"

type SpaceVehicle struct {
	ActiveStage StagedRocket
	RigidBody
}

// Update propellant and elevantion of space vehicle
func (ship *SpaceVehicle) Update(Δt float64) {
	stage := ship.ActiveStage
	stage.Update(Δt) // consume fuel. Affects mass
	ship.RigidBody.Mass = stage.Mass()
	ship.RigidBody.Force = stage.Thrust()
	ship.RigidBody.Integrate(Δt)
}

func NewSpaceVehicle(payload Part, tank Tank, engine Engine) *SpaceVehicle {
	rocket := NewSingleRocket(tank, engine)
	rocket.Payload = payload

	stage := StagedRocket{nil, *rocket}
	rigidBody := NewRigidBody(stage.Mass(), 0)
	ship := SpaceVehicle{stage, *rigidBody}
	return &ship
}

func NewStagedRocket(stages []Rocket) *SpaceVehicle {

	// for i, stage := range stages {
	//
	// }
	return new(SpaceVehicle)
}

func (ship *SpaceVehicle) Launch(Δt, max_duration float64) {
	stage := &ship.ActiveStage
	for t := 0.0; t < max_duration; t += Δt {
		if stage == nil {
			return
		}

		stage.Update(Δt)

		if stage.IsEmpty() {
			stage = ship.StageSeparate()
		}
	}
}

func (ship *SpaceVehicle) StageSeparate() *StagedRocket {
	return nil
}
