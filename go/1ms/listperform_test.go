package list_test

import (
	"container/list"
	"testing"
	"time"

	"./init.go"
)

var basicList = init.InitBasic()

func BenchmarkList2(b *testing.B) {
	nums := list.New()
	c := time.Tick(1 * time.Millisecond)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = <-c
		nums.PushBack(basicList.Front())
	}
}
