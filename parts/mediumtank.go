package parts

import . "github.com/ordovician/rockets/physics"

const (
	mediumDryMass   = 250
	mediumTotalMass = 2300
)

// A medium sized rocket tank
type MediumTank struct {
	propellant Kg
}

func (tank *MediumTank) IsEmpty() bool {
	return tank.propellant == mediumDryMass
}

func (tank *MediumTank) Consume(amount Kg) Kg {
	tank.propellant -= amount
	return tank.propellant
}

func (tank *MediumTank) Refill() {
	tank.propellant = mediumTotalMass - mediumDryMass
}

func (tank *MediumTank) Mass() Kg {
	return tank.propellant + mediumDryMass
}

func NewMediumTank() Tank {
	tank := MediumTank{}
	tank.Refill()
	return &tank
}
