package loop

import (
	"fmt"
	"testing"
)

func TestConvertToBin(t *testing.T) {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(0),
	)
}

func TestPrintFile(t *testing.T) {
	printFile()
}

