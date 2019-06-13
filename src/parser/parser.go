package parser

import (
    "fmt"
)

func Parse(args []string){
    if len(args) > 0 {
        switch args[0] {
        case "login":
            fmt.Println("try to login")
        case "get":
            fmt.Println("try to get")
        case "tweet":
            fmt.Println("try to tweet")
        default:
            fmt.Println("unknown argument")
        }
    }
    fmt.Println(args)
}