package regexp

import (
	"fmt"
	"regexp"
	"testing"
)

func TestFindAllString(t *testing.T) {
	r := regexp.MustCompile(`[a-zA-Z0-9]+@[0-9a-zA-Z]+[0-9a-z-A-Z.]+`)
	match := r.FindAllString(text, -1)
	fmt.Println("match:", match)
}

func TestFindAllStringSubMatch(t *testing.T) {
	r := regexp.MustCompile(`([a-zA-Z0-9]+)@([0-9a-zA-Z]+)([0-9a-z-A-Z.]+)`)
	matches := r.FindAllStringSubmatch(text, -1)
	for _, m:= range matches{
		fmt.Println(m)
	}
}
