package common

import "bytes"

const HashLength = 32

type Hash [HashLength]byte

func (h Hash) String() string {
	return Bytes2Hex(h.Bytes(), true)
}

func (h Hash) Bytes() []byte {
	return h[:]
}

// MarshalJSON implements the json.Marshaler interface.
func (h Hash) MarshalText() ([]byte, error) {
	return MarshalUtil(h.Bytes()), nil
}

// UnmarshalText parses a hash in hex syntax.
func (h *Hash) UnmarshalText(input []byte) error {
	result, err := UnmarshalUtil(input)
	h.setBytes(result)
	return err
}

// SetBytes sets the address to the value of b.
// If b is larger than len(a), b will be cropped from the left.
func (h *Hash) setBytes(b []byte) {
	if len(b) > HashLength {
		b = b[len(b)-HashLength:]
	}
	copy(h[:], b)
}

func (a *Hash) Equal(b Hash) bool {
	return bytes.Equal(a.Bytes(), b.Bytes())
}

func (a Hash) IsEmpty() bool {
	emptyHash := Hash{}
	return a.Equal(emptyHash)
}
