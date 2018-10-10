package queue

import (
	"fmt"
	"errors"
)

type Queue struct {
	Qbf []interface{}
}

func NewQueue(size int) (qu *Queue) {
	qu = &Queue{ Qbf: make([]interface{}, 0, size), }
	return
}

func (qu Queue) Pushback(x interface{}) error {
	if len(qu.Qbf) == cap(qu.Qbf) {
		return errors.New("Queue is full.")
	}
	qu.Qbf = append(qu.Qbf, 1)
	n := len(qu.Qbf)
	qu.Qbf[n] = x
	return nil
}

func (qu *Queue) Popfront() (interface{}, error) {
	if len(qu.Qbf) == 0 {
		return nil, errors.New("Queue is full.")
	}
	x := qu.Qbf[0]
	n := len(qu.Qbf)
	qu.Qbf = qu.Qbf[1:n]
	fmt.Println("Queue Len:", n, len(qu.Qbf))
	return x, nil
}
