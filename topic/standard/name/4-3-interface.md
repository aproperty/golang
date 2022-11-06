
- 接口 命名规则：单个函数 的 接口 名 以 ”er” 作为后缀，如 Reader,Writer。  
接口的实现 则去掉“er”。

  ```go
  type Reader interface {
          Read(p []byte) (n int, err error)
  }
  ```


两个函数 的 接口 名 综合两个函数名

```go
  type WriteFlusher interface {
      Write([]byte) (int, error)
      Flush() error
  }
```



三个以上函数的接口名，抽象这个接口的功能，类似于结构体名

  ```go
  type Car interface {
      Start([]byte) 
      Stop() error
      Recover()
  }
  ```
