
```go
type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32
	elemtype *_type // element type
	sendx    uint   // send index
	recvx    uint   // receive index
	recvq    waitq  // list of recv waiters
	sendq    waitq  // list of send waiters

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex
}

type waitq struct {
	first *sudog
	last  *sudog
}
```




channel 变量是一个存储在函数栈帧上的指针，占用 8 个字节，指向堆上的 hchan 结构体  

## channel 为什么是线程安全的
不同协程通过 channel 进行通信，本身的使用场景就是多线程，为了保证数据的一致性，必须实现线程安全。  
因此 channel 的底层实现中，hchan 结构体中采用 mutex 锁来保证数据读写安全。  
在对循环数组 buf 中的数据进行入队和出队操作时，必须先获取互斥锁，才能操作 channel 数据  

