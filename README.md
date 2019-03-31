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

# to see used regex rules run with flag -i
./makeMeRex -i
```