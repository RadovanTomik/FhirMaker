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
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestCauseOfDeath(t *testing.T) {

	causeOfDeath := CauseOfDeath(rand.New(rand.NewSource(1)), 0)
	assert.Equal(t,
		"https://fhir.bbmri.de/StructureDefinition/CauseOfDeath",
		causeOfDeath["meta"].(Object)["profile"].(Array)[0])
	assert.Equal(t, "Observation", causeOfDeath["resourceType"])
	assert.Equal(t, "bbmri-0-cause-of-death", causeOfDeath["id"])
	assert.Equal(t, "http://loinc.org", causeOfDeath["code"].(Object)["coding"].(Array)[0].(Object)["system"])
	assert.Equal(t, "68343-3", causeOfDeath["code"].(Object)["coding"].(Array)[0].(Object)["code"])
	assert.Equal(t, "http://hl7.org/fhir/sid/icd-10", causeOfDeath["valueCodeableConcept"].(Object)["coding"].(Array)[0].(Object)["system"])
	assert.Equal(t, "Patient/bbmri-0", causeOfDeath["subject"].(Object)["reference"])

}
