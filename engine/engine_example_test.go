package engine

import "fmt"

// Creating a cluster of nine Rutherford engines. Shows how e.g.
// thrust gets combined but specific impulse (Isp) stays the same.
func ExampleCluster() {
	var engine Engine = &Rutherford{}
	var cluster Engine = &Cluster{engine, 9}

	fmt.Println(engine.Mass(), cluster.Mass())
	fmt.Println(engine.Thrust(), cluster.Thrust())
	fmt.Println(engine.Isp(), cluster.Isp())

	// Output:
	// 35 315
	// 25000 225000
	// 311 311
}

func ExampleMerlin() {
	var engine Engine = &Merlin{}
	fmt.Println(engine.Mass())
	fmt.Println(engine.Thrust())

	// Output:
	// 470
	// 845000
}

func ExampleLoadEngines() {
	engines, _ := LoadEngines()
	for name, engine := range engines {
		fmt.Printf("%v has thrust %.1f and Isp %.1f\n", name, engine.Thrust(), engine.Isp())
	}

	// Unordered Output:
	// BE-3U has thrust 0.0 and Isp 0.0
	// 48-7S has thrust 16.9 and Isp 270.0
	// RD-180 has thrust 3830.0 and Isp 311.0
	// Rutherford has thrust 25.0 and Isp 311.0
	// Kestrel 2 has thrust 0.0 and Isp 0.0
	// RS-25 has thrust 1860.0 and Isp 366.0
	// BE-4 has thrust 2400.0 and Isp 0.0
	// BE-3PM  has thrust 490.0 and Isp 310.0
	// LV-T30 has thrust 205.2 and Isp 265.0
	// LV-T45  has thrust 168.0 and Isp 250.0
	// Merlin 1D has thrust 845.0 and Isp 282.0
}
