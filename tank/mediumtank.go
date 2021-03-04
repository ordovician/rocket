package tank

import (
	. "github.com/ordovician/rocket/physics"
)

const (
	mediumDryMass   = 250.0
	mediumTotalMass = 2300.0
)

// A medium sized rocket tank. Similar to the capacity of
// the second stage of an Electron rocket.
type MediumTank struct {
	propellant Kg
}

func (tank *MediumTank) Propellant() Kg {
	return tank.propellant
}

// Is the tank empty?
func (tank *MediumTank) IsEmpty() bool {
	return tank.propellant == 0
}

// Consume given amount of propellant. Returns how much was consumed
func (tank *MediumTank) Consume(amount Kg) Kg {
	if tank.propellant < amount {
		amount = tank.propellant
	}
	tank.propellant -= amount
	return amount
}

// Fill tank up to the max
func (tank *MediumTank) Refill() {
	tank.propellant = mediumTotalMass - mediumDryMass
}

// Max of tank including propellant inside it.
func (tank *MediumTank) Mass() Kg {
	return tank.propellant + mediumDryMass
}

func NewMediumTank() Tank {
	tank := MediumTank{}
	tank.Refill()
	return &tank
}
