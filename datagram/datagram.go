package datagram

import (
	"fmt"
)

// Datagram is a struct that represents a datagram
type Datagram struct {
	Source      string
	Destination string
	Data        string
}

// NewDatagram creates a new datagram
func NewDatagram(source string, destination string, data string) *Datagram {
	return &Datagram{
		Source:      source,
		Destination: destination,
		Data:        data,
	}
}

// String returns a string representation of the datagram
func (d *Datagram) String() string {
	return fmt.Sprintf("Source: %s, Destination: %s, Data: %s", d.Source, d.Destination, d.Data)
}
