package rocket

import (
	. "github.com/ordovician/rocket/parts"
	. "github.com/ordovician/rocket/physics"
)

// The propulsion element of a rocket. The Engine decides how much propellant
// is consumed by the tank each second
type Propulsion struct {
	Tank
	Engine
}

// The mass of the propulsion element. This is not fixed but will fall as
// propellant in the tank is consumed.
func (p *Propulsion) Mass() Kg {
	return p.Tank.Mass() + p.Engine.Mass()
}

// Advance simulation of one time unit, reducing content of propellant tank
// based on mass flow out of rocket engine
func (p *Propulsion) Update(Δt float64) {
	mflow := MassFlow(p.Thrust(), p.Isp())
	p.Consume(Kg(mflow * Δt))
}
