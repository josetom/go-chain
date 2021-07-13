package filedb

import (
	"bufio"
	"bytes"
	"encoding/json"
	"log"
	"os"

	"github.com/josetom/go-chain/db/types"
)

type FileDbIterator struct {
	db *os.File

	start []byte
	limit []byte

	scanner   *bufio.Scanner
	hasSeeked bool

	key   []byte
	value []byte
}

func newFileDbIterator(dbPath string, start []byte, limit []byte) types.Iterator {
	d, err := os.OpenFile(dbPath, os.O_RDONLY, 0600)
	scanner := bufio.NewScanner(d)
	if err != nil {
		log.Fatalln("DB File cannot be opened !") // This should not happen since FileDb should not have gotten created in the first place
	}
	return &FileDbIterator{
		db:      d,
		scanner: scanner,
		start:   start,
		limit:   limit,
	}
}

func (f *FileDbIterator) Key() []byte {
	return f.key
}

func (f *FileDbIterator) Value() []byte {
	return f.value
}

func (f *FileDbIterator) Next() bool {
	if !f.hasSeeked {
		return f.Seek(f.start)
	} else {
		// move pointer
		isScanSuccess, record := f.scan()
		if !isScanSuccess {
			return false
		}
		f.key = record.Key
		f.value = record.Value
		return true
	}
}

func (f *FileDbIterator) Seek(start []byte) bool {
	for isScanSuccess, record := f.scan(); isScanSuccess; {
		if start == nil || bytes.Equal(record.Key, start) {
			f.key = record.Key
			f.value = record.Value
			f.hasSeeked = true
			return true
		}
	}

	return false
}

func (f *FileDbIterator) scan() (bool, types.Record) {
	if !f.scanner.Scan() {
		return false, types.Record{}
	}

	if err := f.scanner.Err(); err != nil {
		return false, types.Record{}
	}

	var record types.Record
	log.Println(string(f.scanner.Bytes()))
	err := json.Unmarshal(f.scanner.Bytes(), &record)
	log.Println(record)
	if err != nil {
		return false, types.Record{}
	}

	return true, record
}
