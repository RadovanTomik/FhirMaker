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
	"math"
	"math/rand"
	"time"
)

func BodyHeight(patientIdx string, date time.Time, value float64) Object {
	return Object{
		"resourceType":      "Observation",
		"id":                fmt.Sprintf("bbmri-%d-body-height", patientIdx),
		"meta":              meta("https://fhir.bbmri.de/StructureDefinition/BodyHeight"),
		"status":            "final",
		"category":          Array{vitalSigns},
		"subject":           patientReference(patientIdx),
		"code":              codeableConcept(coding("http://loinc.org", "8302-2")),
		"effectiveDateTime": date.Format("2006-01-02"),
		"valueQuantity":     quantity(math.Round(value), "cm"),
	}
}

func RandBodyHeightValue(r *rand.Rand) float64 {
	return r.NormFloat64()*30.8 + 171
}
