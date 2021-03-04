package tank

import . "github.com/ordovician/rocket/physics"

// A tank to hold rocket propellant
type Tank interface {
	Mass() Kg
	Propellant() Kg
	IsEmpty() bool
	Consume(amount Kg) Kg
	Refill()
}
