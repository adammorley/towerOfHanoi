// solve the tower of hanoi problem iteratively
package towerOfHanoi

type stack struct {
	top  *disc
	size uint
}
type disc struct {
	value uint
	next  *disc
}

const noDisc = uint(64) // because stack overflow at <32

func (s *stack) Peek() uint {
	if s.top == nil {
		return noDisc
	}
	return s.top.value
}
func (s *stack) Pop() *disc {
	d := s.top
	s.top = d.next
	s.size--
	d.next = nil
	return d
}
func (s *stack) Push(d *disc) {
	if s.top != nil && d.value > s.Peek() {
		panic("nope")
	}
	d.next = s.top
	s.top = d
	s.size++
}
func moveDiscs(o, b, d *stack) {
	if o.size >= 64 {
		panic("unsupported disc count")
	} else if o.size == 1 && b.size == 0 && d.size == 0 {
		d.Push(o.Pop())
		return
	} else if (o.size+b.size+d.size)%2 == 0 { // even
		compareAndMove(o, b)
		compareAndMove(o, d)
		compareAndMove(b, d)
		if done(o, b) {
			return
		}
		moveDiscs(o, b, d)
	} else { // odd
		compareAndMove(o, d)
		if done(o, b) {
			return
		}
		compareAndMove(o, b)
		compareAndMove(d, b)
		moveDiscs(o, b, d)
	}
}

func compareAndMove(a, b *stack) {
	if a.Peek() < b.Peek() {
		b.Push(a.Pop())
	} else {
		a.Push(b.Pop())
	}
}

func done(a, b *stack) bool {
	if a.size == 0 && b.size == 0 {
		return true
	}
	return false
}
