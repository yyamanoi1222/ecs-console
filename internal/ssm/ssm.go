package ssm

import (
  "os/exec"
  "os"
  "fmt"
  "log"
)

type Config struct {
  ClusterName string
  TaskId string
  ContainerId string
  LocalPort string
  RemotePort string
}

func StartPortforward(c Config) error {
  target := fmt.Sprintf("ecs:%s_%s_%s", c.ClusterName, c.TaskId, c.ContainerId)
  parameters := fmt.Sprintf(`{"portNumber":["%s"],"localPortNumber":["%s"]}`, c.RemotePort, c.LocalPort)

  cmd := exec.Command(
    "aws",
    "ssm",
    "start-session",
    "--target",
    target,
    "--document-name",
    "AWS-StartPortForwardingSession",
    "--parameters",
    parameters,
  )

  log.Printf("executing start-session \ncommand: %s \n", cmd.String())

  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  return cmd.Run()
}
