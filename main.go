package main

import (
    "fmt"
    "plugin"
)

func main() {
    readPlugin()
}

func readPlugin() {
    fmt.Println("reading plugin ...")
    p, err := plugin.Open("lib_plugin.so")
    if err != nil {
        panic(err)
    }
    fn, err := p.Lookup("HelloWorld")
    if err != nil {
        panic(err)
    }
    fn.(func())()
}
