package physics

type Newton float64
type Kg float64

const (
	Gravity = 9.80665 // Acceleration of Gravity on Earth in meters / seconds squared
)

// Calculates the exhaust velocity `vₑ` of a rocket engine, by taking specific impulse as input.
func ExhaustVelocity(Isp float64) float64 {
	return Isp * Gravity
}

// Get mass flow in Kg/s of propellant for an engine with given thrust and specific impulse (Isp)
// Use to figure out how fast propellant tanks get depleted
func MassFlow(thrust Newton, Isp float64) float64 {
	return float64(thrust) / (Isp * Gravity)
}
