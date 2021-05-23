package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
)

func ParseJSONBody(source io.ReadCloser, target interface{}) error {
	// un-serialize the origin source.
	dec := json.NewDecoder(source)

	// prevent unknown field appear in decode result.
	dec.DisallowUnknownFields()

	// put the decoded result to target(reference).
	err := dec.Decode(&target)
	if err != nil {
		return err
	}

	// verify the data is clear.
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return err
	}

	return nil
}

func ReaderToString(source io.Reader) (string, error) {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(source); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func StringToReadCloser(source string) io.ReadCloser {
	return ioutil.NopCloser(strings.NewReader(source))
}
