package ecs

import (
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/aws"
  aecs "github.com/aws/aws-sdk-go/service/ecs"
)

type CreateTaskConfig struct {
  TaskDefinition string
  ClusterName string
}
type StopTaskConfig struct {
  TaskArn string
  ClusterName string
}

func RunEcsTask(c CreateTaskConfig) (taskArn string, err error) {
  sess := session.Must(session.NewSessionWithOptions(session.Options{SharedConfigState: session.SharedConfigEnable}))
  ecs := aecs.New(sess)

  _, err = ecs.RunTask(&aecs.RunTaskInput{
    Cluster: aws.String(c.ClusterName),
    TaskDefinition: aws.String(c.TaskDefinition),
  })

  if err != nil {
    return "", err
  }

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

  _, err := ecs.StopTask(&aecs.StopTaskInput{
    Cluster: aws.String(c.ClusterName),
    Task: aws.String(c.TaskArn),
  })
  return err
}
