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

func NewRigidBody(mass Kg, force Newton) *RigidBody {
	body := RigidBody{0, 0, force, mass}
	return &body
}

func (body *RigidBody) GravityForce() float64 {
	return -Gravity * float64(body.Mass)
}

func (body *RigidBody) Acceleration() float64 {
	return float64(body.Force) / float64(body.Mass)
}

func (body *RigidBody) Integrate(Δt float64) {
}
