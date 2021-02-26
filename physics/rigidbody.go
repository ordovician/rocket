package physics

// Represents an inelastic object used
// to simulate simple movement cause
// by forces
type RigidBody struct {
	Elevation float64
	Velocity  float64
	Force     Newton
	Mass      Kg
}

// The force of gravity on the rigid body.
func (body *RigidBody) GravityForce() Newton {
	return Newton(-Gravity * float64(body.Mass))
}

// Acceleration of of this body upwards.
// Gravity is defined as causing downward acceleration.
func (body *RigidBody) Acceleration() float64 {
	return float64(body.Force) / float64(body.Mass)
}

// Perform one integration step using the explict euler method. This will update
// the position and velocity of the body by advancing the time with Δt.
func (body *RigidBody) Integrate(Δt float64) {
	body.Elevation += body.Velocity * Δt
	body.Velocity += body.Acceleration() * Δt
}
