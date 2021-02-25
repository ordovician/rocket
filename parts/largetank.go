package parts

import . "github.com/ordovician/rockets/physics"

const (
	largeDryMass   = 950
	largeTotalMass = 10200
)

// A large rocket tank
type LargeTank Kg

func (tank *LargeTank) IsEmpty() bool {
	return *tank == largeDryMass
}

func (tank *LargeTank) Consume(amount Kg) Kg {
	*tank -= LargeTank(amount)
	return Kg(*tank)
}

func (tank *LargeTank) Refill() {
	*tank = LargeTank(largeTotalMass - largeDryMass)
}

// Current mass of tank taking into account amount of propellant consumed
func (tank *LargeTank) Mass() Kg {
	return Kg(*tank) + largeDryMass
}

func NewLargeTank() Tank {
	tank := LargeTank(0)
	tank.Refill()
	return &tank
}
