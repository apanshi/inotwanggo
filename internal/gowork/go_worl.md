mkdir -p internal/gowork
cd internal/gowork
mkdir hello
cd hello
go mod init example.com/hello
go get golang.org/x/example
vim hello.go
go run example.com/hello
cd ..
go work init ./hello
cat go.work
go run example.com/hello
git clone https://go.googlesource.com/example
go work use ./example
cat go.work
touch example/stringutil/toupper.go
vim example/stringutil/toupper.go
vim hello/hello.go
go run example.com/hello
go build -o hello_test example.com/hello
./hello_test