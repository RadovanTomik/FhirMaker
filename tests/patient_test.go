package tests

import (
	"FhirMaker/res"
	"reflect"
	"testing"
)

func TestReadBogusFile(t *testing.T) {
	patientMap := res.ReadFile("nonexistent")
	if patientMap != nil {
		t.Error()
	}
}
func TestReadRealFile(t *testing.T) {
	patientMap := res.ReadFile("resources/BBM210610122129-012593.XML")
	if reflect.TypeOf(patientMap).String() != "mxj.Map" {
		t.Error()
	}
}
func TestPatientId(t *testing.T) {
	patientMap := res.ReadFile("resources/BBM210610122129-012593.XML")
	patientId, _ := patientMap.ValueForPath("patient.-id")
	if patientId.(string) != "999" {
		t.Error()
	}
}
