package osm

import "io"

// Decoder reads and decode OSM structure
type Decoder struct {
	r io.Reader
}

// NewDecoder returns a new decoder that reads from r.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

func (dec *Decoder) Decode() {

}
