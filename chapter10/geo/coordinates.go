package geo

import "errors"

type Coordinates struct {
	latitude  float64
	longitude float64
}

// Latitude() returns the latitude of the coordinates.
func (c *Coordinates) Latitude() float64 {
	return c.latitude
}

// Longitude() returns the longitude of the coordinates.
func (c *Coordinates) Longitude() float64 {
	return c.longitude
}

// SetLatitude sets the latitude of the coordinates.
func (c *Coordinates) SetLatitude(latitude float64) error {
	if latitude < -90 || latitude > 90 {
		return errors.New("invalid latitude")
	}
	c.latitude = latitude
	return nil
}

// SetLongitude sets the longitude of the coordinates.
func (c *Coordinates) SetLongitude(longitude float64) error {
	if longitude < -180 || longitude > 180 {
		return errors.New("invalid longitude")
	}
	c.longitude = longitude
	return nil
}
