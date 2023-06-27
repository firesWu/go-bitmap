package bitmap

import "math"

type Bitmap []uint32

var bitLen = uint(32)
var maxUint32 = uint32(math.MaxUint32)

const intSize = 32 << (^uint(0) >> 63)

func New(n int) Bitmap {
	return make(Bitmap, n)
}

func (b Bitmap) Len() int {
	return len(b)
}

func (b Bitmap) Contain(i uint) bool {
	blkAt := int(i / bitLen)
	if blkAt >= len(b) {
		return false
	}
	return b[blkAt]>>(i%bitLen)&1 == 1
}

func (b Bitmap) Set(i uint) {
	blkAt := int(i / bitLen)
	if blkAt >= len(b) {
		return
	}
	b[blkAt] |= 1 << (i % bitLen)
}

func (b Bitmap) UnSet(i uint) {
	blkAt := int(i / bitLen)
	if blkAt >= len(b) {
		return
	}
	b[blkAt] &= maxUint32 ^ (1 << (i % bitLen))
}

func (b Bitmap) FillOnes() {
	for i := 0; i < len(b); i++ {
		b[i] = maxUint32
	}
}

func (b Bitmap) RunIterator(start uint, f func(uint)) {
	mx := uint(len(b)) * bitLen
	for i := start; i <= mx; i++ {
		if b.Contain(i) {
			f(i)
		}
	}
}
