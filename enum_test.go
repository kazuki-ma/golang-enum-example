package main

//go:generate bash -c "thrift --gen go *.thrift"

import (
	"fmt"
	"testing"

	"github.com/assertgo/assert"
	"github.com/kazuki-ma/golang-enum-example/gen-go/enum"
)

func debug(v interface{}) {
	fmt.Printf("%T %#v %s\n", v, v, v)
}

func Test_Enum(t *testing.T) {
	assert := assert.New(t)

	// Verify
	assert.
		That(enum.Gender_Female).
		IsEqualTo(enum.Gender_Female)
}

func Test_EnumFromstring(t *testing.T) {
	assert := assert.New(t)

	// Do
	v, err := enum.GenderFromString("Female")

	// Verify
	assert.
		That(err).
		IsNil()
	assert.
		That(v).
		IsEqualTo(enum.Gender_Female)
}

func Test_EnumToValue(t *testing.T) {
	assert := assert.New(t)

	// Do
	v, err := enum.GenderPtr(enum.Gender_Female).Value()

	// Verify
	assert.
		That(err).
		IsNil()
	assert.
		That(v).
		IsEqualTo(int64(2))
}
