package tests

import (
	"testing"
	"time"
)

//! 测试 Stack
func TestMockStack_len(t *testing.T) {
	arr := make([]int, 2)
	s1 := Stack{
		val:    arr,
		length: 2,
	}
	len := s1.Len()
	if len != 2 {
		t.Errorf("the length is %d\n", len)
	}
	v1, _ := s1.Pop()
	t.Logf("the val is %d \n", v1)
	time.Sleep(time.Duration(2) * time.Second)
}

//! 测试 Queue
func TestQueue(t *testing.T) {
	q1 := Queue{
		val:    []int{1, 2, 3},
		length: 3,
	}
	v1, err := q1.Pop()
	if v1 != 1 || err != nil {
		t.Errorf("the value poped is incorrect 1")
	}
	v1, err = q1.Pop()
	if v1 != 2 || err != nil {
		t.Errorf("the value poped is incorrect 2")
	}
	v1, err = q1.Pop()
	if v1 != 3 || err != nil {
		t.Errorf("the value poped is incorrect 3")
	}
	q1.Push(100254)
	v1, err = q1.Pop()
	if v1 != 100254 || err != nil {
		t.Errorf("the value poped is incorrect 100254")
	}
	q1.Push(12)
	q1.Push(22)
	t.Logf("the val is %v \n", q1)
	t.Logf("the val is %v \n", v1)
}
