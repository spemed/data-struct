package queue

import (
	"errors"
	"fmt"
)

type CircleQueue struct {
	Data  []int //底层存储数组
	Front int   //队首
	Tail  int   //队尾
}

func (c *CircleQueue) Full() bool {
	return (c.Tail+1)%len(c.Data) == c.Front
}

func (c *CircleQueue) Empty() bool {
	return c.Front == c.Tail
}

func (c *CircleQueue) GetTop() (int, error) {
	if c.Empty() {
		return 0, errors.New("empty queue")
	}
	data := c.Data[c.Tail]
	return data, nil
}

func (c *CircleQueue) Enqueue(val int) error {
	if c.Full() {
		return errors.New("full queue")
	}
	c.Data[c.Tail] = val
	c.Tail = (c.Tail + 1) % len(c.Data)
	return nil
}

func (c *CircleQueue) Dequeue() (int, error) {
	if c.Empty() {
		return 0, errors.New("empty queue")
	}
	data := c.Data[c.Front]
	c.Front = (c.Front + 1) % len(c.Data)
	return data, nil
}

func (c *CircleQueue) String() string {
	if c.Empty() {
		return "[]"
	}

	data := make([]int, 0)
	front, tail := c.Front, c.Tail
	//队列未满
	for {
		if front%len(c.Data) == tail {
			break
		}
		data = append(data, c.Data[front])
		front = (front + 1) % len(c.Data)
	}
	return fmt.Sprintf("%v", data)
}

type CircleQueueFactory struct {
}

func (factory *CircleQueueFactory) Create(length int) Queue {
	return &CircleQueue{
		Data:  make([]int, length),
		Front: 0,
		Tail:  0,
	}
}
