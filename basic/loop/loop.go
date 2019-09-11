package loop

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	var result string
	if n == 0 {
		return "0"
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

const filename  = "./loop.go"

func printFile() {
	// open file
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// new scanner
	scanner := bufio.NewScanner(file)

	// scan
	for scanner.Scan(){
		fmt.Printf("%s\n", scanner.Text())
	}
}
