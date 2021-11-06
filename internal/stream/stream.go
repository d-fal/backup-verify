// streamer is package to stream large json files
// if the input file was too large, this is not wise to
// load the json blocks in memory
// thus, we stream data block by block
package streamming

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type stream struct {
	file *os.File
}

type decoder struct {
	jsonDecoder *json.Decoder
}

type streamer interface {
	open(filename string) (*os.File, error)
	close() error
	Stream() (*decoder, error)
	Get() *stream
}

// NewStreamer would crate a data streamer to open large json files
func NewStreamer(filename string) (streamer, error) {
	var err error
	s := new(stream)

	if s.file, err = s.open(filename); err != nil {
		// file not found
		return nil, fmt.Errorf("cannot open %s", filename)
	}

	return s, nil
}

func (s *stream) Get() *stream {
	return s
}

// open opens the file for reading
func (s *stream) open(filename string) (*os.File, error) {
	return os.Open(filename)
}

func (s *stream) close() error {
	if s != nil && s.file != nil {
		return s.file.Close()
	}
	return fmt.Errorf("file is not open")
}

func (s *stream) reset() error {
	_, err := s.file.Seek(0, io.SeekStart)
	return err
}
func (s *stream) Stream() (*decoder, error) {

	d := newDecoder(s.file)

	if err := d.start(); err != nil {
		return nil, err
	}

	return d, nil

}

// check if a word is inside the file
func (s *stream) Match(block interface{}) (bool, error) {

	if err := s.reset(); err != nil {
		return false, err
	}

	d, err := s.Stream()

	if err != nil {
		return false, nil
	}

	return d.match(block), nil

}
