package ecs_exec

import (
  "os/exec"
  "os"
)

type Config struct {
  ClusterName string
  Container string
  TaskArn string
  Command string
}

func Start(c Config) {
  cmd := exec.Command(
    "aws",
    "ecs",
    "execute-command",
    "--cluster",
    c.ClusterName,
    "--task",
    c.TaskArn,
    "--container",
    c.Container,
    "--interactive",
    "--command",
    c.Command,
  )

  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  cmd.Run()
}
