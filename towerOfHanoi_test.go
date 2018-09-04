package towerOfHanoi

import "testing"

func Test(t *testing.T) {
	var cases []uint = []uint{1, 2, 3, 4, 5, 6, 8, 15, 16, 20}
	var n uint
	for i := range cases {
		n = cases[i]
		o, b, d := Disc(n)
		moveDiscs(o, b, d)
		validate(t, n, d)
	}
}

func validate(t *testing.T, n uint, s *stack) {
	d := s.top
	var i uint
	for i = 1; i <= n; i++ {
		if d.value != i {
			t.Error("wrong value", i, d.value)
		}
		d = d.next
	}
}

func Disc(n uint) (s0, s1, s2 *stack) {
	s0, s1, s2 = posts()
	d := &disc{1, nil}
	s0.top = d
	var i uint
	for i = 2; i <= n; i++ {
		d.next = &disc{i, nil}
		d = d.next
	}
	s0.size = n
	return s0, s1, s2
}

func oneDisc() (s0, s1, s2 *stack) {
	s0, s1, s2 = posts()
	s0.top = &disc{1, nil}
	s0.size = 1
	return
}
func twoDisc() (s0, s1, s2 *stack) {
	s0, s1, s2 = posts()
	s0.top = &disc{1, &disc{2, nil}}
	s0.size = 2
	return
}
func threeDisc() (s0, s1, s2 *stack) {
	s0, s1, s2 = posts()
	s0.top = &disc{1, &disc{2, &disc{3, nil}}}
	s0.size = 3
	return
}
func posts() (s0, s1, s2 *stack) {
	s0, s1, s2 = new(stack), new(stack), new(stack)
	return
}
