package main

import (
  "github.com/urfave/cli/v2"
  "os"
  "github.com/yyamanoi1222/ecs-console/internal/runner"
  "log"
)

func main() {
  cli.AppHelpTemplate = `NAME:
   {{.Name}}
USAGE:
   {{.HelpName}} {{if .Commands}} command [command options]{{end}}
   {{if len .Authors}}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
   {{.Copyright}}
   {{end}}{{if .Version}}
VERSION:
   {{.Version}}
   {{end}}
`
  app := &cli.App{
    Name: "ecs-console",
    Version:  "v0.0.0",
    Commands: []*cli.Command{
      {
        Name: "exec",
        Action: func(cCtx *cli.Context) error {
          err := runner.Exec(&runner.ExecConfig{
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
            Required: true,
          },
          &cli.StringFlag{
            Name: "task-def",
            Usage: "ECS Taskdefinition name (ex. hoge:latest)",
            Required: true,
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
            Required: true,
          },
          &cli.StringFlag{
            Name: "subnets",
            Usage: "subnets name for task placement",
            Required: true,
          },
          &cli.StringFlag{
            Name: "security-groups",
            Usage: "sg for task placement",
            Required: true,
          },
        },
      },
      {
        Name: "portforward",
        Action: func(cCtx *cli.Context) error {
          err := runner.Portforward(&runner.PortforwardConfig{
            ClusterName: cCtx.String("cluster"),
            TaskDefinition: cCtx.String("task-def"),
            Container: cCtx.String("container"),
            Subnets: cCtx.String("subnets"),
            SecurityGroups: cCtx.String("security-groups"),
            LocalPort: cCtx.String("local-port"),
            RemotePort: cCtx.String("remote-port"),
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
            Required: true,
          },
          &cli.StringFlag{
            Name: "task-def",
            Usage: "ECS Taskdefinition name (ex. hoge:latest)",
            Required: true,
          },
          &cli.StringFlag{
            Name: "container",
            Value: "app",
            Usage: "container name for ecs-exec",
            Required: true,
          },
          &cli.StringFlag{
            Name: "subnets",
            Usage: "subnets name for task placement",
            Required: true,
          },
          &cli.StringFlag{
            Name: "security-groups",
            Usage: "sg for task placement",
            Required: true,
          },
          &cli.StringFlag{
            Name: "remote-port",
            Usage: "remote port",
            Required: true,
          },
          &cli.StringFlag{
            Name: "local-port",
            Usage: "local port",
            Required: true,
          },
        },
      },
    },
  }
  app.Run(os.Args)
}
