package main

// Info ...
type Info struct {
	ID     int    `json:"myID"`
	Name   string `json:"myName"`
	Time   string `json:"myTime"`
	Enable bool   `json:"myEnable"`
}

// CacheData ...
var CacheChannel chan Info

func main() {

	CacheChannel = make(chan Info, 10000)

	MainChannel = make(chan Info, 1000)
	go GoCacheData(MainChannel, 5) // 原材料的中间准备

	go goAnalysis()

	consume()
}
