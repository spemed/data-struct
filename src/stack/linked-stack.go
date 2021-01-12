package stack

import (
	"dataStruct/src/list"
	"errors"
	"fmt"
)

type LinkedListStack struct {
	Top *list.LinkedListNode
}

func (l *LinkedListStack) Push(data int) error {
	newNode := &list.LinkedListNode{
		Data: data,
		Next: nil,
	}
	newNode.Next = l.Top.Next
	l.Top.Next = newNode
	return nil
}

func (l *LinkedListStack) Pop() (int, error) {
	if l.Empty() {
		return 0, errors.New("empty stack")
	}
	data := l.Top.Next.Data
	l.Top.Next = l.Top.Next.Next
	return data, nil
}

func (l *LinkedListStack) GetTop() int {
	return l.Top.Next.Data
}

func (l *LinkedListStack) Empty() bool {
	return l.Top.Next == nil
}

func (l *LinkedListStack) Full() bool {
	return false
}

func (l *LinkedListStack) String() string {
	data := make([]int, 0)
	current := l.Top.Next
	for current != nil {
		data = append(data, current.Data)
		current = current.Next
	}
	return fmt.Sprintf("%v", data)
}

type LinkedListFactory struct {
}

func (factory *LinkedListFactory) Create(int) Stack {
	return &LinkedListStack{Top: &list.LinkedListNode{
		Data: 0,
		Next: nil,
	}}
}
