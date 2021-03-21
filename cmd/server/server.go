package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/ordovician/rocket/engine"
	"github.com/ordovician/rocket/tank"
)

type partsJSON struct {
	Parts []string `json:"parts"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var header http.Header = w.Header()
	header.Set("Content-Type", "application/json")
	var encoder *json.Encoder = json.NewEncoder(w)

	parts := partsJSON{
		[]string{"engines", "tanks"},
	}

	encoder.Encode(parts)
}

func handleTanks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		tanks   map[string]tank.Tank
		err     error
		encoder *json.Encoder = json.NewEncoder(w)
	)
	encoder.SetIndent("", "    ")

	tanks, err = tank.LoadTanks("")
	if err != nil {
		log.Fatal("Could not serve info about tanks: ", err)
	}

	if r.URL.Path == "/tanks/" {
		encoder.Encode(tanks)
	} else {
		_, tankName := path.Split(r.URL.Path)
		if tank := tanks[tankName]; tank != nil {
			encoder.Encode(tank)
		}
	}
}

func handleEngines(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		engines map[string]engine.Engine
		err     error
		encoder *json.Encoder = json.NewEncoder(w)
	)
	encoder.SetIndent("", "    ")

	engines, err = engine.LoadEngines("")
	if err != nil {
		log.Fatal("Could not serve info about engines: ", err)
	}

	if r.URL.Path == "/engines/" {
		encoder.Encode(engines)
	} else {
		_, engineName := path.Split(r.URL.Path)
		if engine := engines[engineName]; engine != nil {
			encoder.Encode(engine)
		}
	}
}

func main() {
	pwd, _ := os.Getwd()
	log.Println("Current working directory: ", pwd)

	http.HandleFunc("/", handler)
	http.HandleFunc("/tanks/", handleTanks)
	http.HandleFunc("/engines/", handleEngines)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
