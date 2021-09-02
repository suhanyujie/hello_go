>* 原文来自：[suhanyujie](https://github.com/suhanyujie/hello_go)
>* 原文作者：[suhanyujie](https://www.github.com/suhanyujie)

事实上，这个标题有点标题党了，所谓“不要用 `std::process::exit`” 是指在特定场景下的。

在很多编程语言中，都具备 exit，比如 PHP 中的 `exit()`，Go 中的 `os.Exit()`，Rust 中也不例外，有 `std::process::exit`。它们的作用都是推出当前进程。但在具体的行为上略有差异。

以 PHP 的 `exit()` 为例，其[文档](https://www.php.net/exit)上描述：
>输出一个消息并且退出当前脚本。
>中止脚本的执行。 尽管调用了 exit()， Shutdown函数 以及 object destructors 总是会被执行。

结合上面的描述，用一个简单实例看下效果：

```php
<?php
$someConn = "";
function shutdownHandle() {
	// unset $someConn
	global $someConn;
	var_dump($someConn);
	unset($someConn);
	var_dump($someConn);// Warning: Undefined variable $someConn
	echo "exec shutdownHandle...\n";
}

register_shutdown_function(fn()=>shutdownHandle());

exit(0);
``

输出结果：

```powershell
PS D:\tech> php D:\tech\tmp\t1.php
string(0) ""
PHP Warning:  Undefined variable $someConn in D:\tech\tmp\t1.php on line 8

Warning: Undefined variable $someConn in D:\tech\tmp\t1.php on line 8
NULL
exec shutdownHandle...
```

通过执行 PHP 脚本，可以看出，脚本调用了 exit，但是总是会执行通过 register_shutdown_function 注册的处理函数。

### go 中的 exit
go 中的 Exit 函数位于标准库的 os 包中，[文档](https://pkg.go.dev/os#Exit)中可以看到其详细描述：
>* Exit causes the current program to exit with the given status code. Conventionally, code zero indicates success, non-zero an error. The program terminates immediately; deferred functions are not run.
>* Exit 使当前程序推出并给出推出状态码。按照管理，code 为 0 表示成功；非 0 表示错误。程序会立即终止，延迟函数也不会执行。
>* For portability, the status code should be in the range [0, 125].
>* 为了更好的兼容性，状态码应该在 [1, 125] 范围内。

可以看出 go 在当前的函数栈中调用 Exit 时，延迟函数不会执行。这样在一定场景下会导致资源泄露。

### Rust 中的 exit


## 参考
* 标准库文档 https://doc.rust-lang.org/std/process/fn.exit.html
* https://old.reddit.com/r/rust/comments/peovws/using_simd_acceleration_in_rust_to_create_the/?
* https://stackoverflow.com/questions/39228685/when-is-stdprocessexit-o-k-to-use

