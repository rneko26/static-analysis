package helper

import (
	"fmt"
	"strings"
)

func Foo() {

	rKeySlice := []string{"foo", "bar"}
	rKeyString := strings.Join(rKeySlice, ":")

	fmt.Println(fmt.Sprintf(rKeyString))

}
