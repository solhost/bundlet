<p align="center">
  <img src="assets/logo.svg" alt="Bundlet logo" width="75">
</p>

<h1 align="center">
  Bundlet
</h1>

Bundlet is a simple compiler package for projects that allows you to distribute completely open-source, non compiled software. That allows the user to simply compile your project with the run of a script.

## Requirements

1. a Linux desktop
2. the GCC (GNU C Compiler) installed
3. Golang/Go installed

## Setup

To properly use Bundlet, edit [`worker.c`](./worker.c):

```c
#include <stdlib.h>

int main() {
  system("go build example.go"); // edit "example.go" with the name of your go file, eg; "src/your-software.go"
  return 0;
}
```

We recommend you put your source files in `src/`, however it's not a requirement.

[`bundlet.go`](./bundlet.go):

```go
package main

import (
  "fmt"
  "os"
  "os/exec"
  "time"
)

func main() {
  // make this variable the project being compiled;
  compile := "Example" // change "Example" to your project name, eg; "GoonCode"
  fmt.Println("Bundlet will compile:", compile)
  fmt.Println("in 5 seconds.")
  time.Sleep(5 * time.Second)
  cmd := exec.Command("./worker")
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  if err := cmd.Run(); err != nil {
    fmt.Println("Failed to run compiling worker,", err)
    os.Exit(1)
  }
}
```

## Shipping Your Project

We recommend removing:

1. [`worker.c`](./worker.c), as it is compiled as a seperate binary. *Do not do this if you are making a fork of Bundlet.*
2. [`bundlet_bundler.c`](./bundlet_bundler.c) and `mainbundler`, as the user will only need to run the Bundlet binary.
