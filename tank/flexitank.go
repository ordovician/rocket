package tank

import (
	. "github.com/ordovician/rocket/physics"
)

// A tank with flexible size
type FlexiTank struct {
	DryMass    Kg
	TotalMass  Kg
	propellant Kg
}

func (tank *FlexiTank) Propellant() Kg {
	return tank.propellant
}

// Is the tank empty?
func (tank *FlexiTank) IsEmpty() bool {
	return tank.propellant == 0
}

// Consume given amount of propellant. Returns how much was consumed
func (tank *FlexiTank) Consume(amount Kg) Kg {
	if tank.propellant < amount {
		amount = tank.propellant
	}
	tank.propellant -= amount
	return amount
}

// Fill tank up to the max
func (tank *FlexiTank) Refill() {
	tank.propellant = tank.TotalMass - tank.DryMass
}

// Max of tank including propellant inside it.
func (tank *FlexiTank) Mass() Kg {
	return tank.propellant + tank.DryMass
}

// Creates a tank with given dry mass and total mass, and fills it up to the max.
// Dry mass is the mass of an empty tank.
func NewFlexiTank(drymass, totalmass Kg) *FlexiTank {
	tank := FlexiTank{drymass, totalmass, 0}
	tank.Refill()
	return &tank
}
