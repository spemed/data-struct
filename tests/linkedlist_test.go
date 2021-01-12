package tests

import (
	"dataStruct/src/list"
	"testing"
)

var linkedList *list.LinkedList

func init() {
	linkedList = list.NewLinkedList()
}

func TestInsertLinkedList(t *testing.T) {
	linkedList.Insert(0, 1)
	linkedList.Insert(1, 2)
	linkedList.Insert(2, 2)
	linkedList.Insert(5, 3)
	linkedList.Insert(22, -100)
	t.Logf("%s", linkedList)
}

func TestRemoveLinkedList(t *testing.T) {
	linkedList.Insert(0, 1)
	linkedList.Insert(1, 2)
	linkedList.Insert(2, 22)
	linkedList.Insert(4, 123)
	val, err := linkedList.Remove(-1)
	if err != nil {
		t.Fatalf("fail to remove linked list node,err is %v", err)
	}
	t.Logf("delete val is %d", val)
	t.Logf("%s", linkedList)
}

func TestRemoveValLinkedList(t *testing.T) {
	linkedList.Insert(0, 1)
	linkedList.Insert(1, 2)
	linkedList.Insert(2, 22)
	linkedList.Insert(4, 123)
	err := linkedList.RemoveVal(123)
	if err != nil {
		t.Fatalf("fail to remove linked list node,err is %v", err)
	}
	t.Logf("%s", linkedList)
}

func TestAppendLinkedList(t *testing.T) {
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(22)
	linkedList.Append(123)
	t.Logf("%s", linkedList)
}

func TestDeleteLinkedList(t *testing.T) {
	linkedList.Insert(0, 1)
	linkedList.Insert(1, 2)
	linkedList.Insert(2, 22)
	linkedList.Insert(4, 123)
	linkedList.Insert(4, 123)
	data, _ := linkedList.Delete()
	_, _ = linkedList.Delete()
	_, _ = linkedList.Delete()
	_, _ = linkedList.Delete()
	_, err := linkedList.Delete()
	if err != nil {
		t.Error(err)
	}
	t.Logf("remove data:%d\n", data)
	t.Logf("%s", linkedList)
}
