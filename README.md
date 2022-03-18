# 我并不想学习 [Go](https://go.dev/)

## 在 Debian 安装 [Go](https://go.dev/doc/install)

```shell
wget https://dl.google.com/go/go1.17.7.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.17.7.linux-amd64.tar.gz
echo "PATH=\$PATH:/usr/local/go/bin">>$HOME/.profile
source ~/.profile
go version
```

### `go` 多版本控制

1. 可以通过 [docker](https://www.docker.com/) 来管理；
2. 可以通过 [voidint/g：Golang Version Manager](https://github.com/voidint/g) 来管理，更方便，推荐。

### `go` 版本升级

go1.17.4 升级到 go1.18

#### 1. 源码升级

```shell
wget https://go.dev/dl/go1.18.src.tar.gz
tar -xvf go1.18.src.tar.gz
cd go/src
bash all.bash
cd ../../ 
mv go /usr/local/go
go version

# 1. /usr/local/go 就是原先使用的老版本的 go，如此完成覆盖更新；
# 2. 如果 go 根目录也改变了，则需要去 .bash_profile 中修改 GOPATH坏 境变量；
# 3. 如果打算安装两个（多个）版本，则不要覆盖老版本，并通过修改 GOPATH 在新老版本间切换。
```

1. 可能需要的报错

```shell
Building Go cmd/dist using /usr/local/go. (go1.17.4 linux/amd64)
./make.bash: eval: line 201: unexpected EOF while looking for matching `"'

报错的话，把 201-204 注释掉就可以的。
```

#### 2. release 包升级

```bash
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version
```


### `go` 命令报错

使用 `zsh` 的话假如报错： `zsh: command not found: go。`， 打开 `~/.zshrc`， 在原有 `export PATH` 增加如下配置：

```bash
# go 配置
export GOPATH="$HOME/go"
export PATH="$PATH:/usr/local/go/bin:$GOPATH/bin"
```

保存退出后，执行 `source ~/.zshrc` 生效。

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

### Tips

- Go 并不推荐使用多行注释，建议多用单行注释 `// 单行注释`，虽然有多行注释的语法 `/* 多行注释 *`。
- Go 中 `_` 单独作为标识符出现时，代表 `空标识符`，对应的值会被忽略。
- 输出十进制只能用 `%d`, 而不能像 C 语言那样通过 `%i`，如：`fmt.Printf("age = %d\n", age) //age = 33`。 `%b` 输出二进制。
- `%T` 输出值的类型，如：`fmt.Printf("num1 = %T\n", num1) // int，num1 := 10`。`%v` 打印所有类型数据。
- `defer` 语句常用于 释放资源、解除锁定 以及 错误处理等。
- 单个包中代码执行顺序如下：`main 包 -> 常量 -> 全局变量 -> init 函数 --> main 函数 -> Exit`。
- `nil` 是 `interface`、`function`、`pointer`、`map`、`slice` 和 `channel` 类型变量的默认初始值。
- 很多编程语言使用 `~` 作为一元按位取反（NOT）操作符，Go 重用 `^` XOR 操作符来按位取反。
- AND NOT `&^` 操作符，不同位才取 1。
- `string` 只读的采用 utf8 编码的字节 `slice`，`len()` 函数求出的长度不是**字符个数，是字节个数**。
- `rune` 是 `int32` 的**别名**，保存中文字符；`byte` 是 `uint8` 的别名，可变长度的字节 `slice`，保存**英文字符**。

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

数组的索引 **Go 数组的索引是从 0 开始而不是从 1 开始的。**

#### 切片

切片的构成：1. 指向底层数组中某个元素的指针（指向数组的第一个从切片访问元素）； 2. 长度（length/len 切片的元素个数）； 3. 容量（capacity/cap 为切片分配的存储空间）。

- 切片使用 `len()` 获取长度，`cap()` 获取容量。
- 切片原则：**左闭右开（左含右不含）**。
- 切片的索引不能为负数。

### `switch` 语句

`switch` 从上到下执行，直到匹配后，**不需要加 `break`**。

`switch` 语句还有一个特殊的关键字 `fallthrough` : **用于执行下一个分支的代码**。在 C、Java、JavaScript 中，**下降**是 `switch` 语句的各个分支的默认行为，但是 Go 中，用户必须显式的调用 `fallthrough` 关键字才可以引发下降。

**`case` 后面不仅仅可以放常量，还可以放变量和表达式。**

### 自增自减只有 `++` `--` 语法

Golang 的设计者认为 C 语言中的两种自增自减（例如：`++i` `i++`）写法会引起歧义。所以强制规定 `++` `--` 只能写在变量后面。

`++` `--` **是语句，不是表达式**，必须独立占一行。

#### `defer`

无论你在什么地方注册 `defer` 语句，它都会在所属函数执行完毕之后才会执行, 并且如果注册了多个 `defer` 语句，那么它们会按照**后进先出**的原则执行。

#### 值传递引用

#### 值传递

**值类型**有: `int`系列、`float`系列、`bool`、`string`、`数组`、`结构体`：

1. 值类型通常在栈中分配存储空间
2. 值类型作为函数参数传递, 是拷贝传递
3. 在函数体内修改值类型参数, 不会影响到函数外的值

#### 引用类型

**引用类型**有: 指针、`slice`、`map`、`channel`：

1. 引用类型通常在堆中分配存储空间
2. 引用类型作为函数参数传递,是引用传递
3. 在函数体内修改引用类型参数,会影响到函数外的值

#### 关键字 `iota`

`iota`，特殊常量，可以认为是一个可以被编译器修改的常量。

`iota` 在 `const` 关键字出现时将被重置为 **0** （`const` 内部的第一行之前），`const` 中每新增一行常量声明将使 `iota` 计数一次（`iota` 可理解为 `const` 语句块中的行索引）。

#### 变量覆盖检查

可使用 [vet](https://pkg.go.dev/cmd/vet) 工具来诊断变量覆盖，Go 默认不做覆盖检查，添加 `-shadow` 选项来启用：`go tool vet -shadow main.go`。

注意 vet 不会报告全部被覆盖的变量，可以使用 [`go-nyet`](https://github.com/barakmich/go-nyet) 来做进一步的检测：
`$GOPATH/bin/go-nyet main.go`

#### 运算符的优先级

除了位清除（bit clear）操作符，运算符优先级从高到低：

| Precedence | Operator                               |
|:----------:|:---------------------------------------|
|    5       | `*`  `/`  `%` `<<`  `>>` `&`  `&^`     |
|    4       | `+`  `-`  `\|` `^`                     |
|    3       | `==` `!=` `<` `<=`  `>`  `>=`          |
|    2       | `&&`                                   |
|    1       | `\|\|`                                 |

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
整型：
    int(32 或 64 位) int8 int16 rune(int32 的别名，保存中文字符) int64
    unit(与 int 一样大小) byte(uint8 的别名，保存英文字符) uint16 uint32 uint64
    uintptr(无符号整型，用于存放一个指针) 等
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
条件选择：`if` `else` 和 `else if`
循环控制：
    `switch` `case` `fallthrough`
    `select`(用于通信的 `switch` 语句，每个 case 必须是通信操作：接收或发送，随机执行一个可运行的 **case**)
    for
        `for init; condition; post { }`
        `for condition { } `
        `for { }`
        `for-each range`(可以对 `slice` `map` `数组` `字符串` 等迭代循环)）
    goto
```

#### `fo...range` 循环

```go
for 索引, 值 := range 被遍历数据{
}
```

### 函数

**函数，小写字母开头的函数只能在本包内可见，大写字母开头的函数才能被其他包使用。**

函数的声明和调用语法

```go
// 声明
// 关键字  函数名  形式参数  返回结果
   func    Intn    (n  int)  int {

   }

// 调用
//         包   函数名 实际参数
    num := rand.Intn(10)
```

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
