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
	"time"
)

func TobaccoUse(r *rand.Rand, patientIdx string, time time.Time) Object {
	tobaccoUse := make(map[string]interface{})
	tobaccoUse["resourceType"] = "Observation"
	tobaccoUse["id"] = fmt.Sprintf("bbmri-%d-tobacco-use", patientIdx)
	tobaccoUse["meta"] = meta("https://fhir.bbmri.de/StructureDefinition/TobaccoUse")
	tobaccoUse["status"] = "final"
	tobaccoUse["code"] = codeableConcept(coding("http://loinc.org", "72166-2"))
	tobaccoUse["subject"] = patientReference(patientIdx)
	tobaccoUse["effectiveDateTime"] = time.Format("2006-01-02")
	tobaccoUse["valueCodeableConcept"] = codeableConcept(coding("http://loinc.org", randSmokingStatus(r)))

	return tobaccoUse
}

var smokingStatus = []string{
	"LA18976-3",
	"LA18977-1",
	"LA15920-4",
	"LA18978-9",
	"LA18979-7",
	"LA18980-5",
	"LA18981-3",
	"LA18982-1",
}

func randSmokingStatus(r *rand.Rand) string {
	return smokingStatus[r.Intn(len(smokingStatus))]
}
