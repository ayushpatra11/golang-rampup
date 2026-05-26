package stdlib

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"strings"
)

// ---- encoding/csv ----

func csvWriteDemo() {
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)

	records := [][]string{
		{"name", "age", "city"},
		{"Alice", "30", "New York"},
		{"Bob", "25", "London"},
		{"Charlie", "35", "Tokyo"},
	}

	for _, record := range records {
		w.Write(record)
	}
	w.Flush() // must flush - writes are buffered

	if err := w.Error(); err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
}

func csvReadDemo(input string) ([][]string, error) {
	r := csv.NewReader(strings.NewReader(input))

	// optional settings
	r.Comment = '#'         // lines starting with # are ignored
	r.FieldsPerRecord = -1  // allow variable number of fields per row (-1 = no check)
	r.TrimLeadingSpace = true

	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("csv read: %w", err)
	}
	return records, nil
}

// read CSV row by row (better for large files)
func csvStreamRead(input string) {
	r := csv.NewReader(strings.NewReader(input))

	// skip header
	if _, err := r.Read(); err != nil {
		return
	}

	for {
		record, err := r.Read()
		if err != nil {
			break // io.EOF or parse error
		}
		fmt.Println(record)
	}
}

// ---- encoding/base64 ----
// base64: encode binary data as ASCII text (e.g. for HTTP, email, JSON)

func base64Demo() {
	data := []byte("Hello, World! \x00\x01\x02")

	// standard encoding (uses + and /)
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(decoded))

	// URL-safe encoding (uses - and _ instead of + and /)
	// use for base64 in URLs, filenames, JWT tokens
	urlEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println(urlEncoded)

	// raw encoding: no padding (= characters removed)
	// RawStdEncoding and RawURLEncoding
	rawEncoded := base64.RawURLEncoding.EncodeToString(data)
	fmt.Println(rawEncoded)

	// streaming encoder (for large data)
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write(data)
	enc.Close() // flushes and adds padding
	fmt.Println(buf.String())
}

// ---- encoding/hex ----
// hex: encode bytes as hexadecimal string (2 chars per byte)

func hexDemo() {
	data := []byte{0xDE, 0xAD, 0xBE, 0xEF}

	// encode
	s := hex.EncodeToString(data) // "deadbeef"
	fmt.Println(s)

	// decode
	decoded, err := hex.DecodeString("48656c6c6f")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(decoded)) // "Hello"

	// Dump: debug-friendly hex+ASCII view (like hexdump -C)
	fmt.Println(hex.Dump([]byte("Hello, Go!")))
}

/*
Encoding cheat sheet:
  encoding/json    Go values <-> JSON
  encoding/csv     tabular text data
  encoding/base64  binary <-> ASCII-safe text
  encoding/hex     binary <-> hex string
  encoding/binary  Go values <-> raw bytes (little/big endian)
  encoding/gob     Go-specific binary serialization (fast, not cross-language)
  encoding/xml     Go values <-> XML
*/
