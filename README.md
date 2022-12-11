# MusGo_uint64
Provides serialization of the Golang's `uint64` type.

# How to use
Simply cast `uint64` to `musgo_uint64.Uint64`:
```go
	var n uint64 = 5
	// Marshal
	buf := make([]byte, musgo_uint64.Uint64(n).SizeMUS())
	musgo_uint64.Uint64(n).MarshalMUS(buf)
	// Unmarshal
	_, err := (*musgo_uint64.Uint64)(&n).UnmarshalMUS(buf)
	if err != nil {
		panic(err)
	}
```

# More info
You can find at [github.com/ymz-ncnk/musgo](https://github.com/ymz-ncnk/musgo).

