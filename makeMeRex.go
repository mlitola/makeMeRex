package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    var inputPattern string
    var regex string

    flag.StringVar(&inputPattern, "value", "", "input pattern for which regex is generated,\nfor whitespaced input use quotation marks")
    fixedLength := flag.Bool("fixed", false, "is input pattern and length fixed")
    onePerLine := flag.Bool("s", false, "single occurrance per line")
    info := flag.Bool("i", false, "show regular expression common rules")
    
    flag.Parse()

    if *info == true {
        printRegularExpressionInfo()
    }

    if len(inputPattern) == 0 {
        if *info == false {
            fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
            flag.PrintDefaults()
        }
        os.Exit(1)
    }

    regex = AnalyzeAndMakePattern(inputPattern, fixedLength, onePerLine)
    
    fmt.Println("\nGenerated regex: ", regex)
    fmt.Println("\n◇ Golang: regexp.Compile(\"" + regex + "\")")
    fmt.Println("\n◈ JavaScript: new RegExp('" + regex + "')")
    fmt.Println("\n◆ Kotlin: new Regex(\"" + regex + "\")\n")
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
    fmt.Println("[abc]  character is a, b or c")
    fmt.Println("[a-z]  character is in range of a to z")
    fmt.Println("[A-Z]  character is capital letter from A to Z")
    fmt.Println("[0-9]  character is a number from 0 to 9")
    fmt.Println("[^abc] anything but a, b or c")
    fmt.Println("\\w     is word")
    fmt.Println("\\d     is digit")
    fmt.Println("\\s     is white space")
    fmt.Println("\\n     is new line")
    fmt.Println("\\t     is tab")

    fmt.Println("")
}
