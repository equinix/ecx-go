package ecx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	//given
	input := "test"
	//when
	output := String(input)
	//then
	assert.Equal(t, input, *output, "Input value is same as value of output's pointer")
}

func TestStringValue(t *testing.T) {
	//given
	inputString := "test"
	input := []*string{
		&inputString,
		nil,
	}
	expected := []string{
		inputString,
		"",
	}
	//when
	output := make([]string, len(input))
	for i := range input {
		output[i] = StringValue(input[i])
	}
	//then
	assert.Equal(t, expected, output, "Output matches expected output")
}

func TestInt(t *testing.T) {
	//given
	input := 20
	//when
	output := Int(input)
	//then
	assert.Equal(t, input, *output, "Input value is same as value of output's pointer")
}

func TestIntValue(t *testing.T) {
	//given
	inputInt := 101
	input := []*int{
		&inputInt,
		nil,
	}
	expected := []int{
		inputInt,
		0,
	}
	//when
	output := make([]int, len(input))
	for i := range input {
		output[i] = IntValue(input[i])
	}
	//then
	assert.Equal(t, expected, output, "Output matches expected output")
}

func TestInt64(t *testing.T) {
	//given
	var input int64 = 20
	//when
	output := Int64(input)
	//then
	assert.Equal(t, input, *output, "Input value is same as value of output's pointer")
}

func TestInt64Value(t *testing.T) {
	//given
	var inputInt int64 = 30
	input := []*int64{
		&inputInt,
		nil,
	}
	expected := []int64{
		inputInt,
		0,
	}
	//when
	output := make([]int64, len(input))
	for i := range input {
		output[i] = Int64Value(input[i])
	}
	//then
	assert.Equal(t, expected, output, "Output matches expected output")
}

func TestFloat64(t *testing.T) {
	//given
	var input float64 = 20.55
	//when
	output := Float64(input)
	//then
	assert.Equal(t, input, *output, "Input value is same as value of output's pointer")
}

func TestFloat64Value(t *testing.T) {
	//given
	var inputFloat float64 = 30.69
	input := []*float64{
		&inputFloat,
		nil,
	}
	expected := []float64{
		inputFloat,
		0,
	}
	//when
	output := make([]float64, len(input))
	for i := range input {
		output[i] = Float64Value(input[i])
	}
	//then
	assert.Equal(t, expected, output, "Output matches expected output")
}

func TestBool(t *testing.T) {
	//given
	input := false
	//when
	output := Bool(input)
	//then
	assert.Equal(t, input, *output, "Input value is same as value of output's pointer")
}

func TestBoolValue(t *testing.T) {
	//given
	inputBool := false
	input := []*bool{
		&inputBool,
		nil,
	}
	expected := []bool{
		inputBool,
		false,
	}
	//when
	output := make([]bool, len(input))
	for i := range input {
		output[i] = BoolValue(input[i])
	}
	//then
	assert.Equal(t, expected, output, "Output matches expected output")
}
