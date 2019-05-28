package main

import "fmt"

func zeroval(i int) {
	i = 0
}
func zeroptr(iptr *int) {
	*iptr = 0
}
func main() {
	i := 5
	fmt.Println("Print i val", i)

	zeroval(i)
	fmt.Println("call zeroval function", i)

	zeroptr(&i)
	fmt.Println("call zeroptr function", i)

}
