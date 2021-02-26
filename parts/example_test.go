package parts

import "fmt"

func ExampleElectronBoosterEngines() {
	cluster := EngineCluster{Rutherford{}, 9}
	fmt.Println("Mass: ", cluster.Mass())
	fmt.Println("Thrust: ", cluster.Thrust())
	fmt.Println("Isp: ", cluster.Isp())

	// Output:
	// Mass:  315
	// Thrust:  225000
	// Isp:  311
}
