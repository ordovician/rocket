package vehicles

import (
	. "github.com/ordovician/rockets/parts"
)

type Rocket struct {
	Tank
	Engine
}

func (rocket *Rocket) Mass() Kg {
	return rocket.Tank.Mass() + rocket.Engine.Mass()
}
