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
	"github.com/clbanning/mxj/v2"
	"strings"
)

func Patient(patientXML *mxj.Map) Object {
	patient := make(map[string]interface{})
	var err error
	patient["resourceType"] = "Patient"
	patient["id"], err = patientXML.ValueForPath("patient.-id")
	patient["meta"] = meta("https://fhir.bbmri.de/StructureDefinition/Patient")
	patient["gender"], err = patientXML.ValueForPath("patient.-sex")
	year, err := patientXML.ValueForPath("patient.-year")
	month, err := patientXML.ValueForPath("patient.-month")
	patient["birthDate"] = year.(string) + "-" + strings.Trim(month.(string), "-") + "-01"
	if err != nil {
		return nil
	}
	return patient
}
