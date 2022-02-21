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
