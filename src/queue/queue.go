package queue

//队列
type Queue interface {
	GetTop() (int, error)  //获取队首元素
	Enqueue(int) error     //入队
	Dequeue() (int, error) //出队
	String() string
	Full() bool
	Empty() bool
}

type QueueFactory interface {
	Create(int) Queue
}
