


文件名 应一律使用小写， 不同单词之间 用下划线分割, 不用驼峰式，
命名应尽可能地见名知意。看见文件名 就可以知道这个文件下 的大概内容,其中测试文件以 _test.go 结尾。

- 测试单元
  文件名 *test.go* 或者 文件名平台_test.go。
  例： path_test.go, path_windows_test.go

- 版本区分
  文件名_版本号等。
  例：trap_windows_1.4.go

- 平台区分
  文件名_平台。
  例： file_windows.go, file_unix.go
  可选为windows, unix, posix, plan9, darwin, bsd, linux, freebsd, nacl, netbsd, openbsd, solaris, dragonfly, bsd, notbsd， android，stubs

- CPU 类型区分, 汇编 用的多
  文件名_(平台:可选)_CPU类型.
  例：vdso_linux_amd64.go
  可选为 amd64, none, 386, arm, arm64, mips64, s390,mips64x,ppc64x, nonppc64x, s390x, x86,amd64p32


go build 的时候 会选择性地 编译以系统名 结尾的文件(linux、darwin、windows、freebsd)。例如Linux(Unix)系统下编译只会选择array_linux.go文件，其它系统命名后缀文件全部忽略。


在 xxx.go 文件的文件头上 添加 `// + build !windows (tags)`，可以选择在 windows 系统下面不编译
