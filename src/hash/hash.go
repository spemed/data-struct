package hash

//定义解决冲突策略不同的哈希表
type Hash interface {
	Insert()              //插入
	Find()                //寻找元素
	Delete()              //删除元素
	Reload()              //重载
	hashKey(string) int64 //hash算法
}
