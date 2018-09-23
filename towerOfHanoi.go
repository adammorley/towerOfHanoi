// solve the tower of hanoi problem iteratively
package towerOfHanoi

import "log"

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
		log.Fatal("cannot push onto smaller value")
	}
	d.next = s.top
	s.top = d
	s.size++
}

/*
   the iterative solution operates by noticing a pattern (you have to draw
   out n==1, n==2, n==3, and n==4 (and maybe n==5 unless you're adventurous)
   to see the pattern

   note that there is a recursive element to this, in that if n > 1, the
   function will call itself with the same definition for the three posts.

   one interesting aspect of this is that origin, buffer and destination
   can "flip" if the values need to go the other way (see compareAndMove)
*/
func moveDiscs(o, b, d *stack) {
	if o.size >= 64 {
		log.Fatal("unsupported disc count")
	} else if o.size == 1 && b.size == 0 && d.size == 0 {
		/*
		   if there is only one disk on the origin, simply
		   move it to the destination
		*/
		d.Push(o.Pop())
		return
	} else if (o.size+b.size+d.size)%2 == 0 { // even
		/*
		   if the number of disks is even, move from:
		       origin -> buffer
		       origin -> destination
		       buffer -> destination
		   if no more disks on origin and buffer, complete
		*/
		compareAndMove(o, b)
		compareAndMove(o, d)
		compareAndMove(b, d)
		if done(o, b) {
			return
		}
		moveDiscs(o, b, d)
	} else { // odd
		/*
		   if odd, move:
		       origin -> destination
		           check if done (the last disk moves directly as
		           opposed to being buffered)
		       origin -> buffer
		       destination -> buffer
		*/
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
