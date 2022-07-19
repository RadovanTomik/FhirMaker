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

func Biobank() Object {
	return Object{
		"resourceType": "Organization",
		"id":           "MMCI-0",
		"meta":         meta("https://fhir.bbmri.de/StructureDefinition/Biobank"),
		"extension": Array{
			bbmriExtensionString(
				"OrganizationDescription",
				"Masaryk Memorial Cancer Institute"),
			bbmriExtensionString(
				"JuridicalPerson",
				"amdfsoignsd"),
			bbmriExtensionCodeableConcept(
				"QualityStandard",
				codeableConcept(bbmriCoding("QualityStandard", "oecd-guidelines"))),
			bbmriExtensionCodeableConcept(
				"QualityStandard",
				codeableConcept(bbmriCoding("QualityStandard", "iso-15189"))),
		},
		"identifier": Array{Object{
			"system": "http://www.bbmri-eric.eu/",
			"value":  "bbmri-eric:ID:CZ_MMCI",
		}},
		"name":  "Zentrale Biobank Neustadt",
		"alias": Array{"ZBBN"},
		"telecom": Array{Object{
			"system": "url",
			"value":  "http://biobank.klinikum-neustadt.de",
		}},
		"address": Array{Object{
			"line":       Array{"Krankenhausstr. 12"},
			"city":       "Neustadt",
			"postalCode": "60200",
			"country":    "CZ",
		}},
		"contact": Array{Object{
			"purpose": Object{
				"coding": Array{Object{
					"system":  "http://terminology.hl7.org/CodeSystem/contactentity-type",
					"code":    "ADMIN",
					"display": "Administrative",
				}},
			},
			"name": Object{
				"family": "Chef",
				"given":  Array{"Claudia", "Celine"},
				"prefix": Array{"Prof.", "Dr.", "Dr."},
			},
			"telecom": Array{Object{
				"system": "phone",
				"value":  "0452 4624-24",
			}},
		}, Object{
			"purpose": Object{
				"coding": Array{Object{
					"system":  "https://fhir.bbmri.de/CodeSystem/ContactType",
					"code":    "RESEARCH",
					"display": "Research",
				}},
			},
			"name": Object{
				"family": "Forschungsbeauftragter",
				"given":  Array{"Friedrich"},
				"prefix": Array{"Freiherr", "Dr."},
				"suffix": Array{"M.Sc."},
			},
			"telecom": Array{Object{
				"system": "phone",
				"value":  "0452 4624-28",
			}, Object{
				"system": "email",
				"value":  "forschungsanfragen@biobank.klinikum-neustadt.de",
			}},
			"address": Object{
				"line":       Array{"Brno"},
				"city":       "Brno",
				"postalCode": "60200",
				"country":    "CZ",
			},
		}},
	}
}
