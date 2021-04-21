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
)

func Specimen(mou PatientMOU, specimenIdx int,) Object {
	return Object{
		"resourceType": "Specimen",
		"id":           fmt.Sprintf("bbmri-%d-specimen-%d", mou.Id, specimenIdx),
		"meta":         meta("https://fhir.bbmri.de/StructureDefinition/Specimen"),
		"extension":    Array{storageTemp(), sampleDiagnosis(mou.STS.DMs[0].Diagnosis), custodian(mou.Custodian)},
		"type":         codeableConcept(materialTypeCoding()),
		"subject":      patientReference(mou.Id),
	}
}

func storageTemp() Object {
	coding := coding(
		"https://fhir.bbmri.de/CodeSystem/StorageTemperature",
		storageTemps[0])

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
		stringReference("Organization", fmt.Sprintf("collection-" + "8")))
}


var fastingStatus = []string{"F", "FNA", "NF", "NG"}

func randFastingStatus(r *rand.Rand) string {
	return fastingStatus[r.Intn(len(fastingStatus))]
}

func randIcdOCode(r *rand.Rand) string {
	return fmt.Sprintf("C%02d.%d", r.Intn(100), r.Intn(10))
}

func materialTypeCoding() Object {
	return coding("https://fhir.bbmri.de/CodeSystem/SampleMaterialType",
		materialTypes[0])
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

func randMaterialType(r *rand.Rand) string {
	return materialTypes[r.Intn(len(materialTypes))]
}
