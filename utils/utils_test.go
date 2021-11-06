package utils_test

import (
	"testing"

	"github.com/d-fal/bverify/utils"
	"github.com/stretchr/testify/assert"
)

func TestUtils(t *testing.T) {
	row := utils.SetRow(struct{ ID int }{ID: 100})

	assert.NotNil(t, row)

	assert.NotNil(t, row.Get())

	assert.NotNil(t, row.String())

}

func TestNilRow(t *testing.T) {
	row := utils.SetRow(nil)

	assert.NotNil(t, row.String())
}
