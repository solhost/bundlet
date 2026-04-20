package main

import (
  "fmt"
  "os"
  "os/exec"
  "time"
)

func main() {
  // make this variable the project being compiled;
  compile := "Example"
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
