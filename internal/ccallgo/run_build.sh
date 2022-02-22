rm -rf bin libs
mkdir bin libs
go build -o ./libs/goforc.so -buildmode=c-shared goforc.go
gcc ccallgo.c -o ./bin/ccallgo ./libs/goforc.so
./bin/ccallgo
