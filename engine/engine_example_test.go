package engine

import "fmt"

// Creating a cluster of nine Rutherford engines. Shows how e.g.
// thrust gets combined but specific impulse (Isp) stays the same.
func ExampleCluster() {
	var engine Engine = &Rutherford{}
	var cluster Engine = &engine.Cluster{engine, 9}

	fmt.Println(engine.Mass(), cluster.Mass())
	fmt.Println(engine.Thrust(), cluster.Thrust())
	fmt.Println(engine.Isp(), cluster.Isp())

	// Output:
	// 35 315
	// 25000 225000
	// 311 311
}
