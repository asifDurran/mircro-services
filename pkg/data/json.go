package data

import (
	"encoding/json"
	"io"
)

// ToJSON serialize the given interface into a string base json format
func ToJSON (i interface{}, w io.Writer) error {
	 e := json.NewEncoder(w)

	 return e.Encode(i)
}

// FromJSON deserilize the object from JSON string
// in an io.Reader to the given interface

func FromJSON(i interface{}, r io.Reader) error {

	d := json.NewDecoder(r)

	return d.Decode(i)
}