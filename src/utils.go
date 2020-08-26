package fsm

import (
	"reflect"
	"runtime"
	"strings"
)

func getFunctionName(back CCallBack) string {
	name := runtime.FuncForPC(reflect.ValueOf(back).Pointer()).Name()
	NamesOne := strings.Split(name, "/")
	NamesTwo := strings.Split(NamesOne[len(NamesOne)-1], ".")
	finalName := NamesTwo[len(NamesTwo)-1]
	return finalName
}
