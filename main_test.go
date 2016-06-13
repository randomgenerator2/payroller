package main

import (
	"strings"
	"testing"
)

func TestReadingFile(t *testing.T) {
	data := ReadFile("test_input/test.csv")
	if 6 != len(data) {
		t.Error("test.csv length should be 6, it was:", len(data))
	}
}

func TestHeaderOrder(t *testing.T) {
	data := ReadFile("test_input/test.csv")
	header := strings.Split(data[0], ",")
	HeaderIndexes(header)
	if 4 != startIndex {
		t.Error("index of Start should have been 4 but it was:", startIndex)
	}
	if 3 != endIndex {
		t.Error("index of End should have been 3 but it was:", endIndex)
	}
}

func TestReadingData(t *testing.T) {
	data := ReadFile("test_input/test.csv")
	header := strings.Split(data[0], ",")
	HeaderIndexes(header)
	employees, shifts := ReadData(data, header)
	if 3 != len(employees) {
		t.Error("Employees length should have been 3, but it was:", len(employees))
	}
	if 2 != len(shifts["asd"]) {
		t.Error("There should be 2 shifts for Noob but there was:", len(shifts["asd"]))
	}
}

func TestCreatingAllDateShifts(t *testing.T) {
	data := ReadFile("test_input/test.csv")
	header := strings.Split(data[0], ",")
	HeaderIndexes(header)
	employees, shifts := ReadData(data, header)
	dateshifts := CreateAllDateshifts(employees, shifts)
	for person, dateshift := range dateshifts {
		if 1 != len(dateshift) {
			t.Error("There was too many dateshifts for person:", person,
				"there should only have been 1 but there was:", len(dateshift))
		}
	}
	for _, person := range employees {
		if "1" == person.id {
			ds := dateshifts[person]["3.3.2014"]
			if 23 != ds.time {
				t.Error(person, " should have 23 hours of working hours, but was:", ds.time)
			}
			if 2 != ds.overtime25 {
				t.Error(person, " should have 2 hours of overtime with 25%, but was:", ds.overtime25)
			}
			if 2 != ds.overtime50 {
				t.Error(person, " should have 2 hours of overtime with 50%, but was:", ds.overtime50)
			}
			if 11 != ds.overtime100 {
				t.Error(person, " should have 11 hours of overtime with 100%, but was:", ds.overtime100)
			}
			if 11 != ds.evening {
				t.Error(person, " should have 11 hours of evening hours, but was:", ds.evening)
			}
		}
	}
}

func TestCalculatingPayRolls(t *testing.T) {
	data := ReadFile("test_input/test.csv")
	header := strings.Split(data[0], ",")
	HeaderIndexes(header)
	employees, shifts := ReadData(data, header)
	dateshifts := CreateAllDateshifts(employees, shifts)
	for _, p := range employees {
		for _, ds := range dateshifts[p] {
			ds.CalculatePayRoll()
		}
	}
	for _, p := range employees {
		for _, ds := range dateshifts[p] {
			if 0 >= ds.payroll {
				t.Error("Person: ", p,
					"did not got his payroll calculated", ds.payroll)
			}
		}
	}
}

func TestRounding(t *testing.T) {
	r := round(1.549, 2)
	if 1.55 != r {
		t.Error("Rounding does not work, 1.549 should have been rounded to 1.55, but was:", r)
	}
}

func TestT(t *testing.T) {

}
