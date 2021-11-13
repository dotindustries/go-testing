package testing

import (
	"log"
	"path"
	"runtime"
	"strconv"
)

func Trace(skip int) {
	callerPtr, file, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	caller := runtime.FuncForPC(callerPtr)
	log.Println("Trace", path.Base(caller.Name()), path.Base(file)+":"+strconv.Itoa(line))
}
