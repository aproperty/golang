package main

import (
	"container/list"
	"time"
)

// CacheDataStruct ...
type CacheDataStruct struct {
	List    list.List
	MinTime string
	SendLen int
}

// GCacheData ...
var GCacheData *CacheDataStruct

func init() {
	node := new(CacheDataStruct)
	node.List = *list.New()
	node.List.Init()
	node.SendLen = 0
	GCacheData = node
}

// SendCacheData ...
func SendCacheData(to chan Info) {
	GCacheData.SendLen = GCacheData.List.Len()
	if GCacheData.SendLen == 0 {
		return
	}
	for i := 0; i < GCacheData.SendLen; i++ {
		iter := GCacheData.List.Front()
		totalmsg := iter.Value.(*Info)
		to <- *totalmsg
		GCacheData.List.Remove(iter)
	}
	GCacheData.SendLen = GCacheData.List.Len()
}

// GoCacheData ...
func GoCacheData(mainChannel chan Info, sec int64) {

	t := time.NewTimer(time.Duration(sec) * time.Second)

	for {

		select {
		case totalmsg := <-CacheChannel:
			GCacheData.List.PushBack(&totalmsg)

		case <-t.C:
			SendCacheData(mainChannel)
			t.Reset(time.Duration(sec) * time.Second)
		}

	}

}
