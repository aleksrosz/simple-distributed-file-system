package metadatanode

import "sync"

type blockReport struct {
	fileName       string
	blockID        int
	dataNodeNumber int
}

// In memory database for storing blockReport results. blockReport methods are
// safe to call concurrently.
type blockReportStore struct {
	sync.Mutex

	results map[int]blockReport
	nextId  int
}

func New() *blockReportStore {
	ts := &blockReportStore{}
	ts.results = make(map[int]blockReport)
	ts.nextId = 0
	return ts
}

func (ts *blockReportStore) Add(result blockReport) {
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

func (ts *blockReportStore) Get(id int) (blockReport, bool) {
	ts.Lock()
	defer ts.Unlock()
	result, ok := ts.results[id]
	return result, ok
}

func (ts *blockReportStore) Update(id int, result blockReport) {
	ts.Lock()
	defer ts.Unlock()
	ts.results[id] = result
}
