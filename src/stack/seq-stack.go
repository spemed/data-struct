package stack

import (
	"errors"

	"fmt"
)

//静态栈
type SeqStack struct {
	Top    int   //栈顶下标
	Length int   //栈的长度
	Data   []int //栈的数据域
}

func (s *SeqStack) String() string {
	data := make([]int, 0)
	for i := s.Top; i >= 0; i-- {
		data = append(data, s.Data[i])
	}
	return fmt.Sprintf("%v", data)
}

type SeqStackFactory struct {
}

//初始化化栈
func (factory *SeqStackFactory) Create(length int) Stack {
	return &SeqStack{
		Top:    -1,
		Length: length,
		Data:   make([]int, length),
	}
}

//入栈
func (s *SeqStack) Push(val int) error {
	if s.Full() {
		return errors.New("full stack")
	}
	s.Top++
	s.Data[s.Top] = val
	return nil
}

//出栈
func (s *SeqStack) Pop() (int, error) {
	if s.Empty() {
		return 0, errors.New("empty stack")
	}
	data := s.Data[s.Top]
	s.Top--
	return data, nil
}

//取得栈顶元素
func (s *SeqStack) GetTop() int {
	data := s.Data[s.Top]
	return data
}

//判断栈是否空
func (s *SeqStack) Empty() bool {
	return s.Top == -1
}

//判断栈是否满
func (s *SeqStack) Full() bool {
	return s.Top+1 == s.Length
}
