#!/bin/bash

# Version 1, normal build by cmd
#rm -rf bin libs
#mkdir bin libs
#go build -o ./libs/goforc.so -buildmode=c-shared goforc.go
#gcc ccallgo.c -o ./bin/ccallgo ./libs/goforc.so
#./bin/ccallgo

# Version 2, build by premake
go build -o ./libs/goforc.so -buildmode=c-shared goforc.go
go build -o ./libs/libgoforc.so -buildmode=c-shared goforc.go
rm -rf libs/goforc.so libs/libgoforc.h
premake5 gmake
sudo bash -c "echo '/mnt/e/selfDatas/self_todo_test/tobegoer/internal/ccallgo/libs'>>/etc/ld.so.conf"
sudo ldconfig
make clean && premake5 gmake && make && ./bin/Debug/ccallgo

#rm -rf *.make obj bin Makefile libs
