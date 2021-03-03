package part

import "fmt"

// Creating a cluster of nine Rutherford engines. Shows how e.g.
// thrust gets combined but specific impulse (Isp) stays the same.
func ExampleEngineCluster() {
	var engine Engine = &Rutherford{}
	var cluster Engine = &EngineCluster{engine, 9}

	fmt.Println(engine.Mass(), cluster.Mass())
	fmt.Println(engine.Thrust(), cluster.Thrust())
	fmt.Println(engine.Isp(), cluster.Isp())

	// Output:
	// 35 315
	// 25000 225000
	// 311 311
}

func ExampleMediumTank_Consume() {
	var tank Tank = NewMediumTank()
	fmt.Println(tank.Propellant())

	tank.Consume(200)
	fmt.Println(tank.Propellant())

	// Output:
	// 2050
	// 1850
}

func ExampleMediumTank_IsEmpty() {
	var tank Tank = NewMediumTank()
	fmt.Println(tank.IsEmpty())

	tank.Consume(tank.Propellant())
	fmt.Println(tank.IsEmpty())

	// Output:
	// false
	// true
}

func ExampleLargeTank_Refill() {
	var tank Tank = NewLargeTank()
	fmt.Println(tank.Propellant())

	tank.Consume(tank.Propellant() / 2)
	fmt.Println(tank.Propellant())

	tank.Refill()
	fmt.Println(tank.Propellant())

	// Output:
	// 9250
	// 4625
	// 9250
}

func ExampleNullTank() {
	var tank Tank = &NullTank{}
	fmt.Println(tank.Propellant())
	fmt.Println(tank.IsEmpty())

	tank.Refill()
	fmt.Println(tank.Propellant())
	fmt.Println(tank.IsEmpty())
	fmt.Println(tank.Mass())

	// Output:
	// 0
	// true
	// 0
	// true
	// 0
}
