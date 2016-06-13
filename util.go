package main

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func round(number float64, digits float64) float64 {
	var roundedNumber float64
	power := math.Pow(10.0, digits)
	poweredNumber := number * power
	checkNumber := math.Floor(poweredNumber) * 10
	if poweredNumber*10-checkNumber > 4 {
		roundedNumber = math.Floor(poweredNumber)/power + 1/power
	} else {
		roundedNumber = math.Floor(poweredNumber) / power
	}
	return roundedNumber
}

func GetTime(time string) float64 {
	timeSlice := strings.Split(time, ":")
	timeHours, err := strconv.ParseFloat(timeSlice[0], 64)
	check(err)
	timeMins, err := strconv.ParseFloat(timeSlice[1], 64)
	check(err)
	return timeHours + timeMins/15/4
}

func ReadFile(file string) []string {
	data, err := ioutil.ReadFile(file)
	check(err)
	return strings.Split(string(data), "\n")
}

func HeaderIndexes(header []string) {
	for i, t := range header {
		switch string(t) {
		case "Person Name":
			nameIndex = i
		case "Person ID":
			idIndex = i
		case "Date":
			dateIndex = i
		case "Start":
			startIndex = i
		case "End":
			endIndex = i
		}
	}
}

func ReadData(data []string, header []string) ([]person, map[string][]*shift) {
	employees := make([]person, 0)
	shifts := make(map[string][]*shift)
	for index, personsShift := range data {
		ps := strings.Split(string(personsShift), ",")
		if index > 0 && len(ps) == len(header) {
			p := person{name: ps[nameIndex], id: ps[idIndex]}
			employees = AddIfMissingEmployee(employees, p)
			s := &shift{date: ps[dateIndex], start: ps[startIndex], end: ps[endIndex]}
			shifts[p.id] = append(shifts[p.id], s)
		}
	}
	return employees, shifts
}

func AddIfMissingEmployee(employees []person, p person) []person {
	for _, employee := range employees {
		if employee.id == p.id {
			return employees
		}
	}
	employees = append(employees, p)
	return employees
}

func CreateAllDateshifts(employees []person, shifts map[string][]*shift) map[person]map[string]*dateshift {
	dateshifts := make(map[person]map[string]*dateshift)
	for _, p := range employees {
		for i, s := range shifts {
			if p.id == i {
				for _, ss := range s {
					CreateNewDateshift(dateshifts, p, ss.date)
					dateshifts[p][ss.date].AddTime(ss)
				}
			}
		}
	}
	return dateshifts
}

func CreateNewDateshift(dateshifts map[person]map[string]*dateshift, p person, date string) {
	for key, _ := range dateshifts[p] {
		if key == date {
			return
		}
	}
	if dateshifts[p] == nil {
		dateshifts[p] = make(map[string]*dateshift)
	}
	dateshifts[p][date] = &dateshift{}
}
