package res

import (
	"fmt"
	"github.com/clbanning/mxj/v2"
	"os"
)

func ReadFile(path string) mxj.Map {
	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer func(xmlFile *os.File) {
		err := xmlFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(xmlFile)
	mv, err := mxj.NewMapXmlReader(xmlFile)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return mv
}

func GetFHIRSpecimen(patient *mxj.Map) []Object {
	LTS, err := patient.ValuesForPath("patient.LTS.*")
	if err != nil {
		return nil
	}
	STS, err := patient.ValuesForPath("patient.STS.*")
	if err != nil {
		return nil
	}
	patientId, _ := patient.ValueForPath("patient.-id")
	var fhirObjects []Object
	for index, sample := range LTS {
		temp := sample.(map[string]interface{})
		temp["patientId"] = patientId
		temp["sampleIndex"] = index
		temp["collection"] = "LTS"
		fhirObjects = append(fhirObjects, Specimen(temp))
	}
	for index, sample := range STS {
		temp := sample.(map[string]interface{})
		temp["patientId"] = patientId
		temp["sampleIndex"] = index + len(LTS)
		temp["collection"] = "STS"
		fhirObjects = append(fhirObjects, Specimen(temp))
	}
	return fhirObjects
}
