package main

import (
	"fmt"

	gomod5a_v1 "github.com/relax-space/gomod5a"
	// gomod5a_v2 "github.com/relax-space/gomod5a/v2"
)

func main() {
	fmt.Printf("gomod5a_v1.AddA(1,2)=%d", gomod5a_v1.AddA(1, 2))
	// fmt.Printf("gomod5a_v2.AddA(1,2)=%d", gomod5a_v2.AddA(1, 2))
}
