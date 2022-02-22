#include <iostream>

extern "C" {
    #include "hello.h"
}

int SayHelloCxx() {
    std::cout << "Hello, World! From Cxx for cgo to used !\n";
    return 0;
}