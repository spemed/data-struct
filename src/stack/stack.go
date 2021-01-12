package stack

type Stack interface {
	Push(int) error
	Pop() (int, error)
	GetTop() int
	Empty() bool
	Full() bool
	String() string
}

//工厂方法
type StackFactory interface {
	Create(int) Stack
}
