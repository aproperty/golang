

### 注释

注释 可以 帮我们很好的 完成文档的工作，写得好 的注释 可以方便我们以后的维护。  
详细的如何写注释可以参考：  
http://golang.org/doc/effective_go.html#commentary

- // 与注释内容间有一个空格

- **所有的函数 以及结构体头部 必须要写注释**，注释的规范是**名称加上说明**，只有通过 golint 规范检测的代码 才可以提交发布。

  ```go
  // HelloWorld print hello world
  func HelloWorld() {
      fmt.Println("Hello World")
  }
  ```

- 协议字段 的注释

- 注释 应该 表明 当前代码 对应的业务场景，特殊处理 的原因







- bug 注释

  针对代码中 出现的 bug，可以采用 如下教程 使用 特殊的注释，在 godocs 可以做到 注释高亮：

  ```go
  // BUG(astaxie):This divides by zero. 
  var i float = 1/0
  ```

  http://blog.golang.org/2011/03/godoc­documenting­go­code.html
