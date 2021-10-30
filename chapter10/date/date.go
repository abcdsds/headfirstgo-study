package date

import "errors"

type Date2 struct {
	year int
	month int
	day int
}

func (d *Date2) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

func (d *Date2) Year() int {
	return d.year
}
