package ssm

import (
  "github.com/yyamanoi1222/ecs-console/internal/ecs_exec"
  "os/exec"
  "os"
  "fmt"
  "log"
)

type Config struct {
  ClusterName string
  TaskId string
  ContainerId string
  Container string
  LocalPort string
  RemotePort string
}

type CommandRunner struct {}
func (c CommandRunner) Command(name string, arg ...string) *exec.Cmd {
  return exec.Command(name, arg...)
}

func StartPortforward(c Config) error {
  err := ecs_exec.CheckAgentRunning(ecs_exec.Config{
    Runner: CommandRunner{},
    ClusterName: c.ClusterName,
    Container: c.Container,
    TaskArn: c.TaskId,
  })
  if err != nil {
    return err
  }

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
