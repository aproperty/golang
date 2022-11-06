
### 变量命名

- 自注释，避免单个字符等方式命名；否则加注释

- Go 语言当中规定了 我们应该使用 **驼峰标准** 来命名变量，不能使用_。多个变量申明放在一起。

  ```go
  var (
      Found bool
      count int
  )
  ```




- 变量名称一般遵循驼峰法，但遇到特有名词时，需要遵循以下规则：

  - 如果变量为 私有，且特有名词 为首个单词，则使用小写，如 apiClient
  - 其它情况都应当使用该名词原有的写法，如 APIClient、repoID、UserID
  - 错误示例：UrlArray，应该写成urlArray或者URLArray









  下面列举了一些常见的特有名词：

  ```go
  // A GonicMapper that contains a list of common initialisms taken from golang/lint
  var LintGonicMapper = GonicMapper{
      "API":   true,
      "ASCII": true,
      "CPU":   true,
      "CSS":   true,
      "DNS":   true,
      "EOF":   true,
      "GUID":  true,
      "HTML":  true,
      "HTTP":  true,
      "HTTPS": true,
      "ID":    true,
      "IP":    true,
      "JSON":  true,
      "LHS":   true,
      "QPS":   true,
      "RAM":   true,
      "RHS":   true,
      "RPC":   true,
      "SLA":   true,
      "SMTP":  true,
      "SSH":   true,
      "TLS":   true,
      "TTL":   true,
      "UI":    true,
      "UID":   true,
      "UUID":  true,
      "URI":   true,
      "URL":   true,
      "UTF8":  true,
      "VM":    true,
      "XML":   true,
      "XSRF":  true,
      "XSS":   true,
  }
  ```
