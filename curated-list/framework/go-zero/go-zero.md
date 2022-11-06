
- [go-zero](https://github.com/tal-tech/go-zero) - 
A web and rpc framework. 
It's born to ensure the stability of the busy sites with resilient design. 
Builtin goctl greatly improves the development productivity.



go-zero is a web and rpc framework written in Go.   
It's born to ensure the stability of the busy sites with resilient design.   
Builtin goctl greatly improves the development productivity.  




go-zero contains   
simple API description syntax and   
code generation tool called `goctl`.   

You can generate Go, iOS, Android, Kotlin, Dart, TypeScript, JavaScript from .api files with `goctl`.



## go-zero

go-zero 是一个 集成了 各种工程实践 的 web 和 rpc 框架。  
通过 弹性设计 保障了 大并发 服务端 的 稳定性，经受了 充分的 实战 检验。

go-zero 包含 极简的 API 定义 和 生成工具 goctl，
可以根据 定义的 api 文件 一键生成 Go, iOS, Android, Kotlin, Dart, TypeScript, JavaScript 代码，并 可 直接运行。

主要特点：
- 强大的 工具支持，尽可能少 的代码编写
- 极简 的接口
- 完全兼容 net/http
- 支持 中间件，方便扩展
- 高性能
- 面向 故障编程，弹性设计
- 内建 服务发现、负载均衡
- 内建 熔断、降载，且 自动触发，自动恢复
- API 参数 自动校验
- 超时 级联控制
- 自动 缓存控制
- 链路跟踪、统计报警 等
- 高并发 支撑，稳定 保障了 晓黑板 疫情期间 每天的 流量洪峰



Advantages of go-zero:

* improve the stability   of the services    with tens of millions of daily active users
* builtin chained timeout control,     concurrency control,    rate limit,    adaptive circuit breaker, adaptive load shedding,      even no configuration needed
* builtin middlewares     also can be integrated into your frameworks
* simple API syntax,    one command to generate couple of different languages
* auto validate the request parameters from clients
* plenty of builtin microservice management      and concurrent toolkits




go-zero 是 好未来 自研的 集成了 各种 工程实践 的 微服务 web 和 rpc 框架。  
通过 弹性设计 保障了 大并发 服务端的稳定性，经受了充分的 实战检验，  
go-zero 提供了一个 完备的、性能可观、开发迅速的 web 框架，但又不像 其他框架一样 有过多的 束缚，开发者 也可以在里边 做自定义的发挥。



- 稳定性：内外同源，好未来内部 也使用的是 github 的版本。
- 工作效率：通过 goctl 工具 生成各种代码，开发人员 只需要 补充 业务逻辑，也可以生成 js、java、go、ts 接口，提高 对接效率。
- 内建 级联 超时控制、限流、自适应熔断、自适应降载 等微服务治理能力，无需 配置 和额外代码
- 微服务治理中间件  可无缝集成 到其它现有框架使用
- 自动校验 客户端 请求参数 合法性
- 大量 微服务治理和并发工具包

缺点：好未来 2020年 8月份 新开源的框架，目前大项目的案例 比较少



### 部分功能说明

#### 日志

go-zero 除了 通用的 日志功能之外，还支持

- 打印 耗时情况，可以在 需要计算耗时的地方 添加计时，比如 sql 执行时间，可以很轻松 统计出慢 sql 的相关信息；
- 链路追踪日志；
- 日志压缩；
- 日志保留天数；
- 日志输出间隔；



- info
- error
- server
- fatal
- slow
- stat



#### 请求链路

微服务架构中，调用链 可能很漫长，从 `http` 到 `rpc` ，又从 `rpc` 到 `http` 。  
而开发者 想了解 每个环节 的 调用情况及性能，最佳方案就是 **全链路跟踪**。  
追踪的方法 就是在一个请求开始时 生成一个自己的 `spanID` ，随着整个请求链路 传下去。  
go-zero 则 通过这个 `spanID` 查看整个链路的情况 和 性能问题。  

**代码结构**

- spancontext：保存链路的 上下文信息「traceid，spanid，或者是 其他想要传递的内容」
- span：链路中 的一个操作，存储时间 和 某些信息



#### 服务自适应降载

服务自适应降载 的目的 是保证系统 不被过量请求 拖垮，在保证系统稳定 的前提下，尽可能 提供更高的吞吐量。

**如何衡量系统负载**

- 是否处于 虚机或容器内，需要读取 cgroup 相关负载
- 用 1000m 表示 100%CPU，推荐使用 800m 表示系统 高负载

**机制设计**

- 计算 CPU 负载时  使用滑动平均 来降低 CPU 负载抖动  带来的不稳定
- 时间窗口 机制，用 滑动窗口机制 来记录之前时间窗口内 的 QPS 和 RT(response time)
- 当满足一定的条件 就拒绝该请求

参考：https://mp.weixin.qq.com/s/cgjCL59e3CDWhsxzwkuKBg