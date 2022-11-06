


### 常量

- golang 当中是有常量的，golang 当中的常量一样用**驼峰标准，首字母大写**。
比如我们起一个常量叫做 app_env，表示当前 app 运行的环境，我们必须要这样定义：

```go
const AppEnv = "env"
```






- 如果是枚举类型 的常量，需要先创建相应类型：

```go
type Scheme string
  
const (
    HTTP  Scheme = "http"
    HTTPS Scheme = "https"
)
```









- 如果模块的功能 较为复杂、常量名称 容易混淆 的情况下，为了更好地区分 枚举类型，可以使用 完整的前缀：

  ```go
  type PullRequestStatus int
  
  const (
      PULL_REQUEST_STATUS_CONFLICT PullRequestStatus = iota
      PULL_REQUEST_STATUS_CHECKING
      PULL_REQUEST_STATUS_MERGEABLE
  )
  ```




运算符 优先级，结合括号 明确 运算顺序


状态、数字等，避免 硬编码，使用常量、枚举等方式