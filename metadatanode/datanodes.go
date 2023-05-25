package metadatanode

import "sync"

type DatanodeItem struct {
	Status         []int //We store last 3 statuses
	DataNodeNumber int32
	IpAddr         string
	LastContact    int64
}

// In memory database for storing datanodes connected to metadatanode. Methods are
// safe to call concurrently.
type datanodeStore struct {
	sync.Mutex

	results map[int]DatanodeItem
	nextId  int
}

func NewDatanodeDatabase() *datanodeStore {
	ts := &datanodeStore{}
	ts.results = make(map[int]DatanodeItem)
	ts.nextId = 0
	return ts
}

func (ts *datanodeStore) Add(result DatanodeItem) {
	ts.Lock()
	defer ts.Unlock()
	ts.results[ts.nextId] = result
	ts.nextId++
}

func (ts *datanodeStore) Delete(id int) {
	ts.Lock()
	defer ts.Unlock()
	delete(ts.results, id)
}

// TODO imho it should return error not only bool
func (ts *datanodeStore) Get(id int) (DatanodeItem, bool) {
	ts.Lock()
	defer ts.Unlock()
	result, ok := ts.results[id]
	return result, ok
}

func (ts *datanodeStore) Update(id int, result DatanodeItem) {
	ts.Lock()
	defer ts.Unlock()
	ts.results[id] = result
}
