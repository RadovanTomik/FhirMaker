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
)

func BiobankBundle() Object {
	entries := make(Array, 0, 11)
	entries = append(entries, entry(Biobank()))
	entries = append(entries, entry(Collection("STS")))
	entries = append(entries, entry(Collection("LTS")))
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
		patientMou, _ := readFile(dir, name)
		if len(patientMou.LTS.Tissues) == 0 && len(patientMou.LTS.Serums) == 0 &&
			len(patientMou.LTS.Genomes) == 0 && len(patientMou.STS.DMs) == 0  {
			continue
		}
		patient := Patient(patientMou)
		entries = append(entries, entry(patient))
		entries = appendSpecimensAndConditions(entries, patientMou)

	}

	return Object{
		"resourceType": "Bundle",
		"id":           uuid.New().String(),
		"type":         "transaction",
		"entry":        entries,
	}
}

func appendSpecimensAndConditions(entries Array, mou PatientMOU) Array {
	conditions := make(map[string]struct{})
	counter := 0
	conCounter := 0
	for i := 0; i < len(mou.LTS.Tissues); i++ {
		if mou.LTS.Tissues[i].Diagnosis == "" {
			continue
		}
		if _, ok := conditions[mou.LTS.Tissues[i].Diagnosis]; ok {
		} else {
			conditions[mou.LTS.Tissues[i].Diagnosis] = struct{}{}
			entries = append(entries, entry(Condition(mou.Id, mou.LTS.Tissues[i].Diagnosis, conCounter)))
			conCounter++
		}
		entries = append(entries, entry(Specimen(mou.Id, mou.LTS.Tissues[i].Diagnosis, mou.LTS.Tissues[i].MaterialType, "LTS", i)))
		counter++
	}
	if len(mou.LTS.Tissues) > 0 {
		for i := 0; i < len(mou.LTS.Serums); i++ {
			if mou.LTS.Tissues[0].Diagnosis == "" {
				continue
			}
			entries = append(entries, entry(Specimen(mou.Id, mou.LTS.Tissues[0].Diagnosis, mou.LTS.Serums[i].MaterialType, "LTS", counter)))
			counter++
		}
		for i := 0; i < len(mou.LTS.Genomes); i++ {
			if mou.LTS.Tissues[0].Diagnosis == "" {
				continue
			}
			entries = append(entries, entry(Specimen(mou.Id, mou.LTS.Tissues[0].Diagnosis, mou.LTS.Genomes[i].MaterialType, "LTS", counter)))
			counter++
		}
	}
		for i := 0; i < len(mou.STS.DMs); i++ {
			if mou.STS.DMs[i].Diagnosis == "" {
				continue
			}
			if _, ok := conditions[mou.STS.DMs[i].Diagnosis]; ok {
			} else {
				conditions[mou.STS.DMs[i].Diagnosis] = struct{}{}
				entries = append(entries, entry(Condition(mou.Id, mou.STS.DMs[i].Diagnosis, conCounter)))
				conCounter++
			}
			entries = append(entries, entry(Specimen(mou.Id, mou.STS.DMs[i].Diagnosis, mou.STS.DMs[i].MaterialType, "STS", counter)))
			counter++
		}
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
