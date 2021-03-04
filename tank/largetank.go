package tank

import . "github.com/ordovician/rocket/physics"

const (
	largeDryMass   = 950.0
	largeTotalMass = 10200.0
)

// A large rocket tank
type LargeTank Kg

func (tank *LargeTank) Propellant() Kg {
	return Kg(*tank)
}

func (tank *LargeTank) IsEmpty() bool {
	return *tank == 0
}

func (tank *LargeTank) Consume(amount Kg) Kg {
	if Kg(*tank) < amount {
		amount = Kg(*tank)
	}
	*tank -= LargeTank(amount)

	return amount
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
