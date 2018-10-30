# HowToDo

**HowoToDo is Go port of [howdoi](https://github.com/gleitz/howdoi "gleitz/howdoi")**

## Abstract

Can you come up with a flawless `tar` command in a second? How about `curl` with some gibberish headers? Do you know how to write a complete `ps`, `rsync`, `ssh`, and `sftp` with all the command line flags you need? **Well, now you know HowToDo, the instant stackoverflow via command line!**

You don't open up a browser, type `www.stackoverflow.com`, and search `tar flags`.
Instead, simply run `howtodo tar flags` on your command line and there you go!

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
1. Compile the code
   ```bash
   $ go install github.com/TipsyPixie/howtodo
   ```
1. Ask em' anything!
   ```bash
   $ $GOPATH/bin/howtodo go handle exception
   ```
   ```go
   https://stackoverflow.com/questions/25025467/catching-panics-in-golang
   
   func main() {
       if len(os.Args) != 2 {
            fmt.Printf("usage: %s [filename]\n", os.Args[0])
            os.Exit(1)
       }
       file, err := os.Open(os.Args[1])
       if err != nil {
           log.Fatal(err)
       }
       fmt.Printf("%s", file)
   }
   ...
   ```

## Usage
```bash
$ howtodo -h
```
```bash
Usage of howtodo:
  -a    Show full content of answer
  -l    Show only link to answer
  -v    Show version
```
