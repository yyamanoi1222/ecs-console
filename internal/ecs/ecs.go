package ecs

import (
  "github.com/aws/aws-sdk-go/aws"
  aecs "github.com/aws/aws-sdk-go/service/ecs"
  "strings"
  "log"
)

type EcsClient interface {
  RunTask(i *aecs.RunTaskInput) (*aecs.RunTaskOutput, error)
  StopTask(i *aecs.StopTaskInput) (*aecs.StopTaskOutput, error)
  WaitUntilTasksRunning(i *aecs.DescribeTasksInput) error
  DescribeTasks(i *aecs.DescribeTasksInput) (*aecs.DescribeTasksOutput, error)
}

type CreateTaskConfig struct {
  Client EcsClient
  TaskDefinition string
  ClusterName string
  Subnets string
  SecurityGroups string
}
type StopTaskConfig struct {
  Client EcsClient
  TaskArn string
  ClusterName string
}

func RunEcsTask(c CreateTaskConfig) (*aecs.Task, error) {
  log.Printf("executing run-task")
  o, err := c.Client.RunTask(&aecs.RunTaskInput{
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
  err = c.Client.WaitUntilTasksRunning(&aecs.DescribeTasksInput{
    Cluster: aws.String(c.ClusterName),
    Tasks: []*string{
      &taskArn,
    },
  })

  return task, err
}

func StopEcsTask(c StopTaskConfig) error {
  log.Printf("executing stop-task %s", c.TaskArn)
  _, err := c.Client.StopTask(&aecs.StopTaskInput{
    Cluster: aws.String(c.ClusterName),
    Task: aws.String(c.TaskArn),
  })
  return err
}

func GetContainerId(client EcsClient, clusterName string, taskId string, containerName string) (string, error) {
  res, err := client.DescribeTasks(&aecs.DescribeTasksInput{
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
