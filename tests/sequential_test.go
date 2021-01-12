package tests

import (
	"dataStruct/src/list"
	"testing"
)

//初始化顺序表
var seq *list.SequentialList

func init() {
	seq = list.NewSequentialList(1)
}

//测试顺序表新增数据
func TestInsert(t *testing.T) {
	seq.Insert(0, 1)
	seq.Insert(1, 2)
	seq.Insert(2, 3)
	seq.Insert(3, 4)
	t.Logf("%v", seq)
}

//测试顺序表删除数据
func TestRemove(t *testing.T) {
	seq.Insert(0, 1)
	seq.Insert(1, 2)
	seq.Insert(2, 3)
	seq.Insert(1, 4)
	seq.Insert(1, 5)
	res, err := seq.Remove(0)
	if err == nil {
		t.Logf("删除了%d\n", res)
	}
	t.Logf("%v", seq)
}
