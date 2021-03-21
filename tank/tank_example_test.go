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
	// FL-T400 has mass 2.2
	// F9 2nd stage has mass 96.6
	// F1 2nd stage has mass 4.8
	// FL-T100 has mass 0.6
	// FL-T200 has mass 1.1
	// Electron 1st stage has mass 10.2
	// Electron 2nd stage has mass 2.3
	// Electron 3rd stage has mass 0.0
	// F9 1st stage has mass 418.8
	// F1 1st stage has mass 41.1
}
