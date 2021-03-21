package engine

import (
	"embed"
	"encoding/csv"
	"io/fs"
	"log"
	"os"
	"strconv"

	. "github.com/ordovician/rocket/physics"
)

var debug *log.Logger = log.New(os.Stdout, "", 0)

//go:embed rocket-engines.csv
var storage embed.FS

// LoadTanks load information about propellant tanks into a dictionary.
func LoadEngines(filename string) (map[string]Engine, error) {
	var (
		file fs.File
		err  error
	)

	if filename == "" {
		file, err = storage.Open("rocket-engines.csv")
	} else {
		file, err = os.Open(filename)
	}

	if err != nil {
		debug.Println("Failed to to open ", filename, ": ", err)
		dir, _ := os.Getwd()
		debug.Println("Currently in directory: ", dir)
		return nil, err
	}

	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		debug.Println("Failed to read CSV file: ", err)
		return nil, err
	}

	engines := make(map[string]Engine)
	for _, line := range lines[1:] {
		name := line[0]
		mass, err := strconv.ParseFloat(line[2], 64)
		if err != err {
			return nil, err
		}

		thrust, err := strconv.ParseFloat(line[3], 64)
		if err != err {
			return nil, err
		}

		isp, err := strconv.ParseFloat(line[6], 64)
		if err != err {
			return nil, err
		}

		engine := NewCustomEngine(Kg(mass), Newton(thrust), isp)
		engines[name] = engine
	}

	return engines, nil
}