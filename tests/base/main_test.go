package base

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddone(t *testing.T) {
	// var (
	// 	input  = 1
	// 	output = 2
	// )
	// actual := AddOne(1)
	// if actual != output {
	// 	t.Error("AddOne(%d),input %d, actual = %d ", input, output, actual)
	// }

	assert.Equal(t, AddOne(2), 4, "AddOne(2) shuild be 3")
	assert.NotEqual(t, 2, 3)
	assert.Nil(t, nil, nil)
}

func TestREquire(t *testing.T) {
	fmt.Println("Not executing")

}

func TestAssert(t *testing.T) {
	assert.Equal(t, 2, 3)
	fmt.Println("executing")
}
