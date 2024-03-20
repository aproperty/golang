

 # Go Mod 命令
```
init        initialize new module in current directory (再当前文件夹下初始化一个新的module, 创建go.mod文件)
vendor      make vendored copy of dependencies (将依赖复制到vendor下)

download    download modules to local cache (下载依赖的module到本地cache))
```

edit go.mod from tools or scripts (编辑go.mod文件)  
`go mod edit -require="github.com/objcoding/testmod@v1.0.1"`

add missing and remove unused modules:  
`go mod tidy`  
        



```
graph       print module requirement graph (打印模块依赖图))
verify      verify dependencies have expected content (校验依赖)
why         explain why packages or modules are needed 
```




# 设置 GO111MODULE
可以用环境变量 GO111MODULE 开启或关闭模块支持，它有三个可选值：off、on、auto，默认值是 auto。

`GO111MODULE=off` 无模块支持，go 会从 GOPATH 和 vendor 文件夹寻找包。
`GO111MODULE=on` 模块支持，go 会忽略 GOPATH 和 vendor 文件夹，只根据 go.mod 下载依赖。
`GO111MODULE=auto` 在 GOPATH/src 外面且根目录有 go.mod 文件时，开启模块支持。

在使用模块的时候，GOPATH 是无意义的，不过它还是会把下载的依赖储存在 GOPATH/pkg/mod 中，也会把`go install`的结果放在`$GOPATH/bin`中。Mod Cache 路径在`GOPATH/pkg/mod/cache`下面。




# Go Mod 使用
* 在一个新的项目中，需要执行`go mod init`来初始化创建文件`go.mod`，`go.mod`中会列出所有依赖包的路径和版本。
* go mod vendor 命令可以在项目中创建 vendor 文件夹将依赖包拷贝过来。
* go mod download 命令用于将依赖包缓存到本地 Cache 起来。
* 显示所有 Import 库信息：`go list -m -json all`


# 常见问题
启用 Go 模块以后，使用 `go get xxx` 时会报错提示 `go: cannot find main module; see 'go help modules'`，因为没有找到 go.mod 文件，所以会报错。你只要在项目根目录下生成一个 go.mod 文件就可以了。

如何在 Go 模块里使用本地依赖包？首先在项目的 go.mod 文件的 `require` 处添加依赖包，然后在 `replace` 处添加替换本地依赖包(路径要处理妥当)。比如：
```
require (
    mytest v0.0.0
)
replace (
    mytest v0.0.0 => ../mytest
)
```

如果在下载一些墙外的包可设置代理
```
export GO111MODULE=on
export GOPROXY=https://goproxy.io
```

---
---
---
对于 Go1.13 的 Module，我们可以通过 GOPROXY 来控制代理，以及通过 GOPRIVATE 控制私有库不走代理。
设置 GOPROXY 代理：  
`go env -w GOPROXY=https://goproxy.cn,direct`

设置 GOPRIVATE 来跳过私有库，比如常用的 Gitlab 或 Gitee 或 Gitea，中间使用逗号分隔：  
`go env -w GOPRIVATE=.gitlab.com,.gitee.com,*.gitea.com`

如果在运行 `go mod vendor` 时，提示 `Get https://sum.golang.org/lookup/xxxxxx: dial tcp 216.58.200.49:443: i/o timeout`，则是因为 Go 1.13 设置了默认的 `GOSUMDB=sum.golang.org`，这个网站是被墙了的，用于验证包的有效性，可以通过如下命令关闭：  
`go env -w GOSUMDB=off`

