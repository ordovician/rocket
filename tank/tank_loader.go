package tank

import (
	"embed"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	. "github.com/ordovician/rocket/physics"
)

var debug *log.Logger = log.New(os.Stdout, "", 0)

//go:embed propellant-tanks.csv
var fs embed.FS

// LoadTanks load information about propellant tanks into a dictionary.
// If the filepath is empty then we will attempt to load from a default location
func LoadTanks(path string) (map[string]Tank, error) {
	var (
		data []byte
		err  error
	)

	if path == "" {
		data, err = fs.ReadFile("propellant-tanks.csv")
	} else {
		data, err = os.ReadFile(path)
	}

	if err != nil {
		return nil, fmt.Errorf("Unsable to open propellant tanks file. Reason: %w", err)
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

		// multiply by 1000 to convert from metric tonnes to Kg as expected
		// by tank struct
		tank := FlexiTank{
			DryMass:   Kg(drymass * 1000),
			TotalMass: Kg(totalMass * 1000),
		}

		name := parts[0]
		tank.Refill()
		tanks[name] = &tank
	}

	return tanks, nil
}
