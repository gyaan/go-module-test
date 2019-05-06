package main

import (
    "fmt"
    "rsc.io/quote"
    app1 "github.com/gyaan/go-module-test/cmd/app1"
    app2 "github.com/gyaan/go-module-test/cmd/app2"
)

func main() {
    testHello()
}
func testHello() {
    fmt.Println(quote.Hello())
    app2.TestApp2()
    app1.TestApp1()
}
