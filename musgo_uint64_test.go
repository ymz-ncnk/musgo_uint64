package musgo_uint64

import "testing"

func TestMusgoUint64(t *testing.T) {
	var n uint64 = 5
	buf := make([]byte, Uint64(n).SizeMUS())
	Uint64(n).MarshalMUS(buf)

	var an uint64
	(*Uint64)(&an).UnmarshalMUS(buf)

	if n != an {
		t.Fatal("something went wrong")
	}
}
