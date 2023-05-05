```bash
go mod init example.com/greetings  #创建一个模块 ,名为 example.com/greetings

go mod edit -replace example.com/greetings=../greetings 	#创建一个模块软链接,指向本地模块(而非远程) example.com/greetings ,其路径为 ../greetings

go get golang.org/x/example  # 添加依赖, 下载到 $GOBIN/pkg 目录
go clean -modcache # 清除所有下载的依赖

go mod tidy # The go mod tidy command adds missing module requirements for imported packages and removes requirements on modules that aren't used anymore.

go run . # 编译并运行(不生成可执行文件)

go test # 执行单元测试

go build # 编译可执行文件

go install # 将可执行文件放入 /Users/light/go/bin 目录
```
 
## 栈上分配还是堆区分配

golang 中数据分配在栈还是堆是由编译器决定的. 

栈上数据(栈帧)随着方法的执行与退出而push、pop, 这决定了其只能存放生命周期与方法一致的数据.   
当数据的生命周期超出了方法(发生指针逃逸), 则只能在堆上分配. 

由于栈的大小有限, 决定了其不适合存放过大的数据, 因此对于较大的数据, 将被分配在堆区 (即使并没有发生指针逃逸)

> https://go.dev/doc/faq#stack_or_heap
> https://cloud.tencent.com/developer/article/1890639

## 传参、返回值 用指针还是值

值传递会发生拷贝,但是分配和回收开销很低(只需要更改栈顶指针); 而指针传递不需要拷贝,但是有分配内存和回收内存的开销.

- 如果是有状态的数据, 那么使用指针传递
- 如果数据较大, 那么使用指针传递(编译器会自动将大的值传递改为引用传递)
- 其他情况,可使用值传递  

