# makeMeRex
A command line tool for making regular expression search pattern(s).

## Description
Another hobby project to study the great and mighty RegEx. The goal is to keep the tool simple and to generate regular expression with proper command line parameters.

### Requirements:

- Go 10.0.0+

### Build
```
# set GOPATH i.e. here project in $HOME/go/src/makeMeRex
export GOPATH=$HOME/go/src

cd $HOME/go/src/makeMeRex

# then compile a binary
go build
```

### Usage
```
# for usage info run with flag --h or --help
./makeMeRex --help

Usage of ./makeMeRex:
  -fixed
    	is input pattern and length fixed
  -rules
    	show regular expression common rules
  -single
    	single occurrance per line
  -value string
    	input pattern for which regex is generated,
    	for whitespaced input use quotation marks, 
    	for wildcard use * symbol
```
```
# to see used regex rules run with flag -rules
./makeMeRex -rules

Using the following Regular Expression rules:
^      match to beginning of line
$      match to end of line
{n}    exactly n item(s)
{n,}   n or more items
{n,m}  between n to m items
?      0 or 1 occurrences
*      0 or more occurrences
+      1 or more occurrences
.      any but new line
[abc]  character is a, b or c
[a-z]  character is in range of a to z
[A-Z]  character is capital letter from A to Z
[0-9]  character is a number from 0 to 9
[^abc] anything but a, b or c
\w     is word
\d     is digit
\s     is white space
\n     is new line
\t     is tab
```
```
# example execution
./makeMeRex -value="+358 40 1234567" -fixed

Generated regex:  +[0-9]{3}\s{1}[0-9]{2}\s{1}[0-9]{7}

◇ Golang: regexp.Compile("+[0-9]{3}\s{1}[0-9]{2}\s{1}[0-9]{7}")
◈ JavaScript: new RegExp('+[0-9]{3}\s{1}[0-9]{2}\s{1}[0-9]{7}')
◆ Kotlin: new Regex("+[0-9]{3}\s{1}[0-9]{2}\s{1}[0-9]{7}")
```