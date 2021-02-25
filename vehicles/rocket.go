package vehicles

import (
	. "github.com/ordovician/rockets/parts"
	. "github.com/ordovician/rockets/physics"
)

type Rocket interface {
	Part
	Tank
	Engine
}

type SingleRocket struct {
	Payload Part
	Tank
	Engine
}

func NewSingleRocket(tank Tank, engine Engine) *SingleRocket {
	rocket := SingleRocket{nil, tank, engine}
	return &rocket
}

func (rocket *SingleRocket) Mass() Kg {
	mass := rocket.Tank.Mass() + rocket.Engine.Mass()
	if rocket.Payload != nil {
		mass += rocket.Payload.Mass()
	}
	return mass
}

func (rocket *SingleRocket) Update(Δt float64) {
	mflow := MassFlow(rocket.Thrust(), rocket.Isp())
	rocket.Consume(Kg(mflow * Δt))
}

type StagedRocket struct {
	payload Rocket
	SingleRocket
}

func (stage *StagedRocket) Mass() Kg {
	mass := stage.SingleRocket.Mass()
	if stage.payload != nil {
		mass += stage.payload.Mass()
	}
	return mass
}
