package pool

import (
	"context"
	"fmt"
	"io"
	"sync"
)

// Worker обрабатывает задачи из канала, пока контекст не отменили
type worker struct {
	id     int
	tasks  chan string
	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup
	writer io.Writer
}

// newWorker создаёт и запускает  воркера
func newWorker(
	id int,
	tasks chan string,
	wg *sync.WaitGroup,
	parentCtx context.Context,
	writer io.Writer,
) *worker {
	ctx, cancel := context.WithCancel(parentCtx)
	w := &worker{id: id, tasks: tasks, ctx: ctx, cancel: cancel, wg: wg, writer: writer}
	wg.Add(1)
	go w.run()
	return w
}

// Run слушает задачи и печатает их, пока не зщакрыт контекст
func (w *worker) run() {
	defer w.wg.Done()
	for {
		select {
		case <-w.ctx.Done():
			return
		case task := <-w.tasks:
			fmt.Fprintf(w.writer, "Worker %d: %s\n", w.id, task)
		}
	}
}
