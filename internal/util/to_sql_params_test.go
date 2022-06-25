package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type testStruct struct {
	ID   int     `db:"id"`
	Name *string `db:"name"`
	Age  *int    `db:"age"`
}

func TestStructToSQLParams(t *testing.T) {
	//no err
	var (
		name = "ivan"
		age  = int(19)
	)

	test := testStruct{
		ID:   GetRandomInt(1, 999),
		Name: &name,
		Age:  &age,
	}

	fields, values, err := StructToInsertableSQLParams(&test)
	require.NoError(t, err)
	require.Equal(t, fields, "id = $1, name = $2, age = $3")
	require.Equal(t, values[0], test.ID)
	require.Equal(t, *values[1].(*string), name)
	require.Equal(t, *values[2].(*int), age)

	fmt.Println(values)

	//with err
	_, _, err = StructToInsertableSQLParams(&struct{}{})
	require.Error(t, err)
}
