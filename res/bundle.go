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
	"github.com/google/uuid"
	"log"
	"os"
	"strings"
)

func BiobankBundle() Object {
	entries := make(Array, 0, 11)
	entries = append(entries, entry(Biobank()))
	for i := 0; i < 10; i++ {
		entries = append(entries, entry(Collection(i)))
	}
	return Object{
		"resourceType": "Bundle",
		"id":           uuid.New().String(),
		"type":         "transaction",
		"entry":        entries,
	}
}

func Bundle(dir string) Object {
	// IDK why 100 TO DO
	entries := make(Array, 0, 100)
	file, err := os.Open("./" + dir)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	for _, name := range list {
		bhImport, _ := readFile(dir, name)
		for _, CP := range bhImport.Patients {
			patient := Patient(CP)
			entries = append(entries, entry(patient))
			for k := 0; k < len(CP.Locations.Locations[0].Events.Events); k++ {
				var name = CP.Locations.Locations[0].Events.Events[k].Type
				switch name {
				case "Sample":
					entries = appendSpecimens(entries, CP, CP.Locations.Locations[0].Events.Events[k])
				case "Histopathology":
					entries = appendConditions(entries, CP.PatientId, CP.Locations.Locations[0].Events.Events[k].LG.Form2.Diagnosis,
						CP.Locations.Locations[0].BasicData.Form.DiagnosisDate)
				}
			}

		}
	}

	return Object{
		"resourceType": "Bundle",
		"id":           uuid.New().String(),
		"type":         "transaction",
		"entry":        entries,
	}
}

func appendConditions(entries Array, patientIdx string, condition string, date string) Array {
	entries = append(entries, entry(Condition(patientIdx, condition[strings.LastIndex(condition, " ")+1:], date)))
	return entries
}

func appendSpecimens(entries Array, CP CohortPatient, sample Event) Array {
	entries = append(entries, entry(Specimen(CP, sample)))
	return entries
}

func entry(resource Object) Object {
	return Object{
		"fullUrl":  fmt.Sprintf("http://example.com/%s/%s", resource["resourceType"], resource["id"]),
		"resource": resource,
		"request": Object{
			"method": "PUT",
			"url":    fmt.Sprintf("%s/%s", resource["resourceType"], resource["id"]),
		},
	}
}
