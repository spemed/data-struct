package tests

import (
	"dataStruct/src/queue"
	"testing"
)

var linkedQueue queue.Queue

func init() {
	linkedQueue = new(queue.LinkedQueueFactory).Create(0)
}

func TestLinkedQueueEnQueue(t *testing.T) {
	linkedQueue.Enqueue(1)
	linkedQueue.Enqueue(2)
	linkedQueue.Enqueue(3)
	linkedQueue.Enqueue(4)
	err := linkedQueue.Enqueue(5)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s", linkedQueue)
}

func TestLinkedQueueDeQueue(t *testing.T) {
	linkedQueue.Enqueue(1)
	linkedQueue.Enqueue(2)
	linkedQueue.Enqueue(3)
	linkedQueue.Enqueue(4)
	t.Log(linkedQueue.Dequeue())
	t.Log(linkedQueue.Dequeue())
	t.Log(linkedQueue.Dequeue())
	t.Log(linkedQueue.Dequeue())
	t.Log(linkedQueue.Dequeue())
	t.Logf("%s", linkedQueue)
}
