package list

import (
	"errors"
	"fmt"
)

type LinkedListNode struct {
	Data int             //数据域
	Next *LinkedListNode //指针域
}

type LinkedList struct {
	Head *LinkedListNode // 头指针,充当哨兵节点
}

//构造链表
func NewLinkedList() *LinkedList {
	head := &LinkedListNode{
		Data: 0,
		Next: nil,
	}
	return &LinkedList{Head: head}
}

//在链表指定的节点后插入
func (l *LinkedList) Insert(index int, val int) {
	newNode := &LinkedListNode{
		Data: val,
		Next: nil,
	}
	current := l.FindIndex(index)
	newNode.Next = current.Next
	current.Next = newNode
}

//删除链表的指定节点
func (l *LinkedList) Remove(index int) (int, error) {
	if index <= 0 {
		return 0, errors.New("param index must be positive integer")
	}
	current := l.Head
	for i := 1; i < index; i++ {
		if current.Next != nil {
			current = current.Next //指针迭代
		} else {
			return 0, errors.New("param index bigger than linkedList length")
		}
	}
	data := current.Next.Data
	current.Next = current.Next.Next
	return data, nil
}

//寻找到目标节点并返回指针
func (l *LinkedList) FindIndex(index int) *LinkedListNode {
	current := l.Head
	//仅有头结点
	if current.Next == nil {
		return l.Head
	}
	for i := 0; i < index; i++ {
		if current.Next != nil {
			current = current.Next //指针迭代
		} else {
			break
		}
	}
	return current
}

/**

 */
func (l *LinkedList) String() string {
	var data []int
	current := l.Head.Next
	for current != nil {
		data = append(data, current.Data)
		current = current.Next
	}
	return fmt.Sprintf("%v", data)
}

//删除值为val的节点
func (l *LinkedList) RemoveVal(data int) error {
	current := l.FindVal(data)
	if current.Next == nil {
		return errors.New("node not found")
	}
	current.Next = current.Next.Next
	return nil
}

//查询值为val的节点
func (l *LinkedList) FindVal(val int) *LinkedListNode {
	current := l.Head
	for current.Next != nil {
		if current.Next.Data == val {
			break
		}
		current = current.Next
	}
	return current
}

//链表尾部追加
func (l *LinkedList) Append(val int) {
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	newNode := &LinkedListNode{
		Data: val,
		Next: nil,
	}
	current.Next = newNode
}

//链表尾部删除
func (l *LinkedList) Delete() (int, error) {
	current := l.Head
	if current.Next == nil {
		return 0, errors.New("empty linked list could not be delete")
	}
	for current.Next != nil {
		if current.Next.Next != nil {
			current = current.Next
		} else {
			break
		}
	}
	data := current.Next.Data
	current.Next = current.Next.Next
	return data, nil
}
