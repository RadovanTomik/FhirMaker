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
)

func Condition(patientId int, condition string, number int) Object {
	return Object{
		"resourceType":  "Condition",
		"id":            fmt.Sprintf("bbmri-%d-condition-%d", patientId, number),
		"meta":          meta("https://fhir.bbmri.de/StructureDefinition/Condition"),
		"subject":       patientReference(patientId),
		"code":          codeableConcept(codingWithVersion("http://hl7.org/fhir/sid/icd-10", "2016", condition)),
	}
}

