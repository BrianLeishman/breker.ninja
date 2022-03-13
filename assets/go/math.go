package breker

import (
	"sync"

	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Integer | constraints.Float
}

type MovingAverage[T Numeric] struct {
	values []T
	sum    T
	size   int
	curr   int
	mx     sync.RWMutex
}

func NewMovingAverage[T Numeric](size int) *MovingAverage[T] {
	avg := &MovingAverage[T]{
		size:   size,
		values: make([]T, 0, size),
	}

	return avg
}

func (a *MovingAverage[T]) Record(n T) {
	a.mx.Lock()
	defer a.mx.Unlock()

	a.curr++
	if a.curr >= a.size {
		a.curr = 0
	}

	if len(a.values) < a.size {
		a.values = append(a.values, n)
	} else {
		a.sum -= a.values[a.curr]
		a.values[a.curr] = n
	}

	a.sum += n
}

func (a *MovingAverage[T]) Average() float64 {
	a.mx.RLock()
	defer a.mx.RUnlock()

	return float64(a.sum) / float64(T(a.size))
}

func ConvertRange[T Numeric](oldValue, oldMin, oldMax, newMin, newMax T) T {
	oldRange := oldMax - oldMin
	if oldRange == 0 {
		return newMin
	}

	newRange := newMax - newMin
	return (((oldValue - oldMin) * newRange) / oldRange) + newMin
}
