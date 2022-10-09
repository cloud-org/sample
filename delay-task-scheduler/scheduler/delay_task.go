package scheduler

// 延时任务

type State int64

const (
	Pending State = iota
	Success
	Failed
)

func (s State) String() string {
	return [...]string{"pending", "success", "failed"}[s]
}

//DelayTask 延时任务
type DelayTask struct {
	Tid      string `json:"tid" bson:"tid"`             // task id -> Id.Hex()
	Name     string `json:"name" bson:"name"`           // task 名字
	Disable  bool   `json:"disable" bson:"disable"`     // 是否禁用当前任务
	State    State  `json:"state" bson:"state"`         // 是否已经完成
	CreateAt int64  `json:"create_at" bson:"create_at"` // 创建时间
	UpdateAt int64  `json:"update_at" bson:"update_at"` // 修改时间
	Delay    int64  `json:"delay" bson:"delay"`         // 延时几秒
}
