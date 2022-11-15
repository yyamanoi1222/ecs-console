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

func RunEcsTask(c CreateTaskConfig) (*aecs.Task, error) {
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
    return &aecs.Task{}, err
  }

  task := o.Tasks[0]
  taskArn := *task.TaskArn

  log.Printf("waiting for running task %s", taskArn)
  err = ecs.WaitUntilTasksRunning(&aecs.DescribeTasksInput{
    Cluster: aws.String(c.ClusterName),
    Tasks: []*string{
      &taskArn,
    },
  })

  return task, err
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

func GetContainerId(clusterName string, taskId string, containerName string) (string, error) {
  sess := session.Must(session.NewSessionWithOptions(session.Options{SharedConfigState: session.SharedConfigEnable}))
  ecs := aecs.New(sess)
  res, err := ecs.DescribeTasks(&aecs.DescribeTasksInput{
    Cluster: aws.String(clusterName),
    Tasks: []*string{aws.String(taskId)},
  })

  if err != nil {
    return "", err
  }

  task := res.Tasks[0]
  container := task.Containers[0]
  for _, cc := range task.Containers {
    if *cc.Name == containerName {
      container = cc
    }
  }

  return *container.RuntimeId, nil
}
