package hash

const (
	empty   = 1 //数据位空置
	busy    = 2 //数据位使用中
	deleted = 3 //数据标识删除
)

/**
使用线性探查解决hash冲突
	优点:实现简单
	缺点:
*/
type LineHash struct {
	Data    []int64 //底层数据存储
	Length  int     //长度
	Divisor int     //重载因子,用于计算哈希桶的利用率
	Status  []int64 //用于标识数据位状态
}
