package main

import (
	"fmt"
	"queue"
	"stack"
)

func main() {
	qu := NewQueue(10)
	qu.Pushback("ABC")
	qu.Pushback("DEF")
	for i := 0; i < len(qu.Qbf); i++ {
		x := qu.Popfront()
		fmt.Println(x.(string))
	}
	fmt.Println("queue.. done")

	st := NewStack(10)

	st.Push("ABC")
	st.Push("DEF")
	for i := 0; i < len(st.Sbf); i++ {
		x := st.Pop()
		fmt.Println(x.(string))

	fmt.Println("stack.. done")
}
