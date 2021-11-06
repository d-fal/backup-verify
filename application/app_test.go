package application_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/d-fal/bverify/application"
	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	diffFilePath := "_sample.txt"
	scenarios := []struct {
		name  string
		dst   string
		src   string
		equal bool
		err   error
	}{
		{
			name:  "comparing similar files",
			src:   "../sample/original.json",
			dst:   "../sample/original.json",
			equal: true,
			err:   nil,
		},
		{
			name:  "comparing different files",
			src:   "../sample/original.json",
			dst:   "../sample/duplicate.json",
			equal: false,
			err:   nil,
		},
		{
			name:  "nonexisting files",
			src:   "sample",
			dst:   "sample",
			equal: false,
			err:   fmt.Errorf("cannot open %s", "sample"),
		},
	}
	for _, cnf := range scenarios {
		t.Run(cnf.name, func(t *testing.T) {
			matched, err := application.Start(cnf.src, cnf.dst,
				application.WithConsolePrint(),
				application.WithDiffSave(diffFilePath))
			assert.Equal(t, err, cnf.err)
			assert.Equal(t, matched, cnf.equal)

		})
	}

	// remove the stored file
	err := os.Remove(diffFilePath)
	assert.Equal(t, err, nil)
}
