package res

// Data imported from XML files exported from MOU
import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// MOU
type PatientMOU struct {
	XMLName    xml.Name `xml:"patient"`
	Id         int      `xml:"id,attr"`
	Sex        string   `xml:"sex,attr"`
	BirthYear  string   `xml:"year,attr"`
	BirthMonth string   `xml:"month,attr"`
	Custodian  string   `xml:"biobank,attr"`
	LTS        LTS      `xml:"LTS"`
	STS        STS      `xml:"STS"`
}

type LTS struct {
	XMLName xml.Name `xml:"LTS"`
	Tissues []Tissue `xml:"tissue"`
}

type Tissue struct {
	XMLName      xml.Name `xml:"tissue"`
	TissueId     string   `xml:"sampleId,attr"`
	MaterialType string   `xml:"materialType"`
}

type STS struct {
	XMLName xml.Name `xml:"STS"`
	DMs     []DM     `xml:"diagnosisMaterial"`
}
type DM struct {
	TissueId     string `xml:"sampleId,attr"`
	Diagnosis    string `xml:"diagnosis"`
	MaterialType string `xml:"materialType"`
}

// COHORT

type BHImport struct {
	XMLName  xml.Name        `xml:"BHImport"`
	Patients []CohortPatient `xml:"BHPatient"`
}

type CohortPatient struct {
	XMLName   xml.Name  `xml:"BHPatient"`
	PatientId string    `xml:"Identifier"`
	Locations Locations `xml:"Locations"`
}

type Locations struct {
	XMLName   xml.Name   `xml:"Locations"`
	Locations []Location `xml:"Location"`
}

type Location struct {
	XMLName   xml.Name  `xml:"Location"`
	Name      string    `xml:"name,attr"`
	BasicData BasicData `xml:"BasicData"`
	Events    Events    `xml:"Events"`
}

type Events struct {
	XMLName xml.Name `xml:"Events"`
	Events  []Event  `xml:"Event"`
}
type Event struct {
	XMLName xml.Name        `xml:"Event"`
	Date    string          `xml:"name,attr"`
	Type    string          `xml:"eventtype,attr"`
	LG      LogitudinalData `xml:"LogitudinalData"`
}
type LogitudinalData struct {
	XMLName xml.Name `xml:"LogitudinalData"`
	Form1   Form1    `xml:"Form1,omitempty"`
	Form2   Form2    `xml:"Form2,omitempty"`
}

type Form1 struct {
	XMLName  xml.Name `xml:"Form1"`
	SampleId string   `xml:"Dataelement_56_2"`
	Type     string   `xml:"Dataelement_54_2"`
	Year     string   `xml:"Dataelement_89_3"`
}

type Form2 struct {
	XMLName   xml.Name `xml:"Form2"`
	Diagnosis string   `xml:"Dataelement_92_1"`
}
type BasicData struct {
	XMLName xml.Name `xml:"BasicData"`
	Form    Form     `xml:"Form"`
}
type Form struct {
	XMLName       xml.Name `xml:"Form"`
	Gender        string   `xml:"Dataelement_85_1"`
	DiagnosisAge  string   `xml:"Dataelement_3_1"`
	DiagnosisDate string   `xml:"Dataelement_51_3"`
}

func readFile(inputDir string, fileName string) (BHImport, error) {
	xmlFile, err := os.Open(inputDir + "/" + fileName)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Successfully Opened %s\n", fileName)

	defer func(xmlFile *os.File) {
		err := xmlFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(xmlFile)

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var bhImport BHImport
	err = xml.Unmarshal(byteValue, &bhImport)
	if err != nil {
		return bhImport, nil
	}
	return bhImport, err
}
