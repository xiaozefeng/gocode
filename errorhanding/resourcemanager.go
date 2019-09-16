package errorhanding

import (
	"bufio"
	"fmt"
	"github.com/xiaozefeng/gocode/advance/functional"
	"os"
)

func WriterFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer:= bufio.NewWriter(file)
	defer writer.Flush()

	var f = functional.Fibonacci()
	for i := 0; i < 20; i++ {
		_, _ = fmt.Fprintln(writer, f())
	}
}
