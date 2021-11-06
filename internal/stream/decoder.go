package streamming

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/google/go-cmp/cmp"
)

func newDecoder(file io.Reader) *decoder {
	return &decoder{jsonDecoder: json.NewDecoder(file)}
}

func (d *decoder) start() error {
	if _, err := d.jsonDecoder.Token(); err != nil {
		return fmt.Errorf("Sub/Nonstandard json file:  %v", err)
	}
	return nil
}

func (d *decoder) Next() (interface{}, bool) {
	var word interface{}

	if d.jsonDecoder.More() {
		if err := d.jsonDecoder.Decode(&word); err != nil {
			return nil, false
		}
		return word, true
	}
	return nil, false
}

func (d *decoder) Close() error {

	if _, err := d.jsonDecoder.Token(); err != nil {
		return fmt.Errorf("no ending delimiter found: %v", err)
	}
	return nil
}

func (d *decoder) match(block interface{}) bool {
	for {
		b, ok := d.Next()

		if !ok {
			break
		}

		if cmp.Equal(b, block) {
			return true
		}
	}
	return false
}
