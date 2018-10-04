package main

import (
    "flag"
    "fmt"
)

func main() {
    linkOnly := flag.Bool("l", false, "Show only link to answer")
    showFullAnswer := flag.Bool("a", false, "Show full content of answer")
    showVersion := flag.Bool("v", false, "Show version")

    // Parse the flags to generate help text for -h flag
    flag.Parse()

    switch {
    case *linkOnly:
        fmt.Println("Show link only")
    case *showFullAnswer:
        fmt.Println("Show full answer")
    case *showVersion:
        fmt.Println("Howtodo v1.0.0")
    }
}
