package bitmap

import (
	"sync"
)

type ConcurrentBitmap struct {
	Bitmap
	sync.RWMutex
}

func NewConcurrentBitmap(n int) ConcurrentBitmap {
	return ConcurrentBitmap{make(Bitmap, n), sync.RWMutex{}}
}

func (c ConcurrentBitmap) Len() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.Bitmap)
}

func (c ConcurrentBitmap) Contain(u uint) bool {
	c.RLock()
	defer c.RUnlock()
	return c.Bitmap.Contain(u)
}

func (c ConcurrentBitmap) Set(u uint) {
	c.Lock()
	defer c.Unlock()
	c.Bitmap.Set(u)
}

func (c ConcurrentBitmap) UnSet(u uint) {
	c.Lock()
	defer c.Unlock()
	c.Bitmap.UnSet(u)
}

func (c ConcurrentBitmap) FillOnes() {
	c.Lock()
	defer c.Unlock()
	c.Bitmap.FillOnes()
}

func (c ConcurrentBitmap) RunIterator(u uint, f func(uint)) {
	c.RLock()
	defer c.RUnlock()
	c.Bitmap.RunIterator(u, f)
}


var _ BitmapInterface = new(ConcurrentBitmap)