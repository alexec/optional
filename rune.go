// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at 2017-07-15 20:14:12.042959048 +0000 UTC

package optional

// Rune is an optional rune
type Rune struct {
	rune    rune
	present bool
}

// EmptyRune returns an empty optional.Rune
func EmptyRune() Rune {
	return Rune{}
}

// OfRune creates an optional.Rune from a rune
func OfRune(r rune) Rune {
	return Rune{rune: r, present: true}
}

// Set sets the rune value
func (o *Rune) Set(r rune) {
	o.rune = r
	o.present = true
}

// Rune returns the rune value
func (o *Rune) Rune() rune {
	return o.rune
}

// Present returns whether or not the value is present
func (o *Rune) Present() bool {
	return o.present
}

// OrElse returns the rune value or a default value if the value is not present
func (o *Rune) OrElse(r rune) rune {
	if o.present {
		return o.rune
	}
	return r
}
