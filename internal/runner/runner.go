package runner

import (
  "github.com/yyamanoi1222/ecs_console/internal/ecs"
  "github.com/yyamanoi1222/ecs_console/internal/ecs_exec"
)

type Config struct {
  ClusterName string
  TaskDefinition string
  Command string
  Container string
}

func Run(c *Config) error {
  // Create ECS Task
  taskArn, err := ecs.RunEcsTask(ecs.CreateTaskConfig{
    TaskDefinition: c.TaskDefinition,
    ClusterName: c.ClusterName,
  })
  if err != nil {
    return err
  }

  // Run ECS Exec
  ecs_exec.Start(ecs_exec.Config{
    ClusterName: c.ClusterName,
    Container: c.Container,
    TaskArn: taskArn,
    Command: c.Command,
  })

  // Stop ECS Task
  err = ecs.StopEcsTask(ecs.StopTaskConfig{
    ClusterName: c.ClusterName,
    TaskArn: taskArn,
  })
  return err
}
