package part

import "fmt"

// Creating a cluster of nine Rutherford engines. Shows how e.g.
// thrust gets combined but specific impulse (Isp) stays the same.
func ExampleEngineCluster() {
	cluster := EngineCluster{Rutherford{}, 9}
	fmt.Println("Mass: ", cluster.Mass())
	fmt.Println("Thrust: ", cluster.Thrust())
	fmt.Println("Isp: ", cluster.Isp())

	// Output:
	// Mass:  315
	// Thrust:  225000
	// Isp:  311
}
