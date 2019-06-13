package main

import (
    "os"
    "./tw_cli/parser"
)


func main()  {
   args := os.Args[1:]
   parser.Parse(args)
}