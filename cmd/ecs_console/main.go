package main

import (
  "github.com/yyamanoi1222/ecs_console/internal/runner"
  "os"
  "fmt"
)

func main() {
  c := &runner.Config{}
  parseArg(c)

  if err := runner.Run(c); err != nil {
    fmt.Printf("error %s \n", err)
    os.Exit(1)
  }
}
