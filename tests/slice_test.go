package tests

import (
	"testing"
	"time"
)

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
