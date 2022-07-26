package cmd

import (
	"FhirMaker/res"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/clbanning/mxj/v2"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"time"
)

var rootCmd = &cobra.Command{
	Use:       "fhir-maker [input directory] [output directory]",
	Version:   "0.4",
	ValidArgs: []string{"input_directory", "output_directory"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("requires input and output directory arguments")
		}
		return checkDir(args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		inputDir := args[0]
		outputDir := args[1]
		err := createDir(outputDir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//err = genTxFile(outputDir, inputDir, 0)
		err = transformFiles(inputDir, outputDir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func checkDir(dir string) error {
	if info, err := os.Stat(dir); os.IsNotExist(err) {
		return fmt.Errorf("directory `%s` doesn't exist", dir)
	} else if !info.IsDir() {
		return fmt.Errorf("`%s` isn't a directory", dir)
	} else {
		return nil
	}
}

func createDir(dir string) error {
	if _, err := os.Stat(dir); os.IsExist(err) {
		return fmt.Errorf("directory `%s` exists", dir)

	}
	{
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func genBiobankTxFile(dir string) error {
	return encodeToFile(dir, "biobank.json", res.BiobankBundle())
}

func transformFiles(inputDir string, outputDir string) error {
	counter := 0
	transactionCounter := 0
	var fhirResources []res.Object
	start := time.Now()
	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if path != inputDir {
			fmt.Printf("File: %s\n", path)
			patientMou := res.ReadFile(path)
			if patientMou != nil {
				counter++
				fhirResources = append(fhirResources, xmlToFhir(patientMou)...)
			}
			if counter%100 == 0 {
				err := encodeToFile(outputDir, fmt.Sprintf("transaction-%d.json", transactionCounter), res.Bundle(fhirResources))
				if err != nil {
					return err
				}
				transactionCounter++
				fhirResources = nil
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = encodeToFile(outputDir, fmt.Sprintf("transaction-%d.json", transactionCounter), res.Bundle(fhirResources))
	if err != nil {
		return err
	}
	elapsed := time.Since(start)
	fmt.Printf("Successfully transformed %d patient files in %s.", counter, elapsed)
	return nil
}

func xmlToFhir(patientMOU mxj.Map) []res.Object {
	return append([]res.Object{res.Patient(&patientMOU)}, res.GetFHIRSpecimen(&patientMOU)...)
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
