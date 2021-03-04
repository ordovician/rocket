package tank

import (
	"log"
	"os"
	"strconv"
	"strings"

	. "github.com/ordovician/rocket/physics"
)

var debug *log.Logger = log.New(os.Stdout, "", 0)

// LoadTanks load information about propellant tanks into a dictionary.
func LoadTanks() (map[string]Tank, error) {

	filename := "../data/propellant-tanks.csv"
	data, err := os.ReadFile(filename)

	if err != nil {
		debug.Println("Failed to to open ", filename, ": ", err)
		dir, _ := os.Getwd()
		debug.Println("Currently in directory: ", dir)
		return nil, err
	}

	str := string(data)
	lines := strings.Split(str, "\n")
	tanks := make(map[string]Tank)

	for _, line := range lines[1:] {
		parts := strings.Split(line, ",")

		dryMassStr := parts[3]
		drymass, err := strconv.ParseFloat(dryMassStr, 64)
		if err != nil {
			debug.Println("Could not parse '", dryMassStr, "' as a floating point value")
			return nil, err
		}

		totalMassStr := parts[2]
		totalMass, err := strconv.ParseFloat(totalMassStr, 64)
		if err != nil {
			debug.Println("Could not parse '", totalMassStr, "' as a floating point value")
			return nil, err
		}

		tank := FlexiTank{
			DryMass:   Kg(drymass),
			TotalMass: Kg(totalMass),
		}

		name := parts[0]
		tank.Refill()
		tanks[name] = &tank
	}

	return tanks, nil
}
