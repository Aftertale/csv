package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func marshalHeader(vt reflect.Type) []string {
	var out []string
	for i := 0; i < vt.NumField(); i++ {
		field := vt.Field(i)
		if curTag, ok := field.Tag.Lookup("csv"); ok {
			out = append(out, curTag)
		}
	}
	return out
}

func marshalOne(vv reflect.Value) ([]string, error) {
	var row []string
	vt := vv.Type()
	for i := 0; i < vt.NumField(); i++ {
		fieldVal := vv.Field(i)
		if _, ok := vt.Field(i).Tag.Lookup("csv"); !ok {
			continue
		}
		switch fieldVal.Kind() {
		case reflect.String:
			row = append(row, fieldVal.String())
		case reflect.Int:
			row = append(row, strconv.Itoa(int(fieldVal.Int())))
		case reflect.Bool:
			row = append(row, strconv.FormatBool(fieldVal.Bool()))
		default:
			continue
		}
	}
	return row, nil
}

func Marshal(i interface{}) ([][]string, error) {
	sliceVal := reflect.ValueOf(i)
	if sliceVal.Kind() != reflect.Slice {
		return nil, fmt.Errorf("must be a slice")
	}

	sliceStruct := sliceVal.Type().Elem()
	if sliceStruct.Kind() != reflect.Struct {
		return nil, fmt.Errorf("must be a slice of structs")
	}

	var csv [][]string
	header := marshalHeader(sliceStruct)
	csv = append(csv, header)
	for i := 0; i < sliceVal.Len(); i++ {
		elem, err := marshalOne(sliceVal.Index(i))
		if err != nil {
			return nil, fmt.Errorf("error while marshalling %+v", sliceVal.Index(i))
		}
		csv = append(csv, elem)
	}

	return csv, nil
}
