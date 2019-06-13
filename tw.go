package main

import (
    "os"
    "./tw_cli"
)


func main()  {
   args := os.Args[1:]
   tw_cli.Parse(args)
}