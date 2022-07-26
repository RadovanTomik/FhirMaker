// Copyright © 2019 The Samply Development Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package res

import (
	"fmt"
	"regexp"
)

func Specimen(sample map[string]interface{}) Object {

	return Object{
		"resourceType": "Specimen",
		"id":           fmt.Sprintf("bbmri-%s-specimen-%d", sample["patientId"].(string), sample["sampleIndex"]),
		"meta":         meta("https://fhir.bbmri.de/StructureDefinition/Specimen"),
		"extension":    generateExtensions(sample),
		"type":         codeableConcept(materialTypeCoding(sample["materialType"])),
		"subject":      patientReference(sample["patientId"].(string)),
	}
}

func storageTemp(collection string, materialType interface{}) Object {
	temp, ok := materialType.(string)
	if !ok {
		return nil
	}
	number := 1              //STS
	if collection == "LTS" { //tissue, serum
		number = 4
	}
	if collection == "LTS" && materialType == "gD" { // Genome
		number = 2
	}
	res1, err := regexp.MatchString("5.", temp) //RNALATER
	if err != nil {
		return nil
	}
	if res1 {
		number = 1
	}
	coding := coding(
		"https://fhir.bbmri.de/CodeSystem/StorageTemperature",
		storageTemps[number])

	return bbmriExtensionCodeableConcept("StorageTemperature", codeableConcept(coding))
}

var storageTemps = []string{
	"temperature2to10",
	"temperature-18to-35",
	"temperature-60to-85",
	"temperatureGN",
	"temperatureLN",
	"temperatureRoom",
	"temperatureOther",
}

func generateExtensions(sample map[string]interface{}) []Object {
	extensions := []Object{custodian(sample["collection"].(string))}
	temp := storageTemp(sample["collection"].(string), sample["materialType"])
	if temp != nil {
		extensions = append(extensions, temp)
	}
	temp = sampleDiagnosis(sample["diagnosis"])
	if temp != nil {
		extensions = append(extensions, temp)
	}
	return extensions
}

func sampleDiagnosis(sampleDiagnosis interface{}) Object {
	diagnosis, ok := sampleDiagnosis.(string)
	if !ok {
		return nil
	}
	locatorDiagnosis := diagnosis
	if len(diagnosis) == 4 {
		suffix := diagnosis[3:]
		locatorDiagnosis = diagnosis[:3] + "." + suffix
	}
	coding := coding("http://hl7.org/fhir/sid/icd-10", locatorDiagnosis)

	return bbmriExtensionCodeableConcept("SampleDiagnosis", codeableConcept(coding))
}

func custodian(custodian string) Object {
	return bbmriExtensionReference(
		"Custodian",
		Object{"identifier": Object{
			"system": "https://bbmri-eric.eu",
			"value":  fmt.Sprintf("bbmri-eric:ID:CZ_MMCI:collection:%s", custodian),
		}})
}

func materialTypeCoding(materialType interface{}) Object {
	// TODO nasty
	temp, ok := materialType.(string)
	if !ok {
		temp = "tissue-other"
	}
	if val, ok := materialTypes2[temp]; ok {
		temp = val
	}
	return coding("https://fhir.bbmri.de/CodeSystem/SampleMaterialType",
		temp)
}

var materialTypes2 = map[string]string{
	"1":  "tissue-frozen",
	"2":  "tissue-frozen",
	"3":  "tissue-frozen",
	"4":  "tissue-frozen",
	"5":  "tissue-frozen",
	"6":  "tissue-frozen",
	"7":  "peripheral-blood-cells-vital",
	"C":  "blood-plasma",
	"K":  "blood-plasma",
	"L":  "liquid-other",
	"PD": "liquid-other",
	"S":  "serum",
	"T":  "blood-plasma",
	"53": "tissue-other",
	"54": "tissue-other",
	"55": "tissue-other",
	"56": "tissue-other",
	"gD": "dna",
	"SD": "serum",
	"A1": "tissue-other",
}
