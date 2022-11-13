package calendar

import "errors"

type Date struct{
	year int
	month int
	day int
}

func (d *Date) GetYear() int {
	return d.year
}

func (d *Date) SetYear(year int) error {	// pass by value
	if year < 1 {
		return errors.New("유효하지 않은 년도입니다.")
	}
	d.year = year

	return nil
}

func (d *Date) GetMonth() int {
	return d.month
}

func (d *Date) SetMonth(month int) error {	// pass by value
	if month < 1 || 12 < month {
		return errors.New("유효하지 않은 달입니다.")
	}
	d.month = month

	return nil
}

func (d *Date) GetDay() int {
	return d.day
}

func (d *Date) SetDay(day int) error {	// pass by value
	if day < 1  || 32 < day{
		return errors.New("유효하지 않은 일입니다.")
	}
	d.day = day

	return nil
}