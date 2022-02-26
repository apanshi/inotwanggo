# I not want [Go](https://go.dev/)

## 在 Debian 安装 [Go](https://go.dev/doc/install)

```shell
wget https://dl.google.com/go/go1.17.7.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.17.7.linux-amd64.tar.gz
echo "PATH=\$PATH:/usr/local/go/bin">>$HOME/.profile
source ~/.profile
go version
```

### 测试安装

```shell
go mod init hello
go mod tidy
```

编写 `hello.go` file, 代码如下

```go
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```

```shell
go build
./hello
```

## GO Notes

### Go 数组和切片

数组是具有 **固定长度** 且拥有 0 个或多个相同数据类型元素的合集。

切片是相同类型元素的 **可变长度** 的集合，通常表示为 `[]type`，同一个切片内的元素数据类型都是同一个 `type`，**切片和数据的最大区别是没有长度**。

#### 数组

- 数组的索引 **Go 数组的索引是从 0 开始而不是从 1 开始的。**

#### 切片

切片的构成：1. 指向底层数组中某个元素的指针（指向数组的第一个从切片访问元素）； 2. 长度（length/len 切片的元素个数）； 3. 容量（capacity/cap 为切片分配的存储空间）。

- 切片使用 `len()` 获取长度，`cap()` 获取容量。
- 切片原则：**左闭右开（左含右不含）**。


### `iota`

`iota`，特殊常量，可以认为是一个可以被编译器修改的常量。

`iota` 在 `const` 关键字出现时将被重置为 **0** （`const` 内部的第一行之前），`const` 中每新增一行常量声明将使 `iota` 计数一次（`iota` 可理解为 `const` 语句块中的行索引）。

### Go 类型

#### 基础类型

```bash
布尔类型： bool
整型： int8 byte(uint8 的别名) int16 int uint uintptr 等
浮点类型： float32 float64
复数类型： complex64 complex128
字符串： string
字符类型： rune  // int32 的别名，表示 Unicode 码点
错误类型： error
```

#### 复合类型

```bash
指针（pointer）
数据（arrry）
切片（slice）
字典（map）
结构体（struct）

接口（interface）
通道（chan）
```

### Go 流程控制

```bash
if else 和 else if
switch case 和 select
for range
goto
```

### 函数

**函数，小写字母开头的函数只能在本包内可见，大写字母开头的函数才能被其他包使用。**

### 并发编程

**不要通过共享内存来通信，应该通过通信来共享内存。**

#### `channel` 语法

```go
var chanName chan ElementType
```

`channel` 写入数据 `ch <- value`， `channel` 写入数据会导致程序阻塞，知道有其他 `goroutine` 从这个 `channel` 种读取数据。

`channel` 读取数据 `value := <- channel`。

带缓冲的 `channel`： `c := make(chan int, 1024)`， 大小为 1024 的 `int` 类型 `channel`，即使没有读取方，写入时也可以在缓冲区没满前一直写 `channel`。

关闭 `channel` : `close(ch)`。

**指定 CPU 使用数目** : `runtime.GOMAXPROCS(16)`，设置使用 16 个 CPU 内核。

### 网络编程

只需要调用 `net.Dail()` 即可：

1. TCP 链接 `conn, err = net.Dail("tcp", "192.168.0.12:2100)`；
2. UDP 链接 `conn, err = net.Dail("udp", "192.168.0.12:975")`;
3. ICMP 链接（使用协议名）`conn, err = net.Dail("ip4:icmp", "wwww.google.com")`;
4. ICMP 链接（使用协议编号）`conn, err = net.Dail("ip4:1", "10.0.0.3")`。

**建立连接后，就可以使用 `conn.Write()` 写数据 或者 `conn.Read()` 读数据。**

#### 其他函数

### 单元测试

在相同的目录中创建一个以 `_test.go` 结尾的新文件作为你的包源文件。 在该文件中，加入 `import "testing"` 并编写以下形式的函数：

```go
func TestFoo(t *testing.T) {
    ...
}
```

在该目录中运行 `go test`。该脚本会查找 `Test` 函数， 构建一个测试二进制文件并运行它。

### 杂记

1. `Go` 没有三元操作符 `?:`，可以用下面的代码代替：

```go
if expr {
    n = trueValue
} else {
    n = falseValue
}
```
