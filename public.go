package main

import (
	"bytes"
	"encoding/json"
	"io"
	pnt "print"
)

// JSON -> struct
func parseJSON(unmsg *[]byte, v interface{}) error {

	dec := json.NewDecoder(bytes.NewReader(*unmsg))
	for {
		if err := dec.Decode(&v); err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}
	return nil
}

// struct -> JSON
func reParseJSON(v interface{}) []byte {
	textbyte, err := json.Marshal(v)
	if err != nil {
		pnt.Errorwd(v, err)
	}
	return textbyte
}
