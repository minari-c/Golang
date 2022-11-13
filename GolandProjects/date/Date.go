package date

import "errors"

type Date struct{
	year int
	month int
	day int
}

func (d *Date) SetYear(year int) error {	// pass by value
	if year < 1 {
		return errors.New("유효하지 않은 년도 입니다.")
	}
	d.year = year

	return nil
}

