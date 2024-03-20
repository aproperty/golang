package waitgrouputil

import "fmt"

// 同步执行
func (w waitGroupWrapper) Wrap(handler func(params ...interface{}), params ...interface{}) {

	w.waitGroup.Add(1)
	defer w.waitGroup.Done()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error: ")
		}
	}()

	handler(params...)
}

// 启动 协程 执行
func (w waitGroupWrapper) WrapGoroutine(handler func(params ...interface{}), params ...interface{}) {

	w.waitGroup.Add(1)

	go func() {

		defer func() {
			if err := recover(); err != nil {
				fmt.Println("error: ")
			}
		}()
		defer w.waitGroup.Done()

		handler(params...)
	}()
}
