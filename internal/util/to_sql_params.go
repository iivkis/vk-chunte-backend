package util

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var ErrMustBePointerToStruct = errors.New("the argument must be a pointer to the structure")

// takes a pointer to the structure
// use tag "db" in structs as field name
// return combined insertionfields in format "name = $1, age = $2"
func StructToInsertableSQLParams(structure interface{}) (combinedInsertionFields string, insertionValues []interface{}, err error) {
	typeof, valueof, ok := getStructTypeOfAndValueOf(structure)
	if !ok {
		return combinedInsertionFields, insertionValues, ErrMustBePointerToStruct
	}

	var (
		insertionFieldIndex = 1
		insertionFields     = []string{}
	)

	for i := 0; i < typeof.NumField(); i++ {
		fieldNameFromTag, ex := typeof.Field(i).Tag.Lookup("db")
		if !ex {
			continue
		}

		fieldValue := valueof.Field(i)
		if fieldValue.Kind() == reflect.Ptr && fieldValue.IsNil() {
			continue
		}

		insertionFields = append(insertionFields, fieldNameFromTag+" = $"+strconv.Itoa(insertionFieldIndex))
		insertionValues = append(insertionValues, fieldValue.Interface())

		insertionFieldIndex++
	}

	if len(insertionFields) == 0 {
		return "", insertionValues, fmt.Errorf("there are no inserted fields")
	}

	combinedInsertionFields = strings.Join(insertionFields, ", ")

	return
}

func getStructTypeOfAndValueOf(s interface{}) (typeof reflect.Type, valueof reflect.Value, ok bool) {
	typeof = reflect.TypeOf(s)
	if typeof.Kind() != reflect.Ptr {
		return
	}

	typeof = typeof.Elem()
	if typeof.Kind() != reflect.Struct {
		return
	}

	valueof = reflect.ValueOf(s).Elem()

	return typeof, valueof, true
}
