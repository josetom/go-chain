package common

import (
	"bytes"
)

const AddressLength = 20

type Address [AddressLength]byte

var ZeroAddress = NewAddress("0x0000000000000000000000000000000000000000")

// Accepts a hex encoded address string eg. 0x30303030
func NewAddress(value string) Address {
	var a Address
	a.SetBytes(Hex2Bytes(value))
	return a
}

func (a Address) Bytes() []byte {
	return a[:]
}

// If len(b) > AddressLength, b will be cropped from the left
func BytesToAddress(b []byte) Address {
	var a Address
	a.SetBytes(b)
	return a
}

// SetBytes sets the address to the value of b.
// If b is larger than len(a), b will be cropped from the left.
func (a *Address) SetBytes(b []byte) {
	if len(b) > AddressLength {
		b = b[len(b)-AddressLength:]
	}
	copy(a[:], b)
}

func (a *Address) Hash() Hash {
	return BytesToHash(a.Bytes())
}

func (a *Address) Equal(b Address) bool {
	return bytes.Equal(a.Bytes(), b.Bytes())
}

// implement stringer interface
func (a Address) String() string {
	return Bytes2Hex(a.Bytes(), true)
}

// MarshalJSON implements the json.Marshaler interface.
func (a Address) MarshalText() ([]byte, error) {
	return MarshalUtil(a.Bytes()), nil
}

// UnmarshalText parses a hash in hex syntax.
func (a *Address) UnmarshalText(input []byte) error {
	result, err := UnmarshalUtil(input)
	a.SetBytes(result)
	return err
}
