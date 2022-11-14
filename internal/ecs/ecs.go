package ecs

import (
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/aws"
  aecs "github.com/aws/aws-sdk-go/service/ecs"
  "strings"
  "log"
)

type CreateTaskConfig struct {
  TaskDefinition string
  ClusterName string
  Subnets string
  SecurityGroups string
}
type StopTaskConfig struct {
  TaskArn string
  ClusterName string
}

func RunEcsTask(c CreateTaskConfig) (taskArn string, err error) {
  sess := session.Must(session.NewSessionWithOptions(session.Options{SharedConfigState: session.SharedConfigEnable}))
  ecs := aecs.New(sess)

  log.Printf("executing run-task")
  o, err := ecs.RunTask(&aecs.RunTaskInput{
    LaunchType: aws.String("FARGATE"),
    EnableExecuteCommand: aws.Bool(true),
    Cluster: aws.String(c.ClusterName),
    TaskDefinition: aws.String(c.TaskDefinition),
    NetworkConfiguration: &aecs.NetworkConfiguration{
      AwsvpcConfiguration: &aecs.AwsVpcConfiguration{
        Subnets: aws.StringSlice(strings.Split(c.Subnets, ",")),
        SecurityGroups: aws.StringSlice(strings.Split(c.SecurityGroups, ",")),
        AssignPublicIp: aws.String("ENABLED"),
      },
    },
  })

  if err != nil {
    return "", err
  }

  taskArn = *o.Tasks[0].TaskArn

  log.Printf("waiting for running task %s", taskArn)
  err = ecs.WaitUntilTasksRunning(&aecs.DescribeTasksInput{
    Cluster: aws.String(c.ClusterName),
    Tasks: []*string{
      &taskArn,
    },
  })

  return taskArn, err
}

func StopEcsTask(c StopTaskConfig) error {
  sess := session.Must(session.NewSessionWithOptions(session.Options{SharedConfigState: session.SharedConfigEnable}))
  ecs := aecs.New(sess)

  log.Printf("executing stop-task %s", c.TaskArn)
  _, err := ecs.StopTask(&aecs.StopTaskInput{
    Cluster: aws.String(c.ClusterName),
    Task: aws.String(c.TaskArn),
  })
  return err
}
