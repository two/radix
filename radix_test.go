package radix

import (
	"fmt"
	"testing"
)

func TestRadix_Insert(t *testing.T) {
	r := NewRadix()
	r.Insert("abc")
	r.Insert("abd")
	r.Insert("ab")
	r.Insert("bcd")
	r.Insert("bd")
	r.Insert("bc")
	fmt.Printf("%v\n", r)
}
