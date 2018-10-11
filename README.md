# HowToDo

**HowoToDo is a Go ported [howdoi](https://github.com/gleitz/howdoi "gleitz/howdoi")**

## Abstract

Can you come up with a flawless `tar` command in a second? How about `curl` with some gibberish headers? Do you know how to write a complete `ps`, `rsync`, `ssh`, and `sftp` with all the command line flags you need? **Well, now you know HowToDo, the instant stackoverflow via the command line!**   

## Requirements

* Go ~1.11.0
* Recommended OS as Windows 10 or Ubuntu 18.04

## Getting Started

1. Install Go compiler at [https://golang.org/dl/](https://golang.org/dl/ "Downloads - the Go Programming Language")
1. Check out the installation is completed
   ```bash
   $ go version
   ```
   ```bash
   go1.11.1 linux/amd64
   ```
1. Make a directory for Go projects
   ```bash
   $ mkdir -p $HOME/go
   ```
1. Set up GOPATH to your Go project directory
   ```bash
   $ export GOPATH="$HOME/go"
   ```
1. Get the source code
   ```bash
   $ go get github.com/TipsyPixie/howtodo
   ```
1. Go to your GOPATH and compile the code
   ```bash
   $ cd $GOPATH && go install github.com/TipsyPixie/howtodo
   ```
1. Wanna know how to slice an array in Python?
   ```bash
   $ $GOPATH/bin/howtodo -a python array slice
   ```
   ```bash
   https://stackoverflow.com/questions/509211/understanding-pythons-slice-notation
   
   It's pretty simple really:
   a[start:end] # items start through end-1
   a[start:]    # items start through the rest of the array
   a[:end]      # items from the beginning through end-1
   a[:]         # a copy of the whole array
   
   There is also the step value, which can be used with any of the above:
   a[start:end:step] # start through not past end, by step
   ...
   ```
