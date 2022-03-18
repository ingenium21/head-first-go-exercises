package calendar

import (
	"errors"
)

type Date struct {
	year  int
	month int
	day   int
}

// SetYear sets the year of the date.
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

// SetMonth sets the month of the date.
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

// SetDay sets the day of the date.
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}

//Getter methods

// Year() returns the year of the date.
func (d *Date) Year() int {
	return d.year
}

// Month() returns the month of the date.
func (d *Date) Month() int {
	return d.month
}

// Day() returns the day of the date.
func (d *Date) Day() int {
	return d.day
}
