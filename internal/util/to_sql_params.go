package util

import (
	"reflect"
	"strconv"
	"strings"
)

//use tag "db" in structs
func StructToSQLParams(structure interface{}, initialInsertionIndex int) (combinedInsertionFields string, insertionValues []interface{}) {
	typeof := reflect.TypeOf(structure)

	if typeof.Kind() != reflect.Struct {
		return
	}

	valueof := reflect.ValueOf(structure)

	var insertionFields []string
	for i := 0; i < typeof.NumField(); i++ {
		nameInDB, ok := typeof.Field(i).Tag.Lookup("db")
		if !ok {
			continue
		}

		fieldValue := valueof.Field(i)
		if fieldValue.IsNil() {
			continue
		}

		insertionFields = append(insertionFields, nameInDB+" = $"+strconv.Itoa(initialInsertionIndex))
		initialInsertionIndex++

		insertionValues = append(insertionValues, fieldValue.Interface())
	}

	return strings.Join(insertionFields, ", "), insertionValues
}
