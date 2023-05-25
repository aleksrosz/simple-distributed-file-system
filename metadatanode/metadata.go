package metadatanode

import "sync"

type BlockReportItem struct {
	FileName       string
	BlockID        int32
	DataNodeNumber int32
	//IpAddr         string //TODO
}

// In memory database for storing blockReport results. blockReport methods are
// safe to call concurrently.
type blockReportStore struct {
	sync.Mutex

	results map[int]BlockReportItem
	nextId  int
}

func NewDatabase() *blockReportStore {
	ts := &blockReportStore{}
	ts.results = make(map[int]BlockReportItem)
	ts.nextId = 0
	return ts
}

func (ts *blockReportStore) Add(result BlockReportItem) {
	ts.Lock()
	defer ts.Unlock()
	ts.results[ts.nextId] = result
	ts.nextId++
}

func (ts *blockReportStore) Delete(id int) {
	ts.Lock()
	defer ts.Unlock()
	delete(ts.results, id)
}

func (ts *blockReportStore) Get(id int) (BlockReportItem, bool) {
	ts.Lock()
	defer ts.Unlock()
	result, ok := ts.results[id]
	return result, ok
}

func (ts *blockReportStore) Update(id int, result BlockReportItem) {
	ts.Lock()
	defer ts.Unlock()
	ts.results[id] = result
}
