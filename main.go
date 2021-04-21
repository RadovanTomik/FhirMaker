package main

import (
	"encoding/json"
	"fmt"
	"github.com/samply/bbmri-fhir-gen/res"
	"os"
	"path/filepath"
)

func genBiobankTxFile(dir string) error {
	return encodeToFile(dir, "biobank.json", res.BiobankBundle())
}

func genTxFile(dir string) error {
	return encodeToFile(dir, fmt.Sprintf("transaction-0.json"), res.Bundle())
}

// encodeToFile encodes the JSON object `o` to the file with `filename` in `dir`
func encodeToFile(dir string, filename string, o res.Object) error {
	f, err := os.OpenFile(filepath.Join(dir, filename), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	e := json.NewEncoder(f)
	e.SetIndent("", "  ")
	err = e.Encode(o)
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

func main() {
	err := genBiobankTxFile("./output")
	if err != nil {
		return
	}
	err = genTxFile("./output")
	if err != nil {
		return
	}

}
