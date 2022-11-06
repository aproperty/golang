
# Go to

1. `godef`       Go to definition          `go get -v github.com/rogpeppe/godef`   
3. `go-outline`  Go to symbol in file      `go get -v github.com/ramya-rao-a/go-outline`   
3. `go-symbols`  Go to symbol in workspace `go get -v github.com/acroca/go-symbols` 
2. `guru`        Find all references and Go to implementation of symbols    `go get -v golang.org/x/tools/cmd/guru`  
 


# Debug

5. `golint`      Linter                `go get -v golang.org/x/lint/golint`   
4. `dlv`         Debugging             `go get -v github.com/go-delve/delve/cmd/dlv` 



# Formatter

1. `goreturns`    Formatter    `go get -v github.com/sqs/goreturns`     
2. `goimports`    Formatter    `go get -v golang.org/x/tools/cmd/goimports`   


* gofmt：大部分的格式问题可以通过 gofmt 解决，gofmt 自动化格式代码，保证所有的 go 代码与官方推荐的格式保持一致，于是所有格式有关问题，都以 gofmt 的结果为准。
* goimports：此工具在 gofmt 的基础上增加了自动删除和引入包。
* go vet： vet 工具可以帮我们 静态分析 源码存在的各种问题，例如多余的代码，提前 return 的逻辑，struct 的 tag 是否符合标准等。编译前 先执行代码静态分析。
* golint：类似于 javascript 中的 jslint 的工具，主要功能就是检测代码中不规范的地方。



# Auto-completion

3. `gocode`        Auto-completion                    `go get -v github.com/mdempsky/gocode`   
4. `gocode-gomod`  Autocompletion, works with Modules `go get -v github.com/stamblerre/gocode` 
5. `gopkgs`        Auto-completion of unimported packages & Add Import feature  `go get -v github.com/uudashr/gopkgs/cmd/gopkgs`   
4. `fillstruct`    Fill structs with defaults  `github.com/davidrjenni/reftools/cmd/fillstruct`
4. `gotests`       Generate unit tests         `go get -v github.com/cweill/gotests/`   




  
 `gorename`  Rename symbols `go get -v golang.org/x/tools/cmd/gorename` 

 `gomodifytags` Modify tags on structs  `go get -v github.com/fatih/gomodifytags` 

 `impl` Stubs for interfaces  `go get -v github.com/josharian/impl`   

 `goplay` The Go playground  `go get -v github.com/haya14busa/goplay/cmd/goplay` 

 `godoctor` Extract to functions and variables   `go get -v github.com/godoctor/godoctor` 

  `gopls`     `go get -v golang.org/x/tools/cmd/gopls` 



go 的辅助工具的话，那在 vs code 中按 Ctrl + Shift + P，输入 `go install`，点击 Go:Install/Update Tools
