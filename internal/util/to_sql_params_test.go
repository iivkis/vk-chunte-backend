package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testStruct struct {
	Name *string `db:"name"`
	Age  *int    `db:"age"`
}

func TestStructToSQLParams(t *testing.T) {

	var (
		name = "ivan"
		age  = int(19)
	)

	test := testStruct{
		Name: &name,
		Age:  &age,
	}

	fields, values := StructToSQLParams(test, 1)
	require.Equal(t, fields, "name = $1, age = $2")
	require.NotNil(t, values[0])
	require.NotNil(t, values[1])
}
