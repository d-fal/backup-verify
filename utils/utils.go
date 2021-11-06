package utils

import (
	"encoding/json"
)

type rowI interface {
	String() string
	Get() *Row
}

type Row struct {
	src interface{}
}

func SetRow(src interface{}) rowI {
	return &Row{
		src: src,
	}
}

func (r *Row) Get() *Row {
	return r
}
func (r *Row) String() string {
	if r != nil {
		src, _ := json.Marshal(r.src)
		return string(src)
	}
	return ""
}
