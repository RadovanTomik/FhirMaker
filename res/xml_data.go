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
	DMs      []DM     `xml:"diagnosisMaterial"`
}
type DM struct {
	TissueId     string `xml:"sampleId,attr"`
	Diagnosis    string `xml:"diagnosis"`
	MaterialType string `xml:"materialType"`
}

func readFile(inputDir string, fileName string) (PatientMOU, error) {
	xmlFile, err := os.Open(inputDir + "/"+ fileName)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Successfully Opened %s.xml\n", fileName)

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