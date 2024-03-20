
# windows 系统下打包 Go 程序到 Linux 平台运行

# 一次性打包各个平台的可执行程序

以下内容摘抄自 https://github.com/mitchellh/gox

## Installation

To install Gox, please use `go get`. 
We tag versions so feel free to checkout that tag and compile.  
`go get github.com/mitchellh/gox`   

Run `gox -h` for help and additional information.




## Usage
If you know how to use `go build`, then you know how to use Gox. For example, to build the current package, specify no parameters and just call `gox`. Gox will parallelize based on the number of CPUs you have by default and build for every platform by default:
```
$ gox
Number of parallel builds: 4

-->      darwin/386: github.com/mitchellh/gox
-->    darwin/amd64: github.com/mitchellh/gox
-->       linux/386: github.com/mitchellh/gox
-->     linux/amd64: github.com/mitchellh/gox
-->       linux/arm: github.com/mitchellh/gox
-->     freebsd/386: github.com/mitchellh/gox
-->   freebsd/amd64: github.com/mitchellh/gox
-->     openbsd/386: github.com/mitchellh/gox
-->   openbsd/amd64: github.com/mitchellh/gox
-->     windows/386: github.com/mitchellh/gox
-->   windows/amd64: github.com/mitchellh/gox
-->     freebsd/arm: github.com/mitchellh/gox
-->      netbsd/386: github.com/mitchellh/gox
-->    netbsd/amd64: github.com/mitchellh/gox
-->      netbsd/arm: github.com/mitchellh/gox
-->       plan9/386: github.com/mitchellh/gox
```

Or, if you want to build a package and sub-packages:
```
$ gox ./...
...
```

Or, if you want to build multiple distinct packages:
```
$ gox github.com/mitchellh/gox github.com/hashicorp/serf
...
```

Or if you want to just build for linux:

`gox -os="linux"`  

Or maybe you just want to build for 64-bit linux:

`gox -osarch="linux/amd64"`  





分别 放在 不同的 文件夹:  
```
$ gox -output "{{.Dir}}_{{.OS}}_{{.Arch}}/swbatch"
```
