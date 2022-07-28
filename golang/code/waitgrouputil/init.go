package waitgrouputil

import (
	"sync"
)

var WaitGroup waitGroupWrapper

type waitGroupWrapper struct {
	waitGroup *sync.WaitGroup
}

func init() {
	WaitGroup.waitGroup = new(sync.WaitGroup)
}
