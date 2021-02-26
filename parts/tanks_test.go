package parts

import "testing"

func TestMakeMediumTank(t *testing.T) {
	tank := NewMediumTank()
	if tank.Mass() != mediumTotalMass {
		t.Errorf("tank.Mass() = %f; want %f", tank.Mass(), mediumTotalMass)
	}
}

func TestMakeLargeTank(t *testing.T) {
	tank := NewLargeTank()
	if tank.Mass() != largeTotalMass {
		t.Errorf("tank.Mass() = %f; want %f", tank.Mass(), largeTotalMass)
	}
}
