# Documentation
In order to use bundlet, you and the user need:
```
- a linux operating system
- the go programming language
- the GCC (GNU C Compiler)
```
You must first now edit the files, such as bundlet.go:

worker.c:
```
#include <stdlib.h>

int main() {
  system("go build example.go"); // edit "example.go" with the name of your go file, eg; "craftmine_real.go"
  return 0;
}
```
bundlet.go:
```
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
To set it up, you as a dev must run:
```
gcc bundlet_bundler.c -o mainbundler
```
```
./mainbundler
```
You have now compiled the basic bundlet runtime.
# Shipping it:
We reccomend deleting
- worker.c | It is now compiled as a seperate binary.
- bundlet_bundler | The user will only have to run the bundlet binary. Not the Compiler that Compiles the Bundlet Compiler 💀
- Also make example.go your actual project lmao.
# Important
*"im not sure if this works..?"* - **Solhost** - *2026*
