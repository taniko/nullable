package nullable

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	var nullString Nullable[string]
	assert.True(t, nullString.IsNull())
	assert.Equal(t, "", nullString.Value())
	nullString = New("string")
	assert.False(t, nullString.IsNull())
	assert.Equal(t, "string", nullString.Value())

	var nullInt Nullable[int]
	assert.True(t, nullInt.IsNull())
	assert.Equal(t, 0, nullInt.Value())
	nullInt = New(1)
	assert.False(t, nullInt.IsNull())
	assert.Equal(t, 1, nullInt.Value())

	var nullPointer Nullable[*int]
	n := 1
	assert.True(t, nullPointer.IsNull())
	assert.Nil(t, nullPointer.Value())
	nullPointer = New(&n)
	assert.False(t, nullPointer.IsNull())
	assert.Equal(t, &n, nullPointer.Value())
}

func TestNullable_MarshalJSON(t *testing.T) {
	expected := struct {
		Int    int    `json:"int"`
		String string `json:"string"`
	}{
		Int:    1,
		String: "string",
	}
	expectedBytes, err := json.Marshal(expected)
	assert.NoError(t, err)

	actual := struct {
		Int    Nullable[int]    `json:"int"`
		String Nullable[string] `json:"string"`
	}{
		Int:    New(1),
		String: New("string"),
	}
	actualBytes, err := json.Marshal(actual)
	assert.NoError(t, err)
	assert.Equal(t, expectedBytes, actualBytes)
}

func TestNullable_UnmarshalJSON(t *testing.T) {
	input := []byte(`{"int":1,"string":"string"}`)
	var expected struct {
		Int    int    `json:"int"`
		String string `json:"string"`
	}
	var actual struct {
		Int    Nullable[int]    `json:"int"`
		String Nullable[string] `json:"string"`
	}
	assert.NoError(t, json.Unmarshal(input, &expected))
	assert.NoError(t, json.Unmarshal(input, &actual))
	assert.False(t, actual.Int.IsNull())
	assert.False(t, actual.String.IsNull())
	assert.Equal(t, expected.Int, actual.Int.Value())
	assert.Equal(t, expected.String, actual.String.Value())
}
