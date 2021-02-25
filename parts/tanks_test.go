package parts

import "testing"

func TestMakeTank(t *testing.T) {
	tank := NewMediumTank()
	if tank.Mass() != mediumTotalMass {
		t.Errorf("tank.Mass() = %f; want %f", tank.Mass(), mediumTotalMass)
	}
}
