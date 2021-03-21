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

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		debug.Println("Failed to read CSV file: ", err)
		return nil, err
	}

	engines := make(map[string]Engine)
	for _, record := range records[1:] {
		name := record[0]
		mass, err := strconv.ParseFloat(record[2], 64)
		if err != err {
			return nil, err
		}

		thrust, err := strconv.ParseFloat(record[3], 64)
		if err != err {
			return nil, err
		}

		isp, err := strconv.ParseFloat(record[6], 64)
		if err != err {
			return nil, err
		}

		// Values read are in metric tonnes and kilo Newton so we need to convert
		// to SI units expected by the Engine struct
		engine := NewCustomEngine(Kg(mass*1000), Newton(thrust*1000), isp)
		engines[name] = engine
	}

	return engines, nil
}
