package stack

import (
	"fmt"
	"errors"
)

type Stack struct {
	Stbf []interface{}
}

func NewStack(size int) (st *Stack) {
	st = &Stack{ Stbf: make([]interface{}, 0, size), }
	return
}

func (st Stack) Pushback(x interface{}) error {
	if len(st.Stbf) == cap(st.Stbf) {
		return errors.New("Stack is full.")
	}
	st.Stbf = append(st.Stbf, 1)
	n := len(st.Stbf)
	st.Stbf[n] = x
	return nil
}

func (st *Stack) Popfront() (interface{}, error) {
	if len(st.Stbf) == 0 {
		return nil, errors.New("Stack is full.")
	}
	x := st.Stbf[0]
	n := len(st.Stbf)
	st.Stbf = st.Stbf[1:n]
	fmt.Println("Stack Len:", n, len(st.Stbf))
	return x, nil
}
