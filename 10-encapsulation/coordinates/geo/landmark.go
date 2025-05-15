package geo

import (
	"errors"
	"unicode/utf8"
)

type Landmark struct {
	name string
	Coordinates
}

func (l *Landmark) Name() string {
	return l.name
}

func (l *Landmark) SetName(name string) error {
	if utf8.RuneCountInString(name) == 0 {
		return errors.New("invalid name")
	}

	l.name = name
	return nil
}
