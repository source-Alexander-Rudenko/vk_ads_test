package pool

type actionType int

const (
	addWorker actionType = iota
	removeWorker
)

type controlMsg struct {
	action   actionType
	workerID int
	resp     chan error
}
