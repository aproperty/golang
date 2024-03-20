


## import 规范

import 在多行的情况下，**goimports会自动帮你格式化**，
但是我们这里 还是规范一下 import 的一些规范，如果你在一个文件里面引入了一个 package，还是建议采用如下格式：

```go
import (
    "fmt"
)
```

如果你的包引入了三种类型的包，
标准库包，
程序内部包，
第三方包，
建议采用如下方式进行组织你的包：

```go
import (
    "encoding/json"
    "strings"

    "myproject/models"
    "myproject/controller"
    "myproject/utils"

    "github.com/astaxie/beego"
    "github.com/go-sql-driver/mysql"
)   
```

有顺序的引入包，
不同的类型采用空格分离，
第一种实标准库，
第二是项目包，
第三是第三方包。










在项目中不要使用相对路径引入包：

```go
// 这是不好的导入
import “../net”

// 这是正确的做法
import “github.com/repo/proj/src/net”
```
