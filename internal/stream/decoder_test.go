package streamming

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/logrusorgru/aurora"
	"github.com/stretchr/testify/assert"
)

var sampleJson = `
	[
		{
			"id": "1"
		},
		{
			"id" : "2",
			"name": "someone"
		}
	]
`

func TestDecoder(t *testing.T) {
	data := []byte(sampleJson)

	reader := bytes.NewReader(data)

	decoder := newDecoder(reader)

	if err := decoder.start(); err != nil {

		assert.NoError(t, err)
		return
	}
	defer decoder.Close()
	for {
		block, ok := decoder.Next()
		if !ok {
			break
		}
		fmt.Println("block : ", aurora.Yellow(block))
	}

}

func TestSearchInFunc(t *testing.T) {
	data := []byte(sampleJson)

	// check for match
	sample := struct {
		ID string `json:"id"`
	}{ID: "1"}

	reader := bytes.NewReader(data)
	decoder := newDecoder(reader)

	if err := decoder.start(); err != nil {

		assert.NoError(t, err)
		return
	}
	defer decoder.Close()
	exists := decoder.match(sample)

	assert.Equal(t, exists, false)
}

func TestStreamming(t *testing.T) {
	streamming, err := NewStreamer("../../sample/original.json")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	decoder, err := streamming.Stream()

	assert.NoError(t, err)
	for {
		_, ok := decoder.Next()
		if !ok {
			break
		}
	}
}
