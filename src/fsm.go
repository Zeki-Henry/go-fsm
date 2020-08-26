package fsm

import (
	"sync"
)

type CCallBack func(opt ...string)

type Transition struct {
	Action string
	From   []string
	To     string
}

type IFSM interface {
	GetState() string
	Trans(name string)
	IsValid(name string) bool
}

type CFSM struct {
	prevState string
	curState  string

	validTables map[string]map[string]string
	callBacks   map[string]CCallBack
	sync.Mutex
}

func (C *CFSM) GetState() string {
	return C.curState
}

func (C *CFSM) IsValid(name string) bool {
	if _, ok := C.validTables[name]; !ok {
		return false
	}
	if _, ok := C.validTables[name][C.curState]; !ok {
		return false
	}

	return true
}

func (C *CFSM) Trans(name string) {
	C.Lock()
	defer C.Unlock()

	if C.IsValid(name) {
		if _, ok := C.callBacks["onLeave" + C.curState]; ok {
			C.callBacks["onLeave" + C.curState](C.curState, C.prevState, name)
		}
		C.prevState = C.curState
		C.curState = C.validTables[name][C.curState]
		if _, ok := C.callBacks["onTransition"]; ok {
			C.callBacks["onTransition"](C.curState, C.prevState, name)
		}
		if _, ok := C.callBacks["onEnter" + C.curState]; ok {
			C.callBacks["onEnter" + C.curState](C.curState, C.prevState, name)
		}
		return
	}

	if _, ok := C.callBacks["onFailed"]; ok {
		C.callBacks["onFailed"](C.curState, C.prevState, name)
	}
	return
}

func NewFSM(trans []Transition, callBacks []CCallBack, defaultState string) IFSM {
	res := &CFSM{}
	res.callBacks = map[string]CCallBack{}
	res.validTables = map[string]map[string]string{}
	res.curState = defaultState

	for _, tran := range trans {
		if _, ok := res.validTables[tran.Action]; !ok {
			res.validTables[tran.Action] = map[string]string{}
		}
		for _, fromName := range tran.From {
			res.validTables[tran.Action][fromName] = tran.To
		}
	}

	for _, v := range callBacks {
		func_name := getFunctionName(v)
		res.callBacks[func_name] = v
	}
	return res
}
