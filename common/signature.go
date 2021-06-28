package common

import "bytes"

const SignatureLength = 65

type Signature [SignatureLength]byte

func (s Signature) String() string {
	return Bytes2Hex(s.Bytes(), true)
}

func (s Signature) Bytes() []byte {
	return s[:]
}

// MarshalJSON implements the json.Marshaler interface.
func (s Signature) MarshalText() ([]byte, error) {
	return MarshalUtil(s.Bytes()), nil
}

// UnmarshalText parses a Signature in hex syntax.
func (s *Signature) UnmarshalText(input []byte) error {
	result, err := UnmarshalUtil(input)
	s.SetBytes(result)
	return err
}

// SetBytes sets the address to the value of b.
// If b is larger than len(a), b will be cropped from the left.
func (s *Signature) SetBytes(b []byte) {
	if len(b) > SignatureLength {
		b = b[len(b)-SignatureLength:]
	}
	copy(s[:], b)
}

func (a *Signature) Equal(b Signature) bool {
	return bytes.Equal(a.Bytes(), b.Bytes())
}

func (a Signature) IsEmpty() bool {
	emptySignature := Signature{}
	return a.Equal(emptySignature)
}
