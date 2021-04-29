package triples // Enkelvoud maken!!!

import "errors"

const (
	// NotFound int = -1
	Wildcard int = -1 // Kiezen we niet best overal voor int64? Dit loopt slecht af op een 32-bit systeem.
)

// Store is a data structure to compactly store triples. Store implements fast
// search.
type Store struct {
	ptrs0 []uint
	ptrs1 []uint
	vals1 []int
	vals2 []int
}

// Triple ...
type Triple struct {
	val0, val1, val2 int
}

// New ...
func New() Store {
	return Store{}
}

// From constructs a store from a list of triples.
func From(ts []Triple) (Store, error) {
	if len(ts) == 0 {
		return Store{}, errors.New("trie: number of triples is zero")
	}

	n := ts[len(ts)-1].val0 + 2
	ptrs0 := make([]uint, 0, n)
	ptrs1 := make([]uint, 0, len(ts)+1)
	vals1 := make([]int, 0, len(ts))
	vals2 := make([]int, 0, len(ts))

	var ptr0, ptr1 uint
	v0, v1 := -1, -1 // Zonder deze sentinels kunnen de values uint/uint64 zijn! Ik kan als sentinel ook math.MaxUint nemen!!!
	for _, t := range ts {
		if t.val0 != v0 {
			v0, v1 = t.val0, t.val1
			ptrs0 = append(ptrs0, ptr0)
			ptrs1 = append(ptrs1, ptr1)
			vals1 = append(vals1, v1)
			ptr0++
		} else if t.val1 != v1 {
			v1 = t.val1
			ptrs1 = append(ptrs1, ptr1)
			vals1 = append(vals1, v1)
			ptr0++
		}
		vals2 = append(vals2, t.val2)
		ptr1++
	}

	// sentinels
	ptrs0 = append(ptrs0, ptr0)
	ptrs1 = append(ptrs1, ptr1)

	return Store{ptrs0: ptrs0, ptrs1: ptrs1, vals1: vals1, vals2: vals2}, nil
}

// Triples ...
func (s *Store) Triples() []Triple {
	ts := make([]Triple, 0, len(s.vals2))

	var lo1, lo2 uint
	for i := 1; i < len(s.ptrs0); i++ {
		hi1 := s.ptrs0[i]
		for j := lo1; j < hi1 && j < uint(len(s.vals1)) && j+1 < uint(len(s.ptrs1)); j++ { // Zijn die extra tests sneller???
			v1 := s.vals1[j]
			hi2 := s.ptrs1[j+1]
			for k := lo2; k < hi2 && k < uint(len(s.vals2)); k++ {
				ts = append(ts, Triple{val0: i - 1, val1: v1, val2: s.vals2[k]})
			}
			lo2 = hi2
		}
		lo1 = hi1
	}

	return ts
}

// Build ...
func (s *Store) Build() {

}

// Select returns an iterator over a list of triples
// that fullfil the given triple selection pattern.
// func (s *Store) Select(tp Triple) Iter {
// 	i := tp.val0
// 	if i < 0 || len(s.ptrs0) <= i+1 {
// 		return Iter{invalid: true}
// 	}

// 	var ok bool
// 	j := s.ptrs0[i]
// 	if tp.val1 != Wildcard {
// 		vals1 := s.vals1[s.ptrs0[i]:s.ptrs0[i+1]]
// 		if j, ok = find(vals1, tp.val1); !ok {
// 			return Iter{invalid: true}
// 		}
// 	}
// 	first := Iter(s.ptrs0.iterAt(i), s.ptrs1.iterAt(j))

// 	if len(s.idxs) <= j {
// 		return Iter{invalid: true}
// 	}

// 	return Iter{i, first, second}
// }

// find ...
func find(arr []int, search int) (uint, bool) {
	var lo, hi uint = 0, uint(len(arr)) // Test doen om lineair te scannen indien lengte kleiner dan threshold!

	for lo < hi {
		m := lo + (hi-lo)>>1
		if arr[m] < search {
			lo = m + 1
		} else {
			hi = m
		}
	}

	return lo, true // We moeten nog NotFound toevoegen als bool!
}

// Iter ...
type Iter struct {
	store                      *Store
	p0, p1                     int // iteration state
	v0, p0Lo, p0Hi, p1Lo, p1Hi int // config iterator
}

// func (t *Triples) Iter() Iter {
// 	return Iter{
// 		t:  t,
// 		p0: p0l,
// 		p1: p1l,
// 		v0: v0,
// 		// p0Lo: p0l,
// 		p0Hi: p0h,
// 		// p1Lo: p1l,
// 		p1Hi: p1h,
// 	}
// }

// func (it *Iter) Triples() Triple {

// 	hi1 := t.ptrs[i+1]
// 	for j := range t.idxs[it.p0Lo:it.p0Hi] {
// 		v1 := t.idxs[lo1+j].val
// 		hi2 := t.idxs[lo1+j+1].ptr
// 		for _, v2 := range t.vals[lo2:hi2] {
// 			tps = append(tps, Triple{val0: i, val1: v1, val2: v2})
// 		}
// 		lo2 = hi2
// 	}

// 	return Triple{val0: it.v0, val1: v1, val2: v2}
// }

func (it *Iter) HasNext() bool {
	return true
}
