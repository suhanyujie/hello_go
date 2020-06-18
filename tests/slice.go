package tests

// 栈，使用 slice 类型模拟栈
// 相关方法：Pop Push IsEmpty
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

// 栈是否为空
func (_this *Stack) IsEmpty() bool {
	return _this.length == 0
}

// ------------------------

// 队列，使用 slice 模拟队列
// 队列相关方法：Push, Pop, IsEmpty, Len
type Queue struct {
	val    []int
	length int
}

func (_this *Queue) Push(value int) {
	_this.val = append(_this.val, value)
	_this.length += 1
}

func (_this *Queue) Pop() (value int, err error) {
	// 第 0 个先出，使用第 0 个之后的所有值作为新切片
	value = _this.val[0]
	// 切片表示的右侧是开区间（不包含）
	_this.val = _this.val[1:_this.length]
	_this.length -= 1
	return value, nil
}

func (_this *Queue) IsEmpty() bool {
	return _this.length == 0
}
