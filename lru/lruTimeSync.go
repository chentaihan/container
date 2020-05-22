package lru

import (
	"time"
)

/**
1：基于过期时间的LRU，将元素保存到集合中，expireTime秒后自动删除
2：采取的是定期删除机制，每100毫秒都会删除一定数量的过期的数据
3：所有接口都支持多协程访问
*/

type Action int

const (
	ACTION_ADD      Action = 1
	ACTION_GET      Action = 2
	ACTION_PEEK     Action = 3
	ACTION_LEN      Action = 4
	ACTION_CLEAR    Action = 5
	ACTION_GETARRAY Action = 6
)

type request struct {
	val        IObject
	expireTime int64
	hashCode   int
}

type response struct {
	result interface{}
	size   int
}

type command struct {
	action   Action
	request  *request
	response chan *response
}

type LruTimeSync struct {
	lruTime ILruTime
	cmdC    chan *command
	stopC   chan struct{}
	doneC   chan struct{}
}

func NewLruTimeSync(cap int) ILruTime {
	lru := &LruTimeSync{
		lruTime: NewLruTime(cap),
		cmdC:    make(chan *command, 6),
		stopC:   make(chan struct{}),
		doneC:   make(chan struct{}),
	}
	go lru.runLoop()
	return lru
}

func (mq *LruTimeSync) Stop() {
	close(mq.stopC)
	<-mq.doneC
}

func (mq *LruTimeSync) runLoop() {
	defer close(mq.doneC)
	for {
		select {
		case cmd := <-mq.cmdC:
			mq.handleCmd(cmd)
		case <-mq.stopC:
			return
		case <-time.After(time.Second / 10):  //定期删除过期数据
			mq.lruTime.RemoveOutOfTime()
		}
	}
}

func (mq *LruTimeSync) handleCmd(cmd *command) {
	switch cmd.action {
	case ACTION_ADD:
		mq.lruTime.Add(cmd.request.val, cmd.request.expireTime)
		cmd.response <- nil
	case ACTION_GET:
		result, exist := mq.lruTime.Get(cmd.request.hashCode)
		size := 0
		if exist {
			size = 1
		}
		cmd.response <- &response{
			result: result,
			size:   size,
		}
	case ACTION_LEN:
		size := mq.lruTime.Len()
		cmd.response <- &response{
			size: size,
		}
	case ACTION_PEEK:
		result, exist := mq.lruTime.Peek()
		size := 0
		if exist {
			size = 1
		}
		cmd.response <- &response{
			result: result,
			size:   size,
		}
	case ACTION_CLEAR:
		mq.lruTime.Clear()
		cmd.response <- nil
	case ACTION_GETARRAY:
		result := mq.lruTime.GetArray()
		cmd.response <- &response{
			result: result,
		}
	}
}

// 添加元素，如果已经存在的就更新过期时间
// expireTime秒后自动删除
func (mq *LruTimeSync) Add(val IObject, expireTime int64) {
	cmd := &command{
		action: ACTION_ADD,
		request: &request{
			val:        val,
			expireTime: expireTime,
			hashCode:   0,
		},
		response: make(chan *response),
	}
	mq.cmdC <- cmd
	<-cmd.response
}

func (mq *LruTimeSync) Get(hashCode int) (IObject, bool) {
	cmd := &command{
		action: ACTION_GET,
		request: &request{
			val:        nil,
			expireTime: 0,
			hashCode:   hashCode,
		},
		response: make(chan *response),
	}
	mq.cmdC <- cmd
	response := <-cmd.response
	var result IObject
	if response.result != nil {
		result = response.result.(IObject)
	}
	return result, response.size > 0
}

func (mq *LruTimeSync) RemoveOutOfTime() int {
	return 0
}

func (mq *LruTimeSync) Peek() (IObject, bool) {
	cmd := &command{
		action:   ACTION_PEEK,
		request:  nil,
		response: make(chan *response),
	}
	mq.cmdC <- cmd
	response := <-cmd.response
	var result IObject
	if response.result != nil {
		result = response.result.(IObject)
	}
	return result, response.size > 0
}

func (mq *LruTimeSync) Len() int {
	cmd := &command{
		action:   ACTION_LEN,
		request:  nil,
		response: make(chan *response),
	}
	mq.cmdC <- cmd
	response := <-cmd.response
	return response.size
}

func (mq *LruTimeSync) Clear() {
	cmd := &command{
		action:   ACTION_CLEAR,
		request:  nil,
		response: make(chan *response),
	}
	mq.cmdC <- cmd
	<-cmd.response
}

func (mq *LruTimeSync) GetArray() []IObject {
	cmd := &command{
		action:   ACTION_GETARRAY,
		request:  nil,
		response: make(chan *response),
	}
	mq.cmdC <- cmd
	response := <-cmd.response
	var result []IObject
	if response.result != nil {
		result = response.result.([]IObject)
	}
	return result
}
