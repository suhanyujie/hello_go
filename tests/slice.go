package tests

type Stack struct {
	val    []int
	length int
}

func (this *Stack) Len() int {
	return this.length
}

// 压栈
func (this *Stack) Push(value int) {
	this.val = append(this.val, value)
}

// 出栈
func (this *Stack) Pop() (int, error) {
	var popVal int
	len := this.length
	popVal = this.val[len-1]
	this.val = this.val[:len-1]
	return popVal, nil
}

//! 模拟栈
func MockStack() int {
	stack := make([]int, 0)
	// 压栈
	stack = append(stack, 10)
	// 出栈
	val := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return val
}
