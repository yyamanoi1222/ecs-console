package runner

import (
  "github.com/yyamanoi1222/ecs_console/internal/ecs"
  "github.com/yyamanoi1222/ecs_console/internal/ecs_exec"
  "github.com/yyamanoi1222/ecs_console/internal/ssm"
  "time"
  "os"
  "os/signal"
  "syscall"
  "strings"
)

type ExecConfig struct {
  ClusterName string
  TaskDefinition string
  Command string
  Container string
  Subnets string
  SecurityGroups string
}

type PortforwardConfig struct {
  ClusterName string
  TaskDefinition string
  Container string
  Subnets string
  SecurityGroups string
  LocalPort string
  RemotePort string
}

func Exec(c *ExecConfig) (err error) {
  taskArn := ""

  defer func() {
    // Stop ECS Task
    if len(taskArn) > 0 {
      ecs.StopEcsTask(ecs.StopTaskConfig{
        ClusterName: c.ClusterName,
        TaskArn: taskArn,
      })
    }
    return
  }()

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
  task, err := ecs.RunEcsTask(ecs.CreateTaskConfig{
    TaskDefinition: c.TaskDefinition,
    ClusterName: c.ClusterName,
    Subnets: c.Subnets,
    SecurityGroups: c.SecurityGroups,
  })
  if err != nil {
    return err
  }
  taskArn = *task.TaskArn

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

  return err
}

func Portforward(c *PortforwardConfig) (err error) {
  taskArn := ""

  defer func() {
    // Stop ECS Task
    if len(taskArn) > 0 {
      ecs.StopEcsTask(ecs.StopTaskConfig{
        ClusterName: c.ClusterName,
        TaskArn: taskArn,
      })
    }
    return
  }()


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
  task, err := ecs.RunEcsTask(ecs.CreateTaskConfig{
    TaskDefinition: c.TaskDefinition,
    ClusterName: c.ClusterName,
    Subnets: c.Subnets,
    SecurityGroups: c.SecurityGroups,
  })
  if err != nil {
    return err
  }
  taskArn = *task.TaskArn

  time.Sleep(time.Second * 20)

  spTaskArn := strings.Split(taskArn, "/")
  taskId := spTaskArn[len(spTaskArn) - 1]

  containerId, err := ecs.GetContainerId(c.ClusterName, taskId, c.Container)
  if err != nil {
    return err
  }

  // Run Portforward
  err = ssm.StartPortforward(ssm.Config{
    ClusterName: c.ClusterName,
    ContainerId: containerId,
    Container: c.Container,
    LocalPort: c.LocalPort,
    RemotePort: c.RemotePort,
    TaskId: taskId,
  })

  if err != nil {
    return err
  }
  return err
}
