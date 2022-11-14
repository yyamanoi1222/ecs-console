package ecs_exec

import (
  "os/exec"
  "os"
  "log"
)

type Config struct {
  ClusterName string
  Container string
  TaskArn string
  Command string
}

func Start(c Config) error {
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

  log.Printf("executing ecs execute-command \ncommand: %s \n", cmd.String())

  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  return cmd.Run()
}
