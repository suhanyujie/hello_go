>* 原文来自：https://github.com/suhanyujie/hello_go/blob/master/notes/2021/go_file_api.md
>* 文章标题：【go 笔记】go 中的 文件 api
>* 作者：[suhanyujie](https://github.com/suhanyujie)
>* 标签：go
>* tip：如果异常，还请指正~

我们知道 go 语言是跨平台的，可以在不同平台上编译出相对应的可执行文件。但因为不同的操作系统，特性有所不同，对 api 的影响也各有差异。
这边笔记主要探讨的是 go 中文件接口的一些问题，在不同操作系统上，有一些差异。

在 go 中，我们可以通过如下方式获取文件的 mode 值：

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	execShowFileStat()
}

func execShowFileStat() {
	arg := os.Args[1]
	fd, _ := os.Stat(arg)
	fmt.Printf("file %s mod is: %o", arg, fd.Mode() & os.ModePerm)
}
```

```shell
$ go run go_file_api.go /etc/hosts
file D:/program/Git/etc/hosts mod is: 666
```



通过源码可以看到 File 类型的声明：

```go,no-exec
// File represents an open file descriptor.
type File struct {
	*file // os specific
}
```


## 参考
* https://fasterthanli.me/articles/i-want-off-mr-golangs-wild-ride