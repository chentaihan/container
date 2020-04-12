package queue

type Queue struct {
	buf  []interface{}
	head int
	tail int
	cap  int
	l    int
}

func NewQueue(cap int) IQueue {
	if cap < 4 {
		cap = 4
	}
	return &Queue{
		buf:  make([]interface{}, cap),
		head: 0,
		tail: -1,
		cap:  cap,
		l:    0,
	}
}

//扩容，小于1M时2倍增加，否则加1M
func (s *Queue) checkCap() {
	if s.l < s.cap {
		return
	}
	newCap := s.cap << 1
	if s.cap > 1024*1024 {
		newCap = s.cap + 1024*1024
	}
	buf := make([]interface{}, newCap)
	start := s.head
	for index := 0; start <= s.tail; index++ {
		buf[index] = s.buf[(start+index)%s.cap]
		start++
	}
	s.buf = buf
	s.cap = newCap
	s.head = 0
	s.tail = s.l - 1
}

//进队列，尾部添加元素
func (s *Queue) Enqueue(val interface{}) {
	s.checkCap()
	s.tail++
	s.buf[s.tail%s.cap] = val
	s.l++
}

//出队列，头往尾走
func (s *Queue) Dequeue() (interface{}, bool) {
	if s.l == 0 {
		return nil, false
	}
	val := s.buf[s.head%s.cap]
	s.buf[s.head%s.cap] = nil //释放对变量的引用
	s.head++
	s.l--
	return val, true
}

func (s *Queue) Empty() bool {
	return s.l == 0
}

func (s *Queue) Len() int {
	return s.l
}

func (s *Queue) Cap() int {
	return s.cap
}

func (s *Queue) Peek() (interface{}, bool) {
	if s.l == 0 {
		return nil, false
	}
	return s.buf[s.head%s.cap], true
}

func (s *Queue) ToList() []interface{} {
	buf := make([]interface{}, s.Len())
	index := 0
	for i := s.head; i <= s.tail; i++ {
		buf[index] = s.buf[i%s.cap]
		index++
	}
	return buf
}
