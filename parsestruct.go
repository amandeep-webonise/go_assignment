package main

import (
	"errors"
	"fmt"
	"reflect"
)

func doParse(data reflect.Value) error {

	t := data.Type()

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i).Name
		fv := data.Field(i)
		if fv.Type().Kind() == reflect.Struct {
			fmt.Println("Key {}", f)
			doParse(data.Field(i))
			continue
		}

		fmt.Println("Key", f)
		fmt.Println("Value", fv)

	}
	return nil

}

func Parse(v interface{}) error {
	ptrRef := reflect.ValueOf(v)
	if ptrRef.Kind() != reflect.Struct {
		return errors.New("not a struct ")
	}

	return doParse(ptrRef)
}

func main() {

	err := Parse(struct {
		Name struct {
			First string
		}
		Age int
	}{
		Name: struct{ First string }{
			First: "one",
		},
		Age: 15,
	})

	if err != nil {
		fmt.Println(err)
	}
}
