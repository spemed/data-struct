package tests

import (
	"dataStruct/src/stack"
	"testing"
)

var seqStack stack.Stack

func init() {
	seqStack = new(stack.SeqStackFactory).Create(5)
}

func TestPushSeqStack(t *testing.T) {
	seqStack.Push(1)
	seqStack.Push(2)
	seqStack.Push(3)
	seqStack.Push(4)
	seqStack.Push(5)
	err := seqStack.Push(7)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s", seqStack)
}

func TestPopSeqStack(t *testing.T) {
	seqStack.Push(1)
	seqStack.Push(2)
	seqStack.Push(3)
	seqStack.Push(4)
	seqStack.Push(5)
	t.Log(seqStack.Pop())
	t.Log(seqStack.Pop())
	t.Log(seqStack.Pop())
	t.Log(seqStack.Pop())
	t.Log(seqStack.Pop())
	t.Log(seqStack.Pop())
	t.Logf("%s", seqStack)
}
