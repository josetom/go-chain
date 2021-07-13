package types

// Iterator iterates over a DB's key/value pairs in key order.
//
// When encounter an error any 'seeks method' will return false and will
// yield no key/value pairs. The error can be queried by calling the Error
// method. Calling Release is still necessary.
//
// An iterator must be released after use, but it is not necessary to read
// an iterator until exhaustion.
// Also, an iterator is not necessarily safe for concurrent use, but it is
// safe to use multiple iterators concurrently, with each in a dedicated
// goroutine.
type Iterator interface {

	// Key returns the key of the current key/value pair, or nil if done.
	// The caller should not modify the contents of the returned slice, and
	// its contents may change on the next call to any 'seeks method'.
	Key() []byte

	// Value returns the value of the current key/value pair, or nil if done.
	// The caller should not modify the contents of the returned slice, and
	// its contents may change on the next call to any 'seeks method'.
	Value() []byte

	// Next moves the iterator to the next key/value pair.
	// It returns false if the iterator is exhausted.
	Next() bool

	// Seek moves the iterator to the first key/value pair whose key is greater
	// than or equal to the given key.
	// It returns whether such pair exist.
	//
	// It is safe to modify the contents of the argument after Seek returns.
	Seek(key []byte) bool
}
