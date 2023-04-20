package datanode

// A blockreport contains a list of all blocks on a DataNode.
type blockReport struct {
	fileName       string
	blockID        int
	dataNodeNumber int
}
