package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewLine(t *testing.T) {
	l := NewLine("TEST")

	assert.Equal(t, "TEST", l.eventName)
	assert.Equal(t, "", l.args)

	l = NewLine("TEST arg1 arg2")
	assert.Equal(t, "TEST", l.eventName)
	assert.Equal(t, "arg1 arg2", l.args)
}

func TestScan_Empty(t *testing.T) {
	l := NewLine("TEST")
	evt := &struct{}{}
	err := l.Scan(evt)

	assert.Nil(t, err)
}

func TestScan_String(t *testing.T) {
	l := NewLine("TEST someString")
	evt := &struct {
		Arg string
	}{}
	err := l.Scan(evt)

	assert.Nil(t, err)
	assert.Equal(t, "someString", evt.Arg)
}

func TestScan_Int(t *testing.T) {
	l := NewLine("TEST 666")
	evt := &struct {
		Arg int
	}{}
	err := l.Scan(evt)

	assert.Nil(t, err)
	assert.Equal(t, 666, evt.Arg)
}

func TestScan_Float(t *testing.T) {
	l := NewLine("TEST 1.23456789")
	evt := &struct {
		Arg float64
	}{}
	err := l.Scan(evt)

	assert.Nil(t, err)
	assert.Equal(t, 1.23456789, evt.Arg)
}

func TestScan_Bool_true(t *testing.T) {
	l := NewLine("TEST true")
	evt := &struct {
		Arg bool
	}{}
	err := l.Scan(evt)

	assert.Nil(t, err)
	assert.Equal(t, true, evt.Arg)
}

func TestScan_Bool_false(t *testing.T) {
	l := NewLine("TEST false")
	evt := &struct {
		Arg bool
	}{}
	err := l.Scan(evt)

	assert.Nil(t, err)
	assert.Equal(t, false, evt.Arg)
}

func TestScan_Time(t *testing.T) {
	l := NewLine("TEST 2021-01-30 00:58:05 GMT")
	evt := &struct {
		Arg time.Time
	}{}
	err := l.Scan(evt)

	assert.Nil(t, err)

	want := time.Date(2021, 01, 30, 00, 58, 05, 0, time.UTC)
	assert.Equal(t, want.UnixNano(), evt.Arg.UnixNano())
}

func TestScan_StringSlice(t *testing.T) {
	l := NewLine("TEST arg1 arg2")
	evt := &struct {
		Arg []string
	}{}
	err := l.Scan(evt)

	assert.Nil(t, err)
	assert.Equal(t, []string{"arg1", "arg2"}, evt.Arg)
}

func TestScan_IntSlice(t *testing.T) {
	l := NewLine("TEST 111 222 333")
	evt := &struct {
		Arg []int
	}{}
	err := l.Scan(evt)

	assert.Nil(t, err)
	assert.Equal(t, []int{111, 222, 333}, evt.Arg)
}

func TestScan_FloatSlice(t *testing.T) {
	l := NewLine("TEST 1.2 2.3 3.4")
	evt := &struct {
		Arg []float64
	}{}
	err := l.Scan(evt)

	assert.Nil(t, err)
	assert.Equal(t, []float64{1.2, 2.3, 3.4}, evt.Arg)
}

func TestScan_BoolSlice(t *testing.T) {
	l := NewLine("TEST true false true")
	evt := &struct {
		Arg []bool
	}{}
	err := l.Scan(evt)

	assert.Nil(t, err)
	assert.Equal(t, []bool{true, false, true}, evt.Arg)
}
