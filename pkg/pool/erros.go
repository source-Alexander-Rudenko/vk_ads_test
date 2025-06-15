package pool

import "errors"

var (
	// ErrWorkerNotFound возвращается, если воркера с указанным id нет
	ErrWorkerNotFound = errors.New("worker not found")

	// ErrPoolStopped возвращается если писать в уже закрытый пул
	ErrPoolStopped = errors.New("pool stopped")
)
