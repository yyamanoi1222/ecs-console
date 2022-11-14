package main

import (
  "github.com/urfave/cli/v2"
  "os"
  "fmt"
)

func main() {
  cli.AppHelpTemplate = fmt.Sprintf(`NAME:
    {{.Name}}
USAGE:
    {{.HelpName}} {{if .VisibleFlags}}[options]{{end}}
OPTIONS:
    {{range .VisibleFlags}}{{.}}
    {{end}}
`)

  app := &cli.App{
    Name: "ecs_console",
    Action: func(*cli.Context) error {
      fmt.Println("ok")
      return nil
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
    },
  }
  app.Run(os.Args)

  /*
  if err := runner.Run(c); err != nil {
    fmt.Printf("error %s \n", err)
    os.Exit(1)
  }
  */
}
