package vehicles

import (
	. "github.com/ordovician/rocket/parts"
	. "github.com/ordovician/rocket/physics"
)

// A space craft is what you would think of as a space ship.
// It is the part of a multistaged rocket which will be navigating
// around in outer space.
type NullRocket struct {
	*NullTank
}

// Will always be zero
func (r *NullRocket) Mass() Kg {
	return Kg(0)
}

// Does nothing, as the NullRocket is nothing
func (p *NullRocket) Update(Δt float64) {

}

// NullRocket has no stages and thus cannot be stage separated. Thus this
// will always return nil
func (r *NullRocket) StageSeparate() Rocket {
	return &NullRocket{}
}

// A NullRocket has not stages because it doesn't really exist.
func (r *NullRocket) NoStages() int {
	return 0
}

// Because a null rocket does not exist it does not produce any thrust.
// This always returns zero.
func (r *NullRocket) Thrust() Newton {
	return 0
}

// A null rocket has no specific impulse
func (r *NullRocket) Isp() float64 {
	return 0
}
