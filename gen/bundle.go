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
	"encoding/xml"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"os"
)

type PatientMOU struct {
	XMLName    xml.Name `xml:"patient"`
	Id         int      `xml:"id,attr"`
	Sex        string   `xml:"sex,attr"`
	BirthYear  string   `xml:"year,attr"`
	BirthMonth string   `xml:"month,attr"`
	Custodian  string   `xml:"biobank,attr"`
	LTS        LTS      `xml:"LTS"`
	STS        STS      `xml:"STS"`
}

type LTS struct {
	XMLName xml.Name `xml:"LTS"`
	Tissues []Tissue `xml:"tissue"`
}

type Tissue struct {
	XMLName      xml.Name `xml:"tissue"`
	TissueId     string   `xml:"sampleId,attr"`
	MaterialType string   `xml:"materialType"`
}

type STS struct {
	XMLName xml.Name `xml:"STS"`
	DMs      []DM     `xml:"diagnosisMaterial"`
}
type DM struct {
	TissueId     string `xml:"sampleId,attr"`
	Diagnosis    string `xml:"diagnosis"`
	MaterialType string `xml:"materialType"`
}

func readFile(string2 string) (PatientMOU, error) {
	xmlFile, err := os.Open("./input/" + string2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened {string2}.xml")

	defer func(xmlFile *os.File) {
		err := xmlFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(xmlFile)

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var patient PatientMOU
	err = xml.Unmarshal(byteValue, &patient)
	if err != nil {
		return patient, nil
	}
	return patient, err
}

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

func Bundle() Object {
	// IDK why 100 TO DO
	entries := make(Array, 0, 100)
	file, err := os.Open("./input")
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
		patientMou, _ := readFile(name)
		patient := Patient(patientMou)
		entries = append(entries, entry(patient))
		entries = appendConditions(entries, patientMou.Id, patientMou.STS.DMs[0].Diagnosis)
		entries = appendSpecimens(entries, patientMou)

	}

	return Object{
		"resourceType": "Bundle",
		"id":           uuid.New().String(),
		"type":         "transaction",
		"entry":        entries,
	}
}

func appendConditions(entries Array, patientIdx int, condition string) Array {
	entries = append(entries, entry(Condition(patientIdx, condition)))
	return entries
}

func appendSpecimens(entries Array, mou PatientMOU) Array {
	for i := 0; i < len(mou.LTS.Tissues); i++ {
		entries = append(entries, entry(Specimen(mou, i)))
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
