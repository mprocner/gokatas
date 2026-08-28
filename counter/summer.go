package counter

type Summer struct {
	total int
}

func (s *Summer) Add(n int) {
	s.total += n
}

func (s *Summer) Reset() {
	s.total = 0
}

func (s *Summer) Value() int {
	return s.total
}
