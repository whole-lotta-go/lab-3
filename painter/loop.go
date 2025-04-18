package painter

import (
	"image"
	"sync"

	"github.com/whole-lotta-go/lab-3/ui"
	"golang.org/x/exp/shiny/screen"
)

var size = image.Pt(ui.WindowSide, ui.WindowSide)

// Receiver отримує текстуру, яка була підготовлена в результаті виконання команд у циклі подій.
type Receiver interface {
	Update(t screen.Texture)
}

// Loop реалізує цикл подій для формування текстури отриманої через виконання операцій отриманих з внутрішньої черги.
type Loop struct {
	Receiver Receiver

	next screen.Texture // текстура, яка зараз формується
	prev screen.Texture // текстура, яка була відправлення останнього разу у Receiver

	oq opQueue

	stop    chan struct{}
	stopReq bool
}

// Start запускає цикл подій. Цей метод потрібно запустити до того, як викликати на ньому будь-які інші методи.
func (l *Loop) Start(s screen.Screen) {
	l.next, _ = s.NewTexture(size)
	l.prev, _ = s.NewTexture(size)

	l.stop = make(chan struct{})
	l.oq.ne = sync.NewCond(&l.oq.mu)
	go func() {
		defer close(l.stop)

		for !(l.stopReq && l.oq.empty()) {
			op := l.oq.pull()

			if update := op.Do(l.next); update {
				l.Receiver.Update(l.next)
				l.next, l.prev = l.prev, l.next
			}
		}
	}()
}

// Post додає нову операцію у внутрішню чергу.
func (l *Loop) Post(op Operation) {
	l.oq.push(op)
}

// StopAndWait сигналізує про необхідність завершити цикл та блокується до моменту його повної зупинки.
func (l *Loop) StopAndWait() {
	l.Post(OperationFunc(func(screen.Texture) {
		l.stopReq = true
	}))
	<-l.stop
}

type opQueue struct {
	ops []Operation

	mu sync.Mutex
	ne *sync.Cond
}

func (oq *opQueue) push(op Operation) {
	oq.mu.Lock()
	defer oq.mu.Unlock()

	isEmpty := oq.lockedEmpty()
	oq.ops = append(oq.ops, op)

	if isEmpty {
		oq.ne.Signal()
	}
}

func (oq *opQueue) pull() Operation {
	oq.mu.Lock()
	defer oq.mu.Unlock()

	for oq.lockedEmpty() {
		oq.ne.Wait()
	}
	op := oq.ops[0]
	oq.ops[0] = nil
	oq.ops = oq.ops[1:]
	return op
}

func (oq *opQueue) lockedEmpty() bool {
	return len(oq.ops) == 0
}

func (oq *opQueue) empty() bool {
	oq.mu.Lock()
	defer oq.mu.Unlock()

	return oq.lockedEmpty()
}
