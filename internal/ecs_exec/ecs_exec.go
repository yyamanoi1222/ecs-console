package ecs_exec

import (
  "os/exec"
  "os"
  "log"
  "github.com/avast/retry-go/v4"
  "time"
)

type Runner interface {
  Command(name string, arg ...string) *exec.Cmd
}

type Config struct {
  Runner Runner
  ClusterName string
  Container string
  TaskArn string
  Command string
}

func Start(c Config) error {
  err := CheckAgentRunning(c)
  if err != nil {
    return err
  }

  cmd := c.Runner.Command(
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

func CheckAgentRunning(c Config) error {
  return retry.Do(
    func() error {
      cmd := c.Runner.Command(
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
        "echo ok",
      )
      return cmd.Run()
    },
    retry.DelayType(func(n uint, err error, config *retry.Config) time.Duration {
      return retry.BackOffDelay(n, err, config)
    }),
    retry.Delay(time.Second),
  )
}
