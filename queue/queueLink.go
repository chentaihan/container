package queue

type queueItem struct {
	next *queueItem
	val  interface{}
}

type QueueLink struct {
	head *queueItem
	tail *queueItem
	l    int
}

func NewQueueLink() IQueue {
	return &QueueLink{}
}

func (s *QueueLink) Enqueue(val interface{}) {
	item := &queueItem{
		val: val,
	}
	if s.tail == nil {
		s.head = item
	} else {
		s.tail.next = item
	}
	s.l++
	s.tail = item
}

func (s *QueueLink) Dequeue() (interface{}, bool) {
	if s.head != nil {
		val := s.head.val
		s.head = s.head.next
		if s.head == nil {
			s.tail = nil
		}
		s.l--
		return val, true
	}
	return nil, false
}

func (s *QueueLink) Empty() bool {
	return s.head == nil
}

func (s *QueueLink) Peek() (interface{}, bool) {
	if s.head != nil {
		return s.head.val, true
	}
	return -1, false
}

func (s *QueueLink) Len() int {
	return s.l
}

func (s *QueueLink) Cap() int {
	return s.l
}

func (s *QueueLink) ToList() []interface{} {
	buf := make([]interface{}, s.Len())
	index := 0
	for i := s.head; i != nil; i = i.next {
		buf[index] = s.head.val
		index++
	}
	return buf
}
