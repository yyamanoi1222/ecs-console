package main

import (
  "github.com/urfave/cli/v2"
  "os"
  "github.com/yyamanoi1222/ecs_console/internal/runner"
  "log"
)

func main() {
  app := &cli.App{
    Name: "ecs_console",
    Commands: []*cli.Command{
      {
        Name: "exec",
        Action: func(cCtx *cli.Context) error {
          err := runner.Run(&runner.Config{
            ClusterName: cCtx.String("cluster"),
            TaskDefinition: cCtx.String("task-def"),
            Command: cCtx.String("command"),
            Container: cCtx.String("container"),
            Subnets: cCtx.String("subnets"),
            SecurityGroups: cCtx.String("security-groups"),
          })
          if err != nil {
            log.Fatal(err)
          }
          return err
        },
        Flags: []cli.Flag{
          &cli.StringFlag{
            Name: "cluster",
            Usage: "ECS Cluster Name",
          },
          &cli.StringFlag{
            Name: "task-def",
            Usage: "ECS Taskdefinition arn",
          },
          &cli.StringFlag{
            Name: "command",
            Value: "/bin/bash",
            Usage: "command passing to ecs-exec",
          },
          &cli.StringFlag{
            Name: "container",
            Value: "app",
            Usage: "container name for ecs-exec",
          },
          &cli.StringFlag{
            Name: "subnets",
            Usage: "subnets name for task",
          },
          &cli.StringFlag{
            Name: "security-groups",
            Usage: "sg for task",
          },
        },
      },
    },
  }
  app.Run(os.Args)
}
