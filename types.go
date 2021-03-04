package rocket

import . "github.com/ordovician/rocket/tank"
import . "github.com/ordovician/rocket/engine"
import . "github.com/ordovician/rocket/physics"

type Payload interface {
	Mass() Kg
}

// A rocket is made up of some kind payload, a propellant tank and a rocket engine.
// This is a composite structure. The payload could potentially be another rocket
// to create a staged rocket.
type Rocket interface {
	Payload
	Tank
	Engine

	Update(Δt float64)
	StageSeparate() Rocket
	NoStages() int
}
