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
	"math/rand"
	"regexp"
)

func Specimen(patientId int, diagnosis string, materialType string, collection string, number int) Object {
	return Object{
		"resourceType": "Specimen",
		"id":           fmt.Sprintf("bbmri-%d-specimen-%d", patientId, number),
		"meta":         meta("https://fhir.bbmri.de/StructureDefinition/Specimen"),
		"extension":    Array{storageTemp(collection, materialType), sampleDiagnosis(diagnosis), custodian(collection)},
		"type":         codeableConcept(materialTypeCoding(materialType)),
		"subject":      patientReference(patientId),
	}
}

func storageTemp(collection string, materialType string) Object {
	number := 1 //STS
	if collection == "LTS" { //tissue, serum
		number = 4
	}
	if collection == "LTS" && materialType == "gD"{ // Genome
		number = 2
	}
	res1, _ := regexp.MatchString("5.", materialType) //RNALATER
	if res1 {
		fmt.Println(materialType)
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

func randStorageTemp(r *rand.Rand) string {
	return storageTemps[r.Intn(len(storageTemps))]
}

func sampleDiagnosis(diagnosis string) Object {
	coding := coding("http://hl7.org/fhir/sid/icd-10", diagnosis)

	return bbmriExtensionCodeableConcept("SampleDiagnosis", codeableConcept(coding))
}

func custodian(custodian string) Object {
	return bbmriExtensionReference(
		"Custodian",
		stringReference("Organization", fmt.Sprintf("collection-" + custodian)))
}


var fastingStatus = []string{"F", "FNA", "NF", "NG"}

func materialTypeCoding(materialType string) Object {
	temp := "tissue-other"
	if val, ok := materialTypes2[materialType]; ok {
		temp = val
	}
	return coding("https://fhir.bbmri.de/CodeSystem/SampleMaterialType",
		temp)
}

var materialTypes = []string{
	"tissue",
	"tissue-formalin",
	"tissue-frozen",
	"tissue-paxgene-or-else",
	"tissue-other",
	"liquid",
	"whole-blood",
	"blood-plasma",
	"blood-serum",
	"peripheral-blood-cells-vital",
	"buffy-coat",
	"bone-marrow",
	"csf-liquor",
	"ascites",
	"urine",
	"saliva",
	"stool-faeces",
	"liquid-other",
	"derivative",
	"dna",
	"cf-dna",
	"rna",
	"derivative-other",
}

var materialTypes2 = map[string]string {
	"1": "tissue-frozen",
	"2": "tissue-frozen",
	"3": "tissue-frozen",
	"4": "tissue-frozen",
	"5": "tissue-frozen",
	"6": "tissue-frozen",
	"7": "peripheral-blood-cells-vital",
	"C": "blood-plasma",
	"K": "blood-plasma",
	"L": "liquid-other",
	"PD": "liquid-other",
	"S": "serum",
	"T": "blood-plasma",
	"53": "tissue-other",
	"54": "tissue-other",
	"55": "tissue-other",
	"56": "tissue-other",
	"gD": "dna",
	"SD": "serum",
	"A1": "tissue-other",



	//TODO
}

func randMaterialType(r *rand.Rand) string {
	return materialTypes[r.Intn(len(materialTypes))]
}
