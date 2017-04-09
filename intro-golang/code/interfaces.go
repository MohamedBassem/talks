package main

import "fmt"

// START OMIT
type Stringer interface {
	String() string
}
type Test struct {
	msg string
}

func (t Test) String() string {
	return t.msg
}
func main() {
	var x Stringer = Test{msg: "Hello World"}
	fmt.Println(x.String())
}

// END OMIT
