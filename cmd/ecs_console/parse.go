package main

import (
  "flag"
  "github.com/yyamanoi1222/ecs_console/internal/runner"
)

func parseArg(c *runner.Config) {
  var (
    taskDefinition string
    clusterName string
  )

  flag.StringVar(&taskDefinition, "task-def", "", "ecs task definition")
  flag.StringVar(&clusterName, "cluster", "", "ecs cluster name")
  flag.Parse()
}
