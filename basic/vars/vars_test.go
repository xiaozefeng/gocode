package vars

import (
	"fmt"
	"testing"
)


func Test(t *testing.T) {

	// zero value
	fmt.Println("variable zero value:")
	variableZeroValue()
	fmt.Println()

	fmt.Println("variable initial value:")
	variableInitialValue()
	fmt.Println()

	fmt.Println("variable type deduction:")
	variableTypeDeduction()
	fmt.Println()

	fmt.Println("variable shorter")
	variableShorter()
	fmt.Println()

	fmt.Println("package variable")
	packageVariable()
	fmt.Println()
}


