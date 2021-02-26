package physics

// An implementation of the following equation:
//     v = at + v0
// It is made so it returns a function of time:
//     v = Velocity(v0, a)  # Returns a function r(t)
//     dist = v(t)
func Velocity(v0, a float64) func(t float64) float64 {
	return func(t float64) float64 {
		return a*t + v0
	}
}

// An implementation of the following equation:
//     r = r0 + v0*t + 0.5a*t^2
// It is made so it returns a function of time:
//     r = Distance(r0, v0, a)  # Returns a function r(t)
//     dist = r(t)
func Distance(r0, v0, a float64) func(t float64) float64 {
	return func(t float64) float64 {
		return r0 + v0*t + 0.5*a*t*t
	}
}
