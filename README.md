<p align="center">
<img src="assets/gb.md.png" alt="logo">
</p>

# Documentation
In order to use Bundlet, Both the Developer and User need these tools:
```
- A Linux Desktop
- Golang/Go (This is mainly needed for the "go build" tool.)
- The GCC (GNU C Compiler) (Only the Developer needs the GNU C Compiler, it is necessary for bundlet_bundler to run.)
```
The Developer **must** first now edit the files, such as bundlet.go:

worker.c:
```
#include <stdlib.h>

int main() {
  system("go build example.go"); // edit "example.go" with the name of your go file, eg; "my.util.go"
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
  // change this variable the project being compiled;
  compile := "Example" // change "Example" to your project name, eg; "GreenUtils"
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
To set it up, The developer must run:
```
gcc bundlet_bundler.c -o mainbundler

**// Compiles the Bundlet Bundler, Without the packages being compiled, Bundlet cannot run.**
```
```
./mainbundler
```
You have now compiled the basic Bundlet package!
# Shipping it:
We recommend you delete these files when you decide to create a release for your project.
- worker.c | It is now compiled as a seperate binary and Bundlet does not need the source code.
- bundlet_bundler.go and mainbundler | The user will only need to run the Bundlet Binary, The source code is no longer needed.
## Notes
Bundlet is not a bundler, It is a build tool. I made this name before I knew what bundler meant. But Bundlet is a template which kind of loops back to a actual bundler?
