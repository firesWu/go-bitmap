package bitmap

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func BenchmarkBitmap(b *testing.B) {
	b.Run("Len", func(b *testing.B) {
		bm := New(128/32)
		assert.Equal(b, bm.Len(), 128/32)
	})

	b.Run("Set", func(b *testing.B) {
		bm := New(128/32)
		bm.Set(10000)
		assert.Equal(b, bm.Contain(10000), false)
	})

	b.Run("Set", func(b *testing.B) {
		bm := New(128/32)
		bm.Set(127)
		assert.Equal(b, bm.Contain(127), true)
	})

	b.Run("Unset", func(b *testing.B) {
		bm := New(128/32)
		bm.Set(12)
		assert.Equal(b, bm.Contain(12), true)
		bm.UnSet(12)
		assert.Equal(b, bm.Contain(12), false)
	})

	b.Run("FillOnes", func(b *testing.B) {
		bm := New(128/32)
		bm.FillOnes()
		for i := 0; i < len(bm); i++ {
			assert.Equal(b, bm[i], maxUint32)
		}
	})

	b.Run("RunIterator", func(b *testing.B) {
		tc := []uint{2, 9, 1, 30, 10, 20}
		bm := New(32/32)
		for _, v := range tc {
			bm.Set(v)
		}

		sort.Slice(tc, func(i, j int) bool {
			return tc[i] < tc[j]
		})

		for i := range tc {
			tc2 := tc[i:]
			var res []uint
			bm.RunIterator(tc[i], func(i uint) {
				res = append(res, i)
			})

			assert.Equal(b, len(tc2), len(res))
			for i, v := range res {
				assert.Equal(b, tc2[i], v)
			}
		}
	})
}

func BenchmarkConcurrentMap(b *testing.B) {
	b.Run("Len", func(b *testing.B) {
		bm := NewConcurrentBitmap(128/32)
		assert.Equal(b, bm.Len(), 128/32)
	})

	b.Run("Set", func(b *testing.B) {
		bm := NewConcurrentBitmap(128/32)
		bm.Set(10000)
		assert.Equal(b, bm.Contain(10000), false)
	})

	b.Run("Set", func(b *testing.B) {
		bm := NewConcurrentBitmap(128/32)
		bm.Set(127)
		assert.Equal(b, bm.Contain(127), true)
	})

	b.Run("Unset", func(b *testing.B) {
		bm := NewConcurrentBitmap(128/32)
		bm.Set(12)
		assert.Equal(b, bm.Contain(12), true)
		bm.UnSet(12)
		assert.Equal(b, bm.Contain(12), false)
	})

	b.Run("FillOnes", func(b *testing.B) {
		bm := NewConcurrentBitmap(128/32)
		bm.FillOnes()
		for i := 0; i < bm.Len(); i++ {
			assert.Equal(b, bm.Bitmap[i], maxUint32)
		}
	})

	b.Run("RunIterator", func(b *testing.B) {
		tc := []uint{2, 9, 1, 30, 10, 20}
		bm := NewConcurrentBitmap(32/32)
		for _, v := range tc {
			bm.Set(v)
		}

		sort.Slice(tc, func(i, j int) bool {
			return tc[i] < tc[j]
		})

		for i := range tc {
			tc2 := tc[i:]
			var res []uint
			bm.RunIterator(tc[i], func(i uint) {
				res = append(res, i)
			})

			assert.Equal(b, len(tc2), len(res))
			for i, v := range res {
				assert.Equal(b, tc2[i], v)
			}
		}
	})
}