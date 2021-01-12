package list

import (
	"errors"
	"fmt"
)

/*
	go实现数据结构之,顺序表
*/
type SequentialList struct {
	Length   int   //顺序表当前长度
	Capacity int   //顺序表的容积
	Data     []int //用于存储真实数据的内部数组
}

//初始化顺序表
func NewSequentialList(capacity int) *SequentialList {
	return &SequentialList{
		Length:   0,
		Capacity: capacity,
		Data:     make([]int, capacity),
	}
}

/**
顺序表在第index个节点插入数据
*/
func (s *SequentialList) Insert(index, value int) {
	//避免超过当前长度
	if index > s.Length {
		index = s.Length
	}
	//如果index为负数
	if index < 0 {
		index = 0
	}
	//如果长度超过了容积,则进行二倍扩容操作
	if s.Length+1 >= s.Capacity {
		capacity := 2 * s.Capacity
		data := make([]int, capacity)
		for i, v := range s.Data {
			data[i] = v
		}
		s.Data = data
		s.Capacity = capacity
	}
	for i := s.Length; i >= index; i-- {
		s.Data[i+1] = s.Data[i]
	}
	s.Data[index] = value
	s.Length++
}

/**
顺序表删除任意节点的数据
*/
func (s *SequentialList) Remove(index int) (int, error) {
	//避免超过当前长度
	if index < 0 || index > s.Length {
		return 0, errors.New("param index invalid")
	}
	val := s.Data[index]
	//执行删除操作
	for i := index + 1; i < s.Length; i++ {
		s.Data[i-1] = s.Data[i]
	}
	s.Length--
	return val, nil
}

/**
当字符串打印
*/
func (s *SequentialList) String() string {
	return fmt.Sprintf("%v", s.Data[:s.Length])
}

/**
打印所有值
*/
func (s *SequentialList) Print() {
	for _, v := range s.Data[:s.Length] {
		fmt.Println(v)
	}
}
