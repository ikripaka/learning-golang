package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

// Marshal is a function that marshals the object into an io.Reader.
// By default, it uses the JSON marshaller.
var Marshal = func(v *interface{})(io.Reader, error){
	b,err := json.MarshalIndent(v, "","\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

// Unmarshal is a function that unmarshals the data from the
// reader into the specified value.
// By default, it uses the JSON unmarshaller.
var Unmarshal = func(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// Save saves a representation of v to the file at path.
func Save(path string, s *server, v interface{}) error{
	s.mutex.Lock()
	defer s.mutex.Unlock()
	f,err := os.Create(path)
	if err!=nil{
		return err
	}

	defer f.Close()
	r, err := Marshal(&v)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err

	return nil
}

func Load(path string, s *server,v interface{}) error{
	s.mutex.Lock()
	defer s.mutex.Unlock()
	f, err := os.Open(path)
	if err != nil{
		return err
	}
	defer f.Close()

	return Unmarshal(f, v)
}
