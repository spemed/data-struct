package queue

import (
	"dataStruct/src/list"
	"errors"
	"fmt"
)

type LinkedQueue struct {
	Front *list.LinkedListNode //队头
	Tail  *list.LinkedListNode //队尾
}

func (l *LinkedQueue) GetTop() (int, error) {
	if l.Empty() {
		return 0, errors.New("empty queue")
	}
	return l.Front.Data, nil
}

func (l *LinkedQueue) Enqueue(i int) error {
	node := &list.LinkedListNode{
		Data: i,
		Next: nil,
	}
	l.Tail.Next = node
	l.Tail = node
	return nil
}

func (l *LinkedQueue) Dequeue() (int, error) {
	if l.Empty() {
		return 0, errors.New("full queue")
	}
	node := l.Front.Next
	l.Front.Next = node.Next
	return node.Data, nil
}

func (l *LinkedQueue) String() string {
	current := l.Front.Next
	data := make([]int, 0)
	for current != nil {
		data = append(data, current.Data)
		current = current.Next
	}
	return fmt.Sprintf("%v", data)
}

func (l *LinkedQueue) Full() bool {
	return false
}

func (l *LinkedQueue) Empty() bool {
	return l.Tail.Next == nil
}

type LinkedQueueFactory struct {
}

func (factory *LinkedQueueFactory) Create(int) Queue {
	node := &list.LinkedListNode{
		Data: 0,
		Next: nil,
	}
	return &LinkedQueue{
		Front: node,
		Tail:  node,
	}
}
