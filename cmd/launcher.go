package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/RadovanTomik/FhirMaker/res"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)
var rootCmd = &cobra.Command{
	Use: "fhir-maker [directory]",
	Version: "0.2",
	ValidArgs: []string{"input_directory", "output_directory"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("requires input and output directory arguments")
		}
		return checkDir(args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		dir := "./" + args[1]
		inputDir := "./" + args[0]
		err := createDir(dir)
		err = genBiobankTxFile(dir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = genTxFile(dir, inputDir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}


func Execute()  {
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

func createDir(dir string) error  {
	if _, err := os.Stat(dir); os.IsExist(err) {
		return fmt.Errorf("directory `%s` exists", dir)

	}
	{
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return err
		}}
	return nil
}

func genBiobankTxFile(dir string) error {
	return encodeToFile(dir, "biobank.json", res.BiobankBundle())
}

func genTxFile(dir string, inputDir string) error {
	return encodeToFile(dir, fmt.Sprintf("transaction-0.json"), res.Bundle(inputDir))
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