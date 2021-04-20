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

package gen

import (
	"fmt"
	"strings"
)

func Patient(mou PatientMOU) Object {
	patient := make(map[string]interface{})
	patient["resourceType"] = "Patient"
	patient["id"] = fmt.Sprintf("bbmri-%d", mou.Id)
	patient["meta"] = meta("https://fhir.bbmri.de/StructureDefinition/Patient")
	patient["gender"] = mou.Sex
	patient["birthDate"] = mou.BirthYear + "-" + strings.Trim(mou.BirthMonth, "-") + "-01"

	return patient
}
