package part

import (
	. "github.com/ordovician/rocket/physics"
)

// Useful as payload for a spacecraft or rocket.
type Probe struct {
	mass Kg
}

func (probe *Probe) Mass() Kg {
	return probe.mass
}

func NewProbe(mass Kg) *Probe {
	return &Probe{mass}
}
