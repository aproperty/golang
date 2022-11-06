
## new 和 make 的区别？

都是 内建函数  

new 函数定义：`func new(Type) *Type`  
作用 是 初始化 一个 指向类型 的 指针(*T)，  使用 new 函数来分配空间。
返回值 是 指向 这个 新分配的零值 的指针。  


make 函数定义：`func make(Type, size IntegerType) Type`   第一个参数是一个类型，第二个参数是长度，  返回值是一个类型    
作用是为 slice，map 或 chan 初始化 并返回 引用(T)。  
make 仅用于 创建 Slice, Map 和 Channel，并且返回类型 是 T（不是T*） 的 一个初始化的（不是零值） 的 实例。  