package test

import (
    "fmt"
    "main/internal/stgo"
    "main/internal/tforgo"
    "testing"
)

func TestMain(t *testing.T) {
    fmt.Println("Test for package Main")
    tforgo.Tmain()
    stgo.Smain()
}
