package main

var hourWage, eveningWage float64 = 3.75, 1.15
var over25, over50, over100 float64 = 0.25, 0.5, 1

type shift struct {
	date  string
	start string
	end   string
}

type person struct {
	id   string
	name string
}

type dateshift struct {
	time        float64
	overtime25  float64
	overtime50  float64
	overtime100 float64
	evening     float64
	payroll     float64
}

func (s *shift) CalculateShiftLength() float64 {
	endTime := GetTime(s.end)
	startTime := GetTime(s.start)
	if endTime < startTime {
		endTime += 24
	}
	return endTime - startTime
}

func (s *shift) GetEveningTime() float64 {
	startTime, endTime := GetTime(s.start), GetTime(s.end)
	eveningTime, eveningStartTime, eveningEndTime := 0.0, 18.0, endTime
	startTimeEvening := 0.0

	if endTime > 0 && endTime <= 6 {
		eveningEndTime = endTime + 24
	}
	if startTime > 18 {
		eveningStartTime = startTime
	} else if startTime > 0 && startTime < 6 {
		eveningStartTime = startTime + 24
	}
	if startTime < 6 && endTime > 6 {
		startTimeEvening = 6 - startTime
		eveningStartTime = 18
	}

	eveningTime = eveningEndTime - eveningStartTime
	if eveningTime > 0 {
		return eveningTime + startTimeEvening
	}

	return 0.0 + startTimeEvening
}

func (d *dateshift) AddTime(s *shift) {
	time := s.CalculateShiftLength()
	newTime := d.time + time
	if newTime > 8 && newTime <= 10 {
		d.overtime25 = newTime - 8
	} else if newTime > 10 && newTime <= 12 {
		d.overtime25 = 2
		d.overtime50 = newTime - 10
	} else if newTime > 12 {
		d.overtime25 = 2
		d.overtime50 = 2
		d.overtime100 = newTime - 12
	}
	d.evening += s.GetEveningTime()
	d.time = newTime
}

func (d *dateshift) CalculatePayRoll() {
	d.payroll = d.time*hourWage + d.evening*eveningWage + d.overtime25*(hourWage*over25) + d.overtime50*(hourWage*over50) + d.overtime100*(hourWage*over100)
}
