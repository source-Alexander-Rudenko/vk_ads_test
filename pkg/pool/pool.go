package pool

import (
	"context"
	"io"
	"sync"
)

type Pool struct {
	tasks   chan string
	control chan controlMsg
	workers map[int]*worker

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	mu     sync.Mutex
	nextID int
	writer io.Writer
}

// NewPool создаёт пул с указанным размером буфера и writer для вывода
func NewPool(buffer int, writer io.Writer) *Pool {
	ctx, cancel := context.WithCancel(context.Background())
	return &Pool{
		tasks:   make(chan string, buffer),
		control: make(chan controlMsg),
		workers: make(map[int]*worker),
		ctx:     ctx,
		cancel:  cancel,
		writer:  writer,
	}
}

// Run запускает менеджер пула в своей горутине
func (p *Pool) Run() {
	go p.controlLoop()
}

// controlLoop слушает команды управления и сигналы завершения
func (p *Pool) controlLoop() {
	for {
		select {
		case <-p.ctx.Done():
			// graceful shutdown всех воркеров
			p.mu.Lock()
			for _, w := range p.workers {
				w.cancel()
			}
			p.mu.Unlock()
			p.wg.Wait()
			return

		case msg := <-p.control:
			switch msg.action {
			case addWorker:
				p.mu.Lock()
				id := p.nextID
				p.nextID++
				w := newWorker(id, p.tasks, &p.wg, p.ctx, p.writer)
				p.workers[id] = w
				p.mu.Unlock()
				msg.resp <- nil

			case removeWorker:
				p.mu.Lock()
				w, ok := p.workers[msg.workerID]
				if !ok {
					p.mu.Unlock()
					msg.resp <- ErrWorkerNotFound
				} else {
					w.cancel()
					delete(p.workers, msg.workerID)
					p.mu.Unlock()
					msg.resp <- nil
				}
			}
		}
	}
}

// AddWorker добавляет нового воркера в пул
func (p *Pool) AddWorker() error {
	resp := make(chan error)
	p.control <- controlMsg{action: addWorker, resp: resp}
	return <-resp
}

// RemoveWorker удаляет воркера по id
func (p *Pool) RemoveWorker(id int) error {
	resp := make(chan error)
	p.control <- controlMsg{action: removeWorker, workerID: id, resp: resp}
	return <-resp
}

// Submit ставит задачу в очередь или возвращает ощтбку
func (p *Pool) Submit(task string) error {
	select {
	case p.tasks <- task:
		return nil
	case <-p.ctx.Done():
		return ErrPoolStopped
	}
}

// Shutdown останавливает и пул и всех воркеров
func (p *Pool) Shutdown() {
	p.cancel()
}
