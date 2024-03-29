# Learning GoLang

Back in 2016 I started learning GoLang.  I wrote my notes in the repo [golang_cheet_sheet](https://github.com/cbitter78/golang_cheat_sheet).  Sense that time I have busy with a few other things.  Its now time (Jan 2022) to refresh my GoLang skills.

I am working through the this [udemy](https://www.udemy.com/course/learn-how-to-code/) corse with a few frends.   This repo is where I will keep my notes and example projects.

## Installing (Mac with ZSH)

First if you dont run [O My ZSH](https://github.com/ohmyzsh/ohmyzsh) you should!  It makes all this eaiser.

Just install go from [here](https://go.dev/doc/install).  Then set up your GOPATH.  I set these in my [.zprofile](https://zsh.sourceforge.io/Intro/intro_3.html) file

```shell
export GOPATH=~/workspace/go
```

In this case I have a folder named workspace where I do all my work and in that a folder named go where I will do golang devlopment.

From there I created src, pkg and bin.

```shell
cd $GOPATH
mkdir src pkg bin
```

## Editor (IDE)

Just use Visual Stuidos Code.  You can add the [Go Extension](https://marketplace.visualstudio.com/items?itemName=golang.go) and you are set.

## Package Management

Go modules are the package manager for go.  Its like pip for python.  The nice part is you no longer need to set up a $GOPATH/src folder and work out of there.  You can just keep your projects go.mod file up to date and you are good.  Do read about modules [here](https://go.dev/blog/using-go-modules) and [here](https://go.dev/ref/mod)


## Helpful links

- [The Go Programming Language Specification](https://go.dev/ref/spec)
- [Effective Go](https://go.dev/doc/effective_go): A guide to writing golang well.
- [Go Blog](https://go.dev/blog/): Just AWESOME!
- [Go Doc (pkg.go.dev)](https://pkg.go.dev/): Docs on the standard lib and other packages.
- [GoLang Docs](https://pkg.go.dev/): Offical golang docs for the standard lib.
- [Go by example](https://gobyexample.com/): This site is awesome to see working go code to reinforce how to do somthing or in my case look it up when you have forgotten the syntax.
- [Go Forum](https://forum.golangbridge.org/)
- [Go Playground](https://go.dev/play): A place where you can write and share go code in a web browser.   Super handy to just try somthing out or to share working code with someone.
- [Go Koans](https://github.com/cdarwin/go-koans): Fun coding exercises in golang
- [Go Style and conventions used by Uber](https://github.com/uber-go/guide/blob/master/style.md#introduction): This is a super useful page for dos and donts.  Well worth a read then a re-read.
- [Common Go Mistakes when writing loops and go goroutines](https://github.com/golang/go/wiki/CommonMistakes)
- [Golang code review referance](https://github.com/golang/go/wiki/CodeReviewComments): This is a list of common code review comments left on pull requests for golang.  You can link to each do or do not listed here.  For example so in your reviews you can say somthing like this: "Your error string should not be cap.  See [Error Strings](https://github.com/golang/go/wiki/CodeReviewComments#error-strings)"  This helps others see why and learn.
- [Learning Golang with Todd McLeod git Repo](https://github.com/GoesToEleven/GolangTraining):  Todd teaches the [udemy](https://www.udemy.com/course/learn-how-to-code/) corse I started with.  There are some nice examples in his git repo.
- [Json to Go](https://mholt.github.io/json-to-go/): A tool to take json and generate a struct to marshel it to in go.
- [Go proverbs](https://go-proverbs.github.io/): A list of sayings by rob pike.  Totally worth a read.
- [Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)
- [Encapsulation or Data Hiding in Go](https://aaronzhuo.medium.com/encapsulation-in-go-328c44dcb362)
  
## Notes

[Short variable declarations](https://go.dev/tour/basics/10) is this guy `:=` he is super cute and useful.  It declairs and assigns with an implicit type.

[Empty Ninja Level go file](https://github.com/cbitter78/golang_learning/blob/830f70d688e0383cb54b9bc4e6216f6bf6937d33/examples/structs/ninja_level5.go)

[How to build your program](https://www.ardanlabs.com/blog/2015/09/composition-with-go.html):  Composition in go with William (Bill) Kennedy

## The next course I want to take

[Ultimate Go with Bill Kennedy](https://www.ardanlabs.com/ultimate-go/)

## Code Snippits

[Generate Random String 3 Ways](https://go.dev/play/p/mJ1evVtBRnT)

## Cool Packages

### Logging

- [logrus](https://github.com/sirupsen/logrus): Killer logger that lets you do formatting, colors, and plugins.
- [logrus plugin hook for ElasticSearch](https://github.com/sohlich/elogrus)
  

### Debuging

- [delve](https://github.com/go-delve/delve): Delve is a debugger for the Go programming language


### Call Stack

Sometimes you want to knwo where you are in the call stack and what file and line of code.  [go-stack](https://pkg.go.dev/github.com/go-stack/stack) can do this.