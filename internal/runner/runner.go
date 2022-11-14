package runner

import (
  "github.com/yyamanoi1222/ecs_console/internal/ecs"
  "github.com/yyamanoi1222/ecs_console/internal/ecs_exec"
  "time"
  "os"
  "os/signal"
  "syscall"
)

type Config struct {
  ClusterName string
  TaskDefinition string
  Command string
  Container string
  Subnets string
  SecurityGroups string
}

func Run(c *Config) error {
  taskArn := ""

  go func() {
    sig := make(chan os.Signal, 1)
    signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTSTP)

    _ = <-sig

    if len(taskArn) > 0 {
      ecs.StopEcsTask(ecs.StopTaskConfig{
        ClusterName: c.ClusterName,
        TaskArn: taskArn,
      })
    }
    os.Exit(1)
  }()

  // Create ECS Task
  taskArn, err := ecs.RunEcsTask(ecs.CreateTaskConfig{
    TaskDefinition: c.TaskDefinition,
    ClusterName: c.ClusterName,
    Subnets: c.Subnets,
    SecurityGroups: c.SecurityGroups,
  })
  if err != nil {
    return err
  }

  time.Sleep(time.Second * 20)

  // Run ECS Exec
  err = ecs_exec.Start(ecs_exec.Config{
    ClusterName: c.ClusterName,
    Container: c.Container,
    TaskArn: taskArn,
    Command: c.Command,
  })

  if err != nil {
    return err
  }

  // Stop ECS Task
  err = ecs.StopEcsTask(ecs.StopTaskConfig{
    ClusterName: c.ClusterName,
    TaskArn: taskArn,
  })

  return err
}
