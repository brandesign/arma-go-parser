package parser

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Line struct {
	raw       string
	eventName string
	args      string
}

func (l Line) String() string {
	return fmt.Sprintf("%s %v", l.eventName, l.args)
}

func NewLine(str string) Line {
	spl := strings.SplitN(str, " ", 2)
	l := Line{
		raw:       str,
		eventName: spl[0],
		args:      "",
	}

	if len(spl) > 1 {
		l.args = spl[1]
	}

	return l
}

func (l Line) Scan(dest interface{}) error {
	rf := reflect.ValueOf(dest)
	if rf.Kind() == reflect.Ptr {
		rf = rf.Elem()
	}

	if rf.Kind() != reflect.Struct {
		return fmt.Errorf("can only scan sctructs, got %v", rf.Type())
	}

	nf := rf.NumField()
	if nf < 1 {
		return nil
	}

	args := strings.SplitN(l.args, " ", nf)
	if len(args) < 1 {
		return nil
	}

	for i, arg := range args {
		field := rf.Field(i)
		if err := scanArg(arg, field); err != nil {
			return err
		}
	}
	return nil
}

func scanArg(arg string, refVal reflect.Value) error {
	if !refVal.CanSet() {
		return errors.New("cannot scan arg")
	}

	switch refVal.Kind() {
	case reflect.String:
		refVal.SetString(arg)
	case reflect.Int, reflect.Int32, reflect.Int64:
		v, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			return err
		}
		refVal.SetInt(v)
	case reflect.Float32, reflect.Float64:
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			return err
		}
		refVal.SetFloat(v)
	case reflect.Bool:
		v, err := strconv.ParseBool(arg)
		if err != nil {
			return err
		}
		refVal.SetBool(v)
	case reflect.Struct:
		// TODO: there must be a better way to check this
		if refVal.Type().PkgPath() == "time" && refVal.Type().Name() == "Time" {
			t, err := time.Parse("2006-01-02 15:04:05 MST", arg)
			if err != nil {
				return err
			}

			refVal.Set(reflect.ValueOf(t))
		} else {
			return fmt.Errorf("unsupported field type: %v", refVal.Type())
		}
	case reflect.Slice:
		subArgs := strings.Split(arg, " ")
		sliceRef := reflect.MakeSlice(refVal.Type(), len(subArgs), len(subArgs))
		for i, subArg := range subArgs {
			if err := scanArg(subArg, sliceRef.Index(i)); err != nil {
				return err
			}
		}

		refVal.Set(sliceRef)
	default:
		return fmt.Errorf("unsupported field type: %v", refVal.Type())
	}

	return nil
}
