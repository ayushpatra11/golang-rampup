package stdlib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// struct tags control JSON serialization behavior
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"` // omit if empty string
	Password string `json:"-"`               // always excluded
	Age      int    `json:"age,omitempty"`   // omit if 0
}

// ---- Marshal: Go value -> JSON bytes ----

func marshalDemo() {
	u := User{ID: 1, Name: "Alice", Email: "alice@example.com", Password: "secret"}
	data, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	// {"id":1,"name":"Alice","email":"alice@example.com"}
	// Password omitted by "-", Age omitted by omitempty+zero

	// pretty print
	pretty, _ := json.MarshalIndent(u, "", "  ")
	fmt.Println(string(pretty))
}

// ---- Unmarshal: JSON bytes -> Go value ----

func unmarshalDemo() {
	data := []byte(`{"id":2,"name":"Bob","age":25}`)
	var u User
	if err := json.Unmarshal(data, &u); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", u)  // {ID:2 Name:Bob Email: Password: Age:25}
}

// unknown fields are silently ignored by default
// unknown JSON keys that don't map to a struct field are dropped

// ---- dynamic JSON with map[string]any ----

func dynamicJSON() {
	data := []byte(`{"event":"click","x":42,"active":true}`)
	var m map[string]any
	json.Unmarshal(data, &m)

	// type assertion needed for concrete values
	if x, ok := m["x"].(float64); ok { // JSON numbers decode to float64 in any
		fmt.Println("x:", x)
	}
}

// ---- Encoder/Decoder: stream JSON to/from io.Reader/Writer ----

func encodeToWriter(w io.Writer, users []User) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(users) // appends \n automatically
}

func decodeFromReader(r io.Reader) ([]User, error) {
	var users []User
	dec := json.NewDecoder(r)
	if err := dec.Decode(&users); err != nil {
		return nil, err
	}
	return users, nil
}

// streaming multiple JSON objects from a single reader
func decodeStream(r io.Reader) {
	dec := json.NewDecoder(r)
	for {
		var u User
		err := dec.Decode(&u)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "decode error:", err)
			continue
		}
		fmt.Println(u.Name)
	}
}

// ---- custom marshal/unmarshal ----
// implement json.Marshaler / json.Unmarshaler

type Color struct{ R, G, B uint8 }

func (c Color) MarshalJSON() ([]byte, error) {
	hex := fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
	return json.Marshal(hex)
}

func (c *Color) UnmarshalJSON(data []byte) error {
	var hex string
	if err := json.Unmarshal(data, &hex); err != nil {
		return err
	}
	_, err := fmt.Sscanf(hex, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	return err
}

// ---- RawMessage: delay/skip decoding of part of a JSON doc ----

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

func rawMessageDemo() {
	data := []byte(`{"type":"click","payload":{"x":10,"y":20}}`)
	var e Event
	json.Unmarshal(data, &e)

	// decode payload later based on Type
	var coords struct{ X, Y int }
	json.Unmarshal(e.Payload, &coords)
	fmt.Println(e.Type, coords) // click {10 20}
}

func roundTrip(v any) (string, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		return "", err
	}
	return buf.String(), nil
}

/*
Struct tag cheat sheet:
  json:"name"            rename field
  json:"name,omitempty"  omit zero values (nil, 0, "", false, empty slice/map)
  json:"-"               always skip
  json:",string"         marshal number/bool as JSON string

json.Number type: use when you need the raw number string from dynamic JSON.
*/
