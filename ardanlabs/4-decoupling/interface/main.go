package main

type reader interface {
	read(b []byte) (int, error)
}

// file implements the reader interface
type file struct {
	name string
}

func (file) read(b []byte) (int, error) {
	s := "<rss><channel><title>Going Go</title></channel></rss>"
	copy(b, s)
	return len(s), nil
}

// pipe implements the reader interface
type pipe struct {
	name string
}

func (pipe) read(b []byte) (int, error) {
	s := `{name: "Bill", title: "developer"}`
	copy(b, s)
	return len(s), nil
}

func main() {
	// This is odd
	// var r reader

	r := file{"data.json"}
	r.read([]byte{})

	r2 := pipe{"data.json"}
	r2.read([]byte{})

}
