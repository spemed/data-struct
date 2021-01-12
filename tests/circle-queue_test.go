package tests

import (
	"dataStruct/src/queue"
	"testing"
)

var circleQueue queue.Queue

func init() {
	circleQueue = new(queue.CircleQueueFactory).Create(5)
}

func TestCircleQueueEnQueue(t *testing.T) {
	circleQueue.Enqueue(1)
	circleQueue.Enqueue(2)
	circleQueue.Enqueue(3)
	circleQueue.Enqueue(4)
	err := circleQueue.Enqueue(5)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s", circleQueue)
}

func TestCircleQueueDeQueue(t *testing.T) {
	circleQueue.Enqueue(1)
	circleQueue.Enqueue(2)
	circleQueue.Enqueue(3)
	circleQueue.Enqueue(4)
	t.Log(circleQueue.Dequeue())
	t.Log(circleQueue.Dequeue())
	t.Log(circleQueue.Dequeue())
	t.Log(circleQueue.Dequeue())
	t.Log(circleQueue.Dequeue())
	t.Logf("%s", circleQueue)
}
