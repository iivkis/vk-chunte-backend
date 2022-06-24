package util

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// takes a pointer to the structure
// use tag "db" in structs as field name
// return combined insertionfields in format "name = $1, age = $2"
func StructToSQLParams(structure interface{}) (combinedInsertionFields string, insertionValues []interface{}, err error) {
	typeof := reflect.TypeOf(structure).Elem()
	if typeof.Kind() != reflect.Struct {
		return
	}

	valueof := reflect.ValueOf(structure).Elem()

	var (
		insertionFieldIndex = 1
		insertionFields     = []string{}
	)

	for i := 0; i < typeof.NumField(); i++ {
		fieldName, ok := typeof.Field(i).Tag.Lookup("db")
		if !ok {
			continue
		}

		fieldValue := valueof.Field(i)
		if fieldValue.Kind() == reflect.Ptr && fieldValue.IsNil() {
			continue
		}

		insertionFields = append(insertionFields, fieldName+" = $"+strconv.Itoa(insertionFieldIndex))
		insertionFieldIndex++

		insertionValues = append(insertionValues, fieldValue.Interface())
	}

	if len(insertionFields) == 0 {
		return "", insertionValues, fmt.Errorf("there are no inserted fields")
	}

	return strings.Join(insertionFields, ", "), insertionValues, nil
}
