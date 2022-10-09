package common

const Endpoint = "http://127.0.0.1:9200"
const JobRetIndex = "job_ret"

type JobRet struct {
	Id     string `json:"id"`
	HostId string `json:"hostId"`
	CronId string `json:"cronId"`
	Ctime  int64  `json:"ctime"`
}

type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}
type Total struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}
type Hits struct {
	Total    Total          `json:"total"`
	MaxScore interface{}    `json:"max_score"`
	Hits     []CronDocsItem `json:"hits"`
}
type CronDocs struct {
	Hits Hits `json:"hits"`
}
type Buckets struct {
	Key      string   `json:"key"`
	DocCount int      `json:"doc_count"`
	CronDocs CronDocs `json:"cron_docs"`
}
type Cron struct {
	DocCountErrorUpperBound int       `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int       `json:"sum_other_doc_count"`
	Buckets                 []Buckets `json:"buckets"`
}

type Aggregations struct {
	Cron Cron `json:"cron"`
}

type CronDocsItem struct {
	Index  string      `json:"_index"`
	Type   string      `json:"_type"`
	ID     string      `json:"_id"`
	Score  interface{} `json:"_score"`
	Source Source      `json:"_source"`
	Sort   []int64     `json:"sort"`
}
type Source struct {
	HostID string `json:"hostId"`
	CronID string `json:"cronId"`
	Ctime  int64  `json:"ctime"`
}
