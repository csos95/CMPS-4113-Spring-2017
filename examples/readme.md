## Go examples


1. **hello_world** - hello world
2. **json** - parse a json file and print the message it contains
3. **fraction** - partial fraction package that shows how a package is structured and its use
4. **hello_world_server** - a server example that shows that a package can be split across multiple files
5. **template_server** - a server example that shows how to use a template
6. **stack** - using go examples to test a package

## Go Tour
The official Go introductory tutorial is available [here](https://tour.golang.org)

## How to Setup Go, Git, and Gogland and Open the Example Project

1. Download and install Go (1.7.4)
 * [Windows](https://storage.googleapis.com/golang/go1.7.4.windows-amd64.msi)
 * [OS X](https://storage.googleapis.com/golang/go1.7.4.darwin-amd64.pkg)
 * [Linux](https://storage.googleapis.com/golang/go1.7.4.linux-amd64.tar.gz)
 * [Source](https://storage.googleapis.com/golang/go1.7.4.src.tar.gz)
2. Download and install Gogland [Link](https://www.jetbrains.com/go/download)
3. Setup Git
 1. Download and install [Git](https://git-scm.com/)
 2. Using cmd run `git config --global user.email "you@example.com"` and `git config --global user.name "Your name"`
4. Set GOPATH environment variable:
 1. explorer > right click computer/this pc > properties > advanced system settings > environment variables > new system variable
 2. name: GOPATH value: `C:\Users\[username]\gocode`
5. Using cmd, run `go get github.com/csos95/CMPS-4113-Spring-2017` this will pull the repo and attempt to compile the go source.  
Because there are no go files in the root of the repo, it will say no source files found.  
This is ok, we just wanted to pull the repo and get the gopath file tree setup.  
If you want to see the files, they are in `C:\Users\\[username]\gocode\src\github.com\csos95\CMPS-4113-Spring-2017`
6. Open the example project
 1. Open Gogland and choose open
 2. Browse to `C:\Users\[username]\gocode\src\github.com\csos95\CMPS-4113-Spring-2017\examples` and click ok.
 3. On the top right, next to the run button, you can select which run configuration to use.
 4. Right now there are three example programs to choose from.
