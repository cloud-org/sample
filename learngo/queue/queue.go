package queue

type Quque []interface{} // interface{} 表示各种类型

func (q *Quque) Push(v interface{}) {
	*q = append(*q, v.(int))
}

func (q *Quque) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

func (q *Quque) IsEmpty() bool {
	return len(*q) == 0
}
