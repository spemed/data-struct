package tests

import (
	"dataStruct/src/stack"
	"testing"
)

var linkedStack stack.Stack

func init() {
	linkedStack = new(stack.LinkedListFactory).Create(5)
}

func TestPushLinkedStack(t *testing.T) {
	linkedStack.Push(1)
	linkedStack.Push(2)
	linkedStack.Push(3)
	linkedStack.Push(4)
	linkedStack.Push(5)
	err := linkedStack.Push(7)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s", linkedStack)
}

func TestPopLinkedStack(t *testing.T) {
	linkedStack.Push(1)
	linkedStack.Push(2)
	linkedStack.Push(3)
	linkedStack.Push(4)
	linkedStack.Push(5)
	t.Log(linkedStack.Pop())
	t.Log(linkedStack.Pop())
	t.Log(linkedStack.Pop())
	t.Log(linkedStack.Pop())
	t.Log(linkedStack.Pop())
	t.Log(linkedStack.Pop())
	t.Logf("%s", linkedStack)
}
