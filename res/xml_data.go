package res

// Data imported from XML files exported from MOU
import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type PatientMOU struct {
	XMLName    xml.Name `xml:"patient"`
	Consent    string   `xml:"consent,attr"`
	Id         int      `xml:"id,attr"`
	Sex        string   `xml:"sex,attr"`
	BirthYear  string   `xml:"year,attr"`
	BirthMonth string   `xml:"month,attr"`
	Custodian  string   `xml:"biobank,attr"`
	LTS        LTS      `xml:"LTS,omitempty"`
	STS        STS      `xml:"STS,omitempty"`
}

type LTS struct {
	XMLName xml.Name `xml:"LTS"`
	Tissues []Tissue `xml:"tissue,omitempty"`
	Genomes []Genome `xml:"genome,omitempty"`
	Serums  []Serum  `xml:"serum,omitempty"`
}

type Tissue struct {
	XMLName      xml.Name `xml:"tissue"`
	SampleId     string   `xml:"sampleId,attr"`
	MaterialType string   `xml:"materialType"`
	TakingDate   string   `xml:"takingDate"`
	Diagnosis    string   `xml:"diagnosis"`
}
type Sample struct {
	XMLName      xml.Name	`xml:"tissue"`
	SampleId     string   `xml:"sampleId,attr"`
	MaterialType string   `xml:"materialType"`
	TakingDate   string   `xml:"takingDate"`
	Diagnosis    string   `xml:"diagnosis"`
}

type Genome struct {
	XMLName      xml.Name `xml:"genome"`
	SampleId     string   `xml:"sampleId,attr"`
	MaterialType string   `xml:"materialType"`
}

type Serum struct {
	XMLName      xml.Name `xml:"serum"`
	SampleId     string   `xml:"sampleId,attr"`
	MaterialType string   `xml:"materialType"`
}

type STS struct {
	XMLName xml.Name `xml:"STS"`
	DMs     []DM     `xml:"diagnosisMaterial"`
}
type DM struct {
	SampleId     string `xml:"sampleId,attr"`
	Diagnosis    string `xml:"diagnosis"`
	MaterialType string `xml:"materialType"`
	TakingDate   string `xml:"takingDate"`
}

func readFile(inputDir string, fileName string) (PatientMOU, error) {
	xmlFile, err := os.Open(inputDir + "/" + fileName)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Successfully Opened %s \n", fileName)

	defer func(xmlFile *os.File) {
		err := xmlFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(xmlFile)

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var patient PatientMOU
	err = xml.Unmarshal(byteValue, &patient)
	if err != nil {
		return patient, nil
	}
	return patient, err
}
