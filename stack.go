package stack

// 定义一个泛型的Stack结构体
type Stack[T any] struct {
	Items []T
	Top   int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		Items: make([]T, 0),
		Top:   -1,
	}
}

// 实现Push方法来添加元素
func (s *Stack[T]) Push(item T) {
	s.Items = append(s.Items, item)
	s.Top++
}

// 实现Pop方法来移除并返回栈顶元素
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.Items) == 0 {
		var zero T // 如果栈为空，返回类型的零值(初始化默认零值)和false
		return zero, false
	}
	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	s.Top--
	return item, true
}

// 实现Peek方法来查看栈顶元素但不移除它
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.Items) == 0 {
		var zero T // 如果栈为空，返回类型的零值和false
		return zero, false
	}
	return s.Items[len(s.Items)-1], true
}

// 实现IsEmpty方法来检查栈是否为空
func (s *Stack[T]) IsEmpty() bool {
	return len(s.Items) == 0
}
