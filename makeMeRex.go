package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    var inputPattern string

    flag.StringVar(&inputPattern, "value", "", "input pattern for which regex is generated")
    fixedPattern := flag.Bool("fixed", false, "is input value fixed")
    onePerLine := flag.Bool("s", false, "single occurrance per line")
    info := flag.Bool("i", false, "show regular expression common rules")
    
    flag.Parse()

    if *info == true {
        printRegularExpressionInfo()
        os.Exit(1)
    }

    fmt.Println("value:", inputPattern)
    fmt.Println("fixed:", *fixedPattern)
    fmt.Println("s:", *onePerLine)
}

func printRegularExpressionInfo() {
    fmt.Println("\nUsing the following Regular Expression rules:")

    fmt.Println("\nAssertions and Quantifiers")
    fmt.Println("^      match to beginning of line")
    fmt.Println("$      match to end of line")
    fmt.Println("{n}    exactly n item(s)")
    fmt.Println("{n,}   n or more items")
    fmt.Println("{n,m}  between n to m items")
    fmt.Println("?      0 or 1 occurrences")
    fmt.Println("*      0 or more occurrences")
    fmt.Println("+      1 or more occurrences")

    fmt.Println("\nCharacters")
    fmt.Println(".      any but new line")
    fmt.Println("[abc]  a, b or c")
    fmt.Println("[a-z]  from a to z")
    fmt.Println("[^abc] anything but a, b or c")
    fmt.Println("\\w     is word")
    fmt.Println("\\d     is digit")
    fmt.Println("\\s     is white space")
    fmt.Println("\\n     is new line")
    fmt.Println("\\t     is tab")

    fmt.Println("")
}
