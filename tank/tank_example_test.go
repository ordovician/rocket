package tank

import "fmt"

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

func ExampleLoadTanks() {
	tanks, _ := LoadTanks("")
	for k, v := range tanks {
		fmt.Printf("%v has mass %.1f\n", k, v.Mass())
	}

	// Unordered Output:
	// FL-T200 has mass 1130.0
	// FL-T400 has mass 2250.0
	// Falcon1-S2 has mass 4800.0
	// Electron-S1 has mass 10200.0
	// FL-T100 has mass 560.0
	// Falcon9-S1 has mass 96570.0
	// Falcon1-S1 has mass 41100.0
	// Electron-S2 has mass 2300.0
	// Electron-S3 has mass 0.0
}
