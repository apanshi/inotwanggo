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

### Go语言中一共有 25 个关键字

|1            |2           |3               |4           |5              |6            |7             |8           |
|-------------|------------|----------------|------------|---------------|-------------|--------------|------------|
| ***if***    | ***else*** | ***switch***   | ***case*** | ***default*** | ***break*** | ***return*** | ***goto*** |
| fallthrough | ***for***  | ***continue*** | type       | ***struct***  | var         | ***const***  | map        |
| func        | interface  | range          | import     | package       | defer       | go           | select     |
| chan        |

- `var` 和 `const` ：变量和常量的声明， `var varName type` 或者 `varName : = value`
- `package` and `import` ：导入
- `func` ：用于定义函数和方法
- `return` ：用于从函数返回
- `defer someCode` ：在函数退出之前执行
- `go` ：用于并行
- `select` ：用于选择不同类型的通讯
- `interface` ：用于定义接口
- `struct` ：用于定义抽象数据类型
- `break`、`case`、`continue`、`for`、`fallthrough`、`else`、`if`、`switch`、`goto`、`default` ：流程控制
- `chan` ：用于 channel 通讯
- `type` 用于声明自定义类型
- `map` ：用于声明 `map` 类型数据
- `range` ：用于读取 `slice`、`map`、`chan` 数据

#### Go 语言中除了关键字以外，还有 36 个 `预定义标识符`

##### 内建常量 4 个

| 内建常量 |       |      |     |
|----------|-------|------|-----|
| true     | false | iota | nil |

##### 内建类型 20 个

| 內建类型 |           |            |         |
|----------|-----------|------------|---------|
| int      | int8      | int16      | int32   |
| int64    | uint      | uint8      | uint16  |
| uint32   | uint64    | uintptr    | float32 |
| float64  | complex64 | complex128 | bool    |
| byte     | rune      | string     | error   |

##### 内建函数 12 个

| 內建函数 |       |         |         |
|----------|-------|---------|---------|
| make     | len   | cap     | new     |
| append   | copy  | delete  | real    |
| imag     | panic | recover | complex |

### Go 数组和切片

数组是具有 **固定长度** 且拥有 0 个或多个相同数据类型元素的合集。

切片是相同类型元素的 **可变长度** 的集合，通常表示为 `[]type`，同一个切片内的元素数据类型都是同一个 `type`，**切片和数据的最大区别是没有长度**。

#### 数组

- 数组的索引 **Go 数组的索引是从 0 开始而不是从 1 开始的。**

#### 切片

切片的构成：1. 指向底层数组中某个元素的指针（指向数组的第一个从切片访问元素）； 2. 长度（length/len 切片的元素个数）； 3. 容量（capacity/cap 为切片分配的存储空间）。

- 切片使用 `len()` 获取长度，`cap()` 获取容量。
- 切片原则：**左闭右开（左含右不含）**。

#### `switch` 语句

`switch` 从上到下执行，直到匹配后，**不需要加 `break`**。

`switch` 语句还有一个特殊的关键字 `fallthrough` : **用于执行下一个分支的代码**。在 C、Java、JavaScript 中，**下降**是 `switch` 语句的各个分支的默认行为，但是 Go 中，用户必须显式的调用 `fallthrough` 关键字才可以引发下降。

**`case` 后面不仅仅可以放常量，还可以放变量和表达式。**

#### 自增自减只有 `++` `--` 语法

Golang 的设计者认为 C 语言中的两种自增自减（例如：`++i` `i++`）写法会引起歧义。所以强制规定 `++` `--` 只能写在变量后面。

**`++` `--` 是语句，不是表达式，必须独立占一行。**

#### 值传递引用

##### 值传递

**值类型**有: `int`系列、`float`系列、`bool`、`string`、`数组`、`结构体`：

1. 值类型通常在栈中分配存储空间
2. 值类型作为函数参数传递, 是拷贝传递
3. 在函数体内修改值类型参数, 不会影响到函数外的值

##### 指针引用

**引用类型**有: 指针、`slice`、`map`、`channel`：

1. 引用类型通常在堆中分配存储空间
2. 引用类型作为函数参数传递,是引用传递
3. 在函数体内修改引用类型参数,会影响到函数外的值

### `iota`

`iota`，特殊常量，可以认为是一个可以被编译器修改的常量。

`iota` 在 `const` 关键字出现时将被重置为 **0** （`const` 内部的第一行之前），`const` 中每新增一行常量声明将使 `iota` 计数一次（`iota` 可理解为 `const` 语句块中的行索引）。

### Go 类型

#### 基础类型

**Go 的所有类型没有隐式转换，必须手动转换**，比如：

```golang
var a int8 = 6
var b int32
b = a        //错误，不能把 int8 类型赋值给 int32 类型
b = int32(a) //正确，可以把 int8 类型赋值给 int32 类型
```

```bash
布尔类型： bool(只能设置 true 和 false 两个值)
整型： int(32 或 64 位) int8 int16 rune(int32 的别名，保存中文字符) int64 unit(与 int 一样大小) byte(uint8 的别名，保存英文字符) uint16 uint32 uint64 uintptr(无符号整型，用于存放一个指针) 等
浮点类型： float32 float64
复数类型： complex64 complex128
字符串： string(不可更改)
字符类型： rune(int32 的别名，表示 Unicode 码点) byte
错误类型： error
```

#### 复合类型

```bash
指针（pointer）
数组（arrry）
切片（slice）
字典（map）
结构体（struct）

函数类型
接口（interface）
通道（chan）
```

#### 数据类型定义

##### 1. 标准格式

```go
var 变量名 变量类型
var a int
var b string = "hello world"
```

##### 批量格式

```go
var (
    a int
    b string
    c []float32
    d func() bool
    e struct {
        x int
    }
)
```

##### 简短格式

```go
名字 := 表达式
c, d := 1, "hello world"
```

简短模式（short variable declaration）有以下限制：

1. 定义变量：同时显式初始化；
2. 不能提供数据类型；
3. 只能用于函数内部。

### Go 流程控制

**Go 没有 `while` 循环**

```bash
`if` `else` 和 `else if`
`switch` `case` `fallthrough` 和 `select`(用于通信的 `switch` 语句，每个 case 必须是通信操作：接收或发送，随机执行一个可运行的 **case**)
for （`for init; condition; post { }` `for condition { } ` `for { }` `for-each range`(可以对 `slice` `map` `数组` `字符串` 等迭代循环)）
goto
```

#### `fo...range` 循环

```go
for 索引, 值 := range 被遍历数据{
}
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
