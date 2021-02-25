package parts

type Newton float64
type Kg float64

type Body interface {
	Mass() Kg
}

// A tank to hold rocket propellant
type Tank interface {
	Body
	IsEmpty() bool
	Consume(amount Kg) Kg
	Refill()
}

type Engine interface {
	Body
	Thrust() Newton
}
